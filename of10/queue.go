package of10

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var (
	errUnsupportedQueuePropertyType = errors.New("Unsupported queue property type")
)

// TODO: Properties should be generalized
type PacketQueue struct {
	QueueId    uint32
	Length     uint16
	pad        [2]uint8
	Properties []QueueProperty
}

func readPacketQueues(buf *bytes.Reader) []PacketQueue {
	queues := make([]PacketQueue, 0, 8)
	remain := buf.Len()
	for remain != 0 {
		queue, err := readPacketQueue(buf)
		if err != nil {
			break
		}
		queues = append(queues, *queue)
		remain = buf.Len()
	}
	return queues
}

func readPacketQueue(buf *bytes.Reader) (*PacketQueue, error) {
	var queue PacketQueue
	if err := binary.Read(buf, binary.BigEndian, &queue.QueueId); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &queue.QueueId); err != nil {
		return nil, err
	}
	if err := binary.Read(buf, binary.BigEndian, &queue.pad); err != nil {
		return nil, err
	}
	queue.Properties = readQueueProperties(buf)
	return &queue, nil
}

type QueueProperty interface {
	Header() *QueuePropertyHeader
	FillBody(buf *bytes.Buffer) error
}

const queuePropertyHeaderLength = 8

func readQueueProperty(buf *bytes.Reader) (QueueProperty, error) {
	var header QueuePropertyHeader
	if err := binary.Read(buf, binary.BigEndian, &header); err != nil {
		return nil, err
	}

	property := newQueueProperty(&header)
	if property == nil {
		return nil, errUnsupportedQueuePropertyType
	}

	body := make([]byte, header.Length-queuePropertyHeaderLength)
	if _, err := io.ReadFull(buf, body); err != nil {
		return nil, err
	}

	bodyBuf := bytes.NewBuffer(body)
	if err := property.FillBody(bodyBuf); err != nil {
		return nil, err
	}
	return property, nil
}

func newQueueProperty(h *QueuePropertyHeader) QueueProperty {
	switch h.Property {
	case QueuePropertyTypes.None:
		prop := *h
		return &prop
	case QueuePropertyTypes.MinRate:
		return &QueuePropertyMinRate{QueuePropertyHeader: *h}
	default:
		return nil
	}
}

func readQueueProperties(buf *bytes.Reader) []QueueProperty {
	properties := make([]QueueProperty, 0, 8)
	remain := buf.Len()
	for remain != 0 {
		property, err := readQueueProperty(buf)
		if err != nil {
			break
		}
		remain = buf.Len()
		properties = append(properties, property)
	}
	return properties
}

type QueuePropertyType uint16

type QueuePropertyHeader struct {
	Property QueuePropertyType
	Length   uint16
	_        [4]uint8
}

func (h *QueuePropertyHeader) Header() *QueuePropertyHeader {
	return h
}

func (p *QueuePropertyHeader) FillBody(buf *bytes.Buffer) error {
	return nil
}

type QueuePropertyMinRate struct {
	QueuePropertyHeader
	Rate uint16
	_    [6]uint8
}

func (p *QueuePropertyMinRate) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &p.Rate); err != nil {
		return err
	}
	padding := make([]byte, 6)
	if _, err := io.ReadFull(buf, padding); err != nil {
		return err
	}
	return nil
}
