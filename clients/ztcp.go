package clients

import (
	"fmt"
	"github.com/chengcxy/gotools/configor"
	"log"
	"net"
	"strings"
)

type TcpClient struct {
	listener net.Listener
}

func (t *TcpClient) Read(query interface{}, writer Client) (*WriteResult, error) {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Println("error", err)
			continue
		}
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("error", err)
			continue
		}

		cmds := strings.Split(string(buf[:n]), " ")
		fmt.Println(cmds)

		var response []byte
		response = []byte(strings.Join(cmds, "-"))
		conn.Write(response)
	}

	rr := &ReadResult{
		Data:  nil,
		Error: nil,
	}
	return writer.Write(rr)
}

func (t *TcpClient) Write(rr *ReadResult) (wr *WriteResult, err error) {
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

func (t *TcpClient) Connect(config *configor.Config) Client {
	conf, _ := config.Get("reader.dsn")
	listener, _ := net.Listen("tcp", conf.(string))
	t.listener = listener
	return t
}

func (t *TcpClient) Close() {

}
func init() {
	Register("tcp", &TcpClient{})
}
