package main

import (
	"log"
	"net"
	"time"

	"github.com/aca/go/netutil"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("helloworld\n"))
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 3)

	err = netutil.ConnCheck(conn)
	if err != nil {
		log.Fatal(err)
	}
}
