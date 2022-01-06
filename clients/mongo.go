package clients

import (
	"github.com/chengcxy/datashift/scheduler"
	"github.com/chengcxy/gotools/configor"
	"log"
)

type MongoClient struct {
}

func (m *MongoClient) Read(query interface{}, writer scheduler.Client) (*scheduler.WriteResult, error) {
	q := query.(map[string]string)
	sql := q["sql"]
	log.Printf("sql is %s", sql)
	data := make([]interface{}, 0)
	data = append(data, 0)
	data = append(data, 2)
	rr := &scheduler.ReadResult{
		Data:  data,
		Error: nil,
	}
	return writer.Write(rr)
}

func (m *MongoClient) Write(rr *scheduler.ReadResult) (wr *scheduler.WriteResult, err error) {
	if rr.Error != nil {
		err = rr.Error
		wr = &scheduler.WriteResult{
			Status: 0,
			Error:  rr.Error,
		}
		return
	}
	data := rr.Data
	log.Println("mongo write data", data)
	wr = &scheduler.WriteResult{
		Status: 1,
		Error:  nil,
	}
	return
}

func (m *MongoClient) Connect(config *configor.Config) scheduler.Client {
	return m
}

func (m *MongoClient) Close() {

}
func init() {
	scheduler.Register("mongo", &MongoClient{})
}
