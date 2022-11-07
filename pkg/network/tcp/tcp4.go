package tcp

import (
	"context"
	"fmt"
	"log"
	"net"
)

// https://studygolang.com/articles/9240

const DefaultTcpBufferSize = 1024

func ReadFuncTcp4(conn *net.Conn) ([]byte, error) {
	data := make([]byte, DefaultTcpBufferSize)
	_, err := (*conn).Read(data)
	if nil != err {
		return nil, err
	}
	return data, nil
}

func WriteFuncTcp4(conn *net.Conn, data []byte) error {
	_, err := (*conn).Write(data)
	return err
}

func RequestViaTcp4(address string, port uint, handler func(*net.Conn)) {
	connection, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, port))
	defer connection.Close()
	if nil != err {
		log.Println(err)
		return
	}
	handler(&connection)
}

func ListenViaTcp4(port uint, handler func(*net.Conn)) {
	connection, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	defer connection.Close()
	if nil != err {
		log.Println(err)
		return
	}

	for true {
		cli, errIn := connection.Accept()
		if nil != errIn {
			log.Println(errIn)
			continue
		}
		go handler(&cli)
	}
}

func ListenInterruptableViaTcp4(ctx *context.Context, port uint, handler func(*net.Conn)) {
	connection, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	defer connection.Close()
	if nil != err {
		log.Println(err)
		return
	}
loop:
	for true {
		cli, errIn := connection.Accept()
		if nil != errIn {
			log.Println(errIn)
			continue
		}
		go handler(&cli)
		select {
		case <-(*ctx).Done():
			break loop
		}
	}
}
