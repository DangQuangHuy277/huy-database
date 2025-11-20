package main

import (
	"bufio"
	"fmt"
	"huy-database/protocol"
	"net"
	"os"
	"strings"
)

func main() {
	addr := "localhost:4000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	fmt.Println("Connected to ", addr)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("sql> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if line == "quit" || line == "exit" {
			fmt.Println("Bye.")
			return
		}

		res, err := executeQuery(conn, line)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
}

func executeQuery(conn net.Conn, query string) (string, error) {
	payload := []byte(query)
	req := &protocol.Request{
		Version:     protocol.Version,
		MessageType: protocol.MsgTypeQuery,
		Length:      uint32(len(payload)),
		Payload:     payload,
	}
	res, err := sendRequest(conn, req)
	if err != nil {
		return "", err
	}
	return string(res.Payload), nil

}

func sendRequest(conn net.Conn, req *protocol.Request) (*protocol.Response, error) {
	buf := protocol.EncodeRequest(req)
	if _, err := conn.Write(buf); err != nil {
		return &protocol.Response{}, err
	}

	for {
		res, err := protocol.DecodeResponse(conn)
		if err != nil {
			return &protocol.Response{}, err
		}
		if res != nil {
			return res, nil
		}
	}
}
