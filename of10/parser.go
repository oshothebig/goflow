package of10

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
)

var (
	errUnsupportedMessage error = errors.New("Unsupported message type")
)

func readMessage(reader *bufio.Reader) (Message, error) {
	headerBytes := make([]byte, HeaderLength)
	if _, err := io.ReadFull(reader, headerBytes); err != nil {
		return nil, err
	}

	var header Header
	if err := binary.Read(reader, binary.BigEndian, header); err != nil {
		return nil, err
	}

	msg := newMessage(&header)
	if msg == nil {
		return nil, errUnsupportedMessage
	}

	body := make([]byte, header.Length-HeaderLength)
	if _, err := io.ReadFull(reader, body); err != nil {
		return nil, err
	}
	if err := msg.FillBody(body); err != nil {
		return nil, err
	}

	return msg, nil
}

func newMessage(h *Header) Message {
	switch h.Type {
	case MessageTypes.Hello:
		return &Hello{Header: *h}
	case MessageTypes.Error:
		return &ErrorMessage{Header: *h}
	case MessageTypes.EchoRequest:
		return &EchoRequest{Header: *h}
	case MessageTypes.EchoReply:
		return &EchoReply{Header: *h}
	case MessageTypes.Vendor:
		return &VendorMessage{Header: *h}
	case MessageTypes.FeaturesRequest:
		return &FeaturesRequest{Header: *h}
	case MessageTypes.FeaturesReply:
		return &FeaturesReply{Header: *h}
	case MessageTypes.GetConfigRequest:
		return &GetConfigRequest{Header: *h}
	case MessageTypes.GetConfigReply:
		return &GetConfigReply{Header: *h}
	case MessageTypes.SetConfig:
		return &SetConfig{Header: *h}
	case MessageTypes.PacketIn:
		return &PacketIn{Header: *h}
	case MessageTypes.FlowRemoved:
		return &FlowRemoved{Header: *h}
	case MessageTypes.PortStatus:
		return &PortStatus{Header: *h}
	case MessageTypes.PacketOut:
		return &PacketOut{Header: *h}
	case MessageTypes.FlowMod:
		return &FlowMod{Header: *h}
	case MessageTypes.PortMod:
		return &PortMod{Header: *h}
	case MessageTypes.StatsRequest:
		return &StatsRequest{Header: *h}
	case MessageTypes.StatsReply:
		return &StatsReply{Header: *h}
	case MessageTypes.BarrierRequest:
		return &BarrierRequest{Header: *h}
	case MessageTypes.BarrierReply:
		return &BarrierReply{Header: *h}
	case MessageTypes.QueueGetConfigRequest:
		return &QueueGetConfigRequest{Header: *h}
	case MessageTypes.QueueGetConfigReply:
		return &QueueGetConfigReply{Header: *h}
	default:
		return nil
	}
}
