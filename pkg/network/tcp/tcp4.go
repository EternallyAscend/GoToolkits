package tcp

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
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
	connection, err := net.Dial("tcp", net.JoinHostPort(address, strconv.Itoa(int(port))))
	if nil != err {
		log.Println(err)
		return
	}
	defer connection.Close()
	handler(&connection)
}

func ListenViaTcp4(handler func(*net.Conn), port uint) {
	connection, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if nil != err {
		log.Println(err)
		return
	}
	defer connection.Close()

	for {
		cli, errIn := connection.Accept()
		if nil != errIn {
			log.Println(errIn)
			continue
		}
		go handler(&cli)
	}
}

func ListenInterruptableViaTcp4(ctx *context.Context, handler func(*net.Conn), port uint) {
	connection, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if nil != err {
		log.Println(err)
		return
	}
	defer connection.Close()
	loop:
	for {
		cli, errIn := connection.Accept()
		if nil != errIn {
			log.Println(errIn)
			continue
		}
		go handler(&cli)
		c := *ctx
		select {
		case <- c.Done():
			break loop
		}
	}
}
