package clients

import (
	"github.com/chengcxy/gotools/configor"
	"log"
)

type MongoClient struct {
}

func (m *MongoClient) Read(query interface{}, writer Client) (*WriteResult, error) {
	q := query.(map[string]string)
	sql := q["sql"]
	log.Printf("sql is %s", sql)
	data := make([]interface{}, 0)
	data = append(data, 0)
	data = append(data, 2)
	rr := &ReadResult{
		Data:  data,
		Error: nil,
	}
	return writer.Write(rr)
}

func (m *MongoClient) Write(rr *ReadResult) (wr *WriteResult, err error) {
	if rr.Error != nil {
		err = rr.Error
		wr = &WriteResult{
			Status: 0,
			Error:  rr.Error,
		}
		return
	}
	data := rr.Data
	log.Println("mongo write data", data)
	wr = &WriteResult{
		Status: 1,
		Error:  nil,
	}
	return
}

func (m *MongoClient) Connect(config *configor.Config) Client {
	return m
}

func (m *MongoClient) Close() {

}
func init() {
	Register("mongo", &MongoClient{})
}
