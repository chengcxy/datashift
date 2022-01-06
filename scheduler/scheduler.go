package scheduler

import (
	"github.com/chengcxy/datashift/clients"
	"github.com/chengcxy/gotools/configor"
	"log"
	"strings"
)

type Scheduler struct {
	Config *configor.Config
}

func (s *Scheduler) Run() {
	from, _ := s.Config.Get("reader.type")
	reader := clients.GetClient(from.(string))
	reader = reader.Connect(s.Config)
	log.Printf("reader %v", reader)
	to, _ := s.Config.Get("writer.type")
	writer := clients.GetClient(to.(string))
	writer = writer.Connect(s.Config)
	defer func() {
		reader.Close()
		writer.Close()
	}()
	query := make(map[string]string)
	query["sql"] = "select id,entity_id from z_pe.base_entity_basic_info limit 1000"
	query["method"] = "GET"
	query["url"] = "http://127.0.0.1:5000/v1/api/"
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
