package main

import (
	"fmt"
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
		cnt, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("write data, cnt =", cnt)

		time.Sleep(time.Second)

	}

}
