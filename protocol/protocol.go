package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Protocol constants
const (
	Version = 0x01

	// MsgTypeQuery Message types (Request)
	MsgTypeQuery = 0x00

	// ReturnCodeData Return codes (Response)
	ReturnCodeData     = 0x00
	ReturnCodeComplete = 0x01
	ReturnCodeError    = 0x02
)

type Request struct {
	Version     byte
	MessageType byte
	Length      uint32
	Payload     []byte
}

type Response struct {
	ResponseType byte
	Length       uint32
	Payload      []byte
}

func EncodeRequest(request *Request) []byte {
	reqBytes := make([]byte, 6+len(request.Payload))
	reqBytes[0] = request.Version
	reqBytes[1] = request.MessageType
	binary.BigEndian.PutUint32(reqBytes[2:6], request.Length)
	copy(reqBytes[6:], request.Payload)
	return reqBytes
}

func DecodeRequest(r io.Reader) (*Request, error) {
	header := make([]byte, 6)
	if _, err := io.ReadFull(r, header); err != nil {
		return &Request{}, fmt.Errorf("error reading header: %v", err)
	}

	req := Request{
		Version:     header[0],
		MessageType: header[1],
		Length:      binary.BigEndian.Uint32(header[2:6]),
	}
	req.Payload = make([]byte, req.Length)
	if _, err := io.ReadFull(r, req.Payload); err != nil {
		return &Request{}, fmt.Errorf("error reading payload: %v", err)
	}
	return &req, nil
}

func EncodeResponse(response *Response) []byte {
	resBytes := make([]byte, 5+len(response.Payload))
	resBytes[0] = response.ResponseType
	binary.BigEndian.PutUint32(resBytes[1:5], response.Length)
	copy(resBytes[5:], response.Payload)
	return resBytes
}

func DecodeResponse(r io.Reader) (*Response, error) {
	header := make([]byte, 5)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, fmt.Errorf("error reading header: %v", err)
	}
	resp := Response{
		ResponseType: header[0],
		Length:       binary.BigEndian.Uint32(header[1:5]),
	}
	resp.Payload = make([]byte, resp.Length)
	if _, err := io.ReadFull(r, resp.Payload); err != nil {
		return nil, fmt.Errorf("error reading payload: %v", err)
	}
	return &resp, nil
}
