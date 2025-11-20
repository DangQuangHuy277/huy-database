package main

import (
	"huy-database/protocol"
)

func handleQuery(req *protocol.Request) (*protocol.Response, error) {
	return &protocol.Response{
		ResponseType: protocol.ReturnCodeData,
		Length:       uint32(len(req.Payload)),
		Payload:      req.Payload,
	}, nil
}
