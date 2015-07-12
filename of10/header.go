package of10

import "io"

const HeaderLength = 8

type Packetizable interface {
	io.ReadWriter
	Len() uint
}

type Message interface {
	GetHeader() *Header
	FillBody(bytes []byte) error
}

type Header struct {
	Version uint8
	Type    MessageType
	Length  uint16
	Xid     uint32
}

func (h *Header) GetHeader() *Header {
	return h
}

type GenericMessage struct {
	Header
	Payload []byte
}

func NewXidGenerator() func() uint32 {
	var xid uint32 = 0
	return func() uint32 {
		xid += 1
		return xid
	}
}

type MessageType uint8
