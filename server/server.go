package main

import (
	"huy-database/protocol"
	"log"
	"net"
)

func main() {
	addr := ":4000"
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Error listening on " + addr)
		return
	}
	log.Println("Listening on " + addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection " + err.Error())
			continue
		}
		go func() {
			err := handleConnection(conn)
			if err != nil {
				log.Fatalf("Error handling connection %s, because %s", conn.RemoteAddr(), err.Error())
			}
		}()
	}
}

func handleConnection(conn net.Conn) error {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing connection " + err.Error())
		}
	}(conn)

	log.Println("Handling connection " + conn.RemoteAddr().String())

	for {
		req, err := protocol.DecodeRequest(conn)
		if err != nil {
			err = sendErr(conn, "Error decoding request: "+err.Error())
			return err
		}

		if req.Version != protocol.Version {
			err = sendErr(conn, "Unsupported protocol version")
			return err
		}

		switch req.MessageType {
		case protocol.MsgTypeQuery:
			res, err := handleQuery(req)
			if err != nil {
				err = sendErr(conn, "Error handling query: "+err.Error())
				return err
			}
			err = sendResponse(conn, res)
		default:
			err = sendErr(conn, "Unsupported message type")
		}
		return err
	}
}

func sendErr(conn net.Conn, msg string) error {
	res := &protocol.Response{
		ResponseType: protocol.ReturnCodeError,
		Length:       uint32(len(msg)),
		Payload:      []byte(msg),
	}
	err := sendResponse(conn, res)
	return err
}

func sendResponse(conn net.Conn, response *protocol.Response) error {
	data := protocol.EncodeResponse(response)
	_, err := conn.Write(data)
	return err
}
