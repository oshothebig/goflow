package openflow

import "io"

type Packetizable interface {
	io.ReadWriter
	Len() uint
}

type Message interface {
	Packetizable
	Header() *Header
}

type Header struct {
	Version uint8
	Type    MessageType
	Length  uint16
	Xid     uint32
}

func (h *Header) Header() *Header {
	return h
}

func NewXidGenerator() func() uint32 {
	var xid uint32 = 0
	return func() uint32 {
		xid += 1
		return xid
	}
}

type MessageType uint8
