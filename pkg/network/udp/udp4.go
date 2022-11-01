package udp

import (
	"fmt"
	"log"
	"net"
)

// 优化 https://zhuanlan.zhihu.com/p/357902432

func SendViaUdp4(address string, port uint, msg []byte) error {
	connect, err := net.Dial("udp", fmt.Sprintf("%s:%d", address, port))
	if nil != err {
		return err
	}
	defer func(connect net.Conn) {
		_ = connect.Close()
	}(connect)
	_, err = connect.Write(msg)
	return err
}

func ListenViaUdp4(handler func([]byte), port uint) {
	connection, err := net.ListenPacket("udp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Println(err)
	}
	defer connection.Close()

	for true {
		buffer := make([]byte, 1024)
		_, _, err = connection.ReadFrom(buffer)
		if err != nil {
			log.Println(err)
		}
		handler(buffer)
	}

}
