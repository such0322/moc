package main

import (
	"fmt"
	"moc/network"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9100")

	if err != nil {
		fmt.Println(err)
		return
	}

	for {

		dp := network.NewDataPack()
		msg1 := &network.Message{
			ID:      0,
			DataLen: 5,
			Data:    []byte{'h', 'e', 'l', 'l', 'o'},
		}
		sendData1, err := dp.Pack(msg1)
		if err != nil {
			fmt.Println("client pack msg1 err:", err)
			return
		}
		msg2 := &network.Message{
			ID:      1,
			DataLen: 7,
			Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
		}
		sendData2, err := dp.Pack(msg2)
		if err != nil {
			fmt.Println("client pack msg2 err:", err)
			return
		}
		sendData1 = append(sendData1, sendData2...)

		cnt, err := conn.Write(sendData1)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("write data, cnt =", cnt)

		time.Sleep(time.Second)

	}

}
