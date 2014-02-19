package of13

import . "github.com/oshothebig/goflow/openflow"

type QueuePropertyType uint16

type PacketQueue struct {
	QueueId    uint32
	Port       uint32
	Length     uint32
	pad        [6]uint8
	Properties []QueueProperty
}

type QueueProperty interface {
	Packetizable
	Header() *QueuePropertyHeader
}

type QueuePropertyHeader struct {
	Property QueuePropertyType
	Length   uint16
	pad      [4]uint8
}

func (h *QueuePropertyHeader) Header() *QueuePropertyHeader {
	return h
}

type QueuePropertyMinRate struct {
	QueuePropertyHeader
	Rate uint16
	pad  [6]uint8
}

type QueuePropertyMaxRate struct {
	QueuePropertyHeader
	Rate uint16
	pad  [6]uint8
}

type QueuePropertyExperimenter struct {
	QueuePropertyHeader
	Experimenter uint32
	pad          [4]uint8
	Data         []uint8
}

// corresponds to ofp_queue_properties
const (
	OFPQT_MIN_RATE     QueuePropertyType = 1
	OFPQT_MAX_RATE     QueuePropertyType = 2
	OFPQT_EXPERIMENTER QueuePropertyType = 0xffff
)

var QueueProperties = struct {
	MinRate      QueuePropertyType
	MaxRate      QueuePropertyType
	Experimenter QueuePropertyType
}{
	OFPQT_MIN_RATE,
	OFPQT_MAX_RATE,
	OFPQT_EXPERIMENTER,
}
