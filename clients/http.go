package clients

import (
	"encoding/json"
	"github.com/chengcxy/gotools/configor"
	"io/ioutil"
	"log"
	"net/http"
)

type HttpClient struct {
}

func (h *HttpClient) Read(query interface{}, writer Client) (*WriteResult, error) {
	q := query.(map[string]string)
	method := q["method"]
	url := q["url"]
	req, _ := http.NewRequest(method, url, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	data := make(map[string]interface{})
	json.Unmarshal(bodyText, &data)
	items := make([]map[string]interface{}, 1)
	items[0] = data
	rr := &ReadResult{
		Data:  items,
		Error: err,
	}
	return writer.Write(rr)
}

func (h *HttpClient) Write(rr *ReadResult) (wr *WriteResult, err error) {
	if rr.Error != nil {
		err = rr.Error
		wr = &WriteResult{
			Status: 0,
			Error:  rr.Error,
		}
		return
	}
	data := rr.Data
	log.Println("http write data", data)
	wr = &WriteResult{
		Status: 1,
		Error:  nil,
	}
	return
}

func (h *HttpClient) Connect(config *configor.Config) Client {
	return h
}

func (h *HttpClient) Close() {

}
func init() {
	Register("http", &HttpClient{})
}
