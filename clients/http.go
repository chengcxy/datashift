package clients

import (
	"github.com/chengcxy/datashift/scheduler"
	"github.com/chengcxy/gotools/configor"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
}

func (h *HttpClient) Read(query interface{}, writer scheduler.Client) (*scheduler.WriteResult, error) {
	q := query.(map[string]string)
	method := q["method"]
	url := q["url"]
	req, _ := http.NewRequest(method, url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	data := make(map[string]string)
	data["result"] = string(bodyText)
	items := make([]map[string]string, 1)
	items[0] = data
	rr := &scheduler.ReadResult{
		Data:  items,
		Error: err,
	}
	return writer.Write(rr)
}

func (h *HttpClient) Write(rr *scheduler.ReadResult) (wr *scheduler.WriteResult, err error) {
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

func (h *HttpClient) Connect(config *configor.Config) scheduler.Client {
	return h
}

func (h *HttpClient) Close() {

}
func init() {
	scheduler.Register("http", &HttpClient{})
}
