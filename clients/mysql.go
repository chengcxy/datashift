package clients

import (
	"fmt"
	"github.com/chengcxy/gotools/backends"
	"github.com/chengcxy/gotools/configor"
	"log"
	"strings"
)

type MysqlClient struct {
	db *backends.MysqlClient
}

func (m *MysqlClient) Read(query interface{}, writer Client) (*WriteResult, error) {
	q := query.(map[string]string)
	sql := q["sql"]
	log.Printf("sql is %s", sql)
	data, _, err := m.db.Query(sql)
	rr := &ReadResult{
		Data:  data,
		Error: err,
	}
	return writer.Write(rr)
}

func (m *MysqlClient) Write(rr *ReadResult) (wr *WriteResult, err error) {
	if rr.Error != nil {
		err = rr.Error
		wr = &WriteResult{
			Status: 0,
			Error:  rr.Error,
		}
		return
	}
	data := rr.Data
	items := data.([]map[string]string)
	log.Println("mysql write data", items)
	//目前写死 正常应该包含在rr *scheduler.ReadResult对象里面
	baseInsertSql := "insert into blog.blog2(id,title)values %s"
	qs := make([]string, len(items))
	values := make([]interface{}, 0)
	for i := 0; i < len(items); i++ {
		qs[i] = "(?,?)"
		item := items[i]
		values = append(values, item["id"])
		values = append(values, item["title"])
	}
	qsStr := strings.Join(qs, ",")
	InsertSql := fmt.Sprintf(baseInsertSql, qsStr)
	m.db.Execute(InsertSql, values...)
	wr = &WriteResult{
		Status: 1,
		Error:  nil,
	}
	return
}

func (m *MysqlClient) Connect(config *configor.Config) Client {
	db, _ := backends.NewMysqlClient(config, "reader.conn")
	m.db = db
	return m
}
func (m *MysqlClient) Close() {
	m.db.Close()
}
func init() {
	Register("mysql", &MysqlClient{})
}
