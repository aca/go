package main

import (
	"io"
	"os"
	"testing"

	"github.com/aca/go/wip/fakenet"
)

func TestWriteMsg(t *testing.T) {
	t.Log("start")
	// server, client := net.Pipe()
	conn := fakenet.NewConn("stdio", os.Stdin, os.Stdout)

	go func() {
		buf := make([]byte, 10)
		_, err := io.ReadFull(conn, buf)
		if err != nil {
			panic(err)
		}
		t.Log(string(buf))
	}()
	
	err := WriteMsg(conn, "helloworld")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("read")



}
