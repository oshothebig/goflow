package of10

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var (
	errUnsupportedMessage error = errors.New("Unsupported message type")
)

type Decoder struct {
	rd *bufio.Reader
}

func NewDecoder(rd io.Reader) *Decoder {
	return &Decoder{bufio.NewReader(rd)}
}

func (d *Decoder) header() (*Header, error) {
	headerBytes, err := d.rd.Peek(HeaderLength)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewReader(headerBytes)
	var header Header
	if err := binary.Read(buf, binary.BigEndian, &header); err != nil {
		return nil, err
	}
	return &header, nil
}

func (d *Decoder) Decode() (Message, error) {
	header, err := d.header()
	if err != nil {
		return nil, err
	}

	message := emptyMessage(header.Type)
	if message == nil {
		return nil, errors.New("Unsupported message type")
	}

	data := make([]byte, header.Length)
	if _, err := io.ReadFull(d.rd, data); err != nil {
		return nil, err
	}

	if err := message.UnmarshalBinary(data); err != nil {
		return nil, err
	}

	return message, nil
}

func emptyMessage(t MessageType) Message {
	switch t {
	case MessageTypes.Hello:
		return new(Hello)
	case MessageTypes.Error:
		return new(ErrorMessage)
	case MessageTypes.EchoRequest:
		return new(EchoRequest)
	case MessageTypes.EchoReply:
		return new(EchoReply)
	case MessageTypes.Vendor:
		return new(VendorMessage)
	case MessageTypes.FeaturesRequest:
		return new(FeaturesRequest)
	case MessageTypes.FeaturesReply:
		return new(FeaturesReply)
	case MessageTypes.GetConfigRequest:
		return new(GetConfigRequest)
	case MessageTypes.GetConfigReply:
		return new(GetConfigReply)
	case MessageTypes.SetConfig:
		return new(SetConfig)
	case MessageTypes.PacketIn:
		return new(PacketIn)
	case MessageTypes.FlowRemoved:
		return new(FlowRemoved)
	case MessageTypes.PortStatus:
		return new(PortStatus)
	case MessageTypes.PacketOut:
		return new(PacketOut)
	case MessageTypes.FlowMod:
		return new(FlowMod)
	case MessageTypes.PortMod:
		return new(PortMod)
	case MessageTypes.StatsRequest:
		return new(StatsRequest)
	case MessageTypes.StatsReply:
		return new(StatsReply)
	case MessageTypes.BarrierRequest:
		return new(BarrierRequest)
	case MessageTypes.BarrierReply:
		return new(BarrierReply)
	case MessageTypes.QueueGetConfigRequest:
		return new(QueueGetConfigRequest)
	case MessageTypes.QueueGetConfigReply:
		return new(QueueGetConfigReply)
	default:
		return nil
	}
}
