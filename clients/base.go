package clients

import (
	"github.com/chengcxy/gotools/configor"
	"log"
)

type Client interface {
	Connect(config *configor.Config) Client
	Read(query interface{}, writer Client) (*WriteResult, error)
	Write(rr *ReadResult) (*WriteResult, error)
	Close()
}

type ReadResult struct {
	Data  interface{}
	Error error
}

type WriteResult struct {
	Status int
	Error  error
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
