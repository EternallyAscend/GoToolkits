package tcp

import (
	"log"
	"net"
	"testing"
)

func TestListenViaTcp4(t *testing.T) {
	ListenViaTcp4(func(conn *net.Conn) {
		data := make([]byte, 1024)
		length, err := (*conn).Read(data)
		if nil != err {
			log.Println(err)
			return
		}
		log.Println(string(data[0:length]))
		log.Println(string(data))
		WriteFuncTcp4(conn, []byte("Huawei."))
	}, 9000)
}

func TestRequestViaTcp4(t *testing.T) {
	data := []byte("Message.")
	RequestViaTcp4("127.0.0.1", 9000, func(conn *net.Conn, port uint, data []byte) {
		_, err := (*conn).Write(data)
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
	}, 0, data)
}
