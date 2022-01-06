package scheduler

import (
	"github.com/chengcxy/gotools/configor"
	"log"
	"strings"
)

type Scheduler struct {
	Config *configor.Config
}

type Client interface {
	Connect(config *configor.Config) Client
	Read(query interface{}, writer Client) (*WriteResult, error)
	Write(rr *ReadResult) (*WriteResult, error)
	Close()
}

var Clients = make(map[string]Client)

func Register(clinetType string, client Client) {
	if _, exists := Clients[clinetType]; exists {
		log.Fatalln(clinetType, "client already registered")
	}

	log.Println("Register", clinetType, "client")
	Clients[clinetType] = client
}

func GetClient(clinetType string) (client Client) {
	if client, exists := Clients[clinetType]; exists {
		return client
	}
	return nil
}

func (s *Scheduler) Run() {
	from, _ := s.Config.Get("reader.type")
	reader := GetClient(from.(string))
	reader = reader.Connect(s.Config)
	log.Printf("reader %v", reader)
	to, _ := s.Config.Get("writer.type")
	writer := GetClient(to.(string))
	writer = writer.Connect(s.Config)
	defer func() {
		reader.Close()
		writer.Close()
	}()
	query := make(map[string]string)
	query["sql"] = "select id,entity_id from z_pe.base_entity_basic_info limit 1000"
	query["method"] = "GET"
	query["url"] = "http://www.baidu.com"
	wr, err := reader.Read(query, writer)
	if err != nil {
		log.Printf("error %v", err)
	}
	log.Printf("wr:%v", wr)
	var path strings.Builder
	path.Grow(1)
	path.WriteString("/shhh")
	path.WriteString("ssss")
	log.Printf("path:%s", path.String())
	var buf []byte
	buf = append(buf, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"...)
	log.Printf("buf %v", buf)
}
