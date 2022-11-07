package tcp

import (
	"log"
	"net"
	"testing"
)

func TestListenViaTcp4(t *testing.T) {
	ListenViaTcp4(9000, func(conn *net.Conn) {
		data := make([]byte, 1024)
		length, err := (*conn).Read(data)
		if nil != err {
			log.Println(err)
			return
		}
		log.Println(string(data[0:length]))
		log.Println(string(data))
		WriteFuncTcp4(conn, []byte("Huawei."))
	})
}

func TestRequestViaTcp4(t *testing.T) {
	RequestViaTcp4("127.0.0.1", 9000, func(conn *net.Conn) {
		_, err := (*conn).Write([]byte("Message."))
		if err != nil {
			log.Println(err)
			return
		}
		d, err := ReadFuncTcp4(conn)
		if nil != err {
			log.Println(err)
			return
		}
		log.Println(string(d))
	})
}
