package of10

// TODO: Properties should be generalized
type PacketQueue struct {
	QueueId    uint32
	Length     uint16
	pad        [2]uint8
	Properties []uint8
}

type QueueProperty interface {
	Packetizable
	Header() *QueuePropertyHeader
}

type QueuePropertyType uint16

const (
	OFPQT_NONE QueuePropertyType = iota
	OFPQT_MIN_RATE
)

var QueuePropertyTypes = struct {
	None    QueuePropertyType
	MinRate QueuePropertyType
}{
	OFPQT_NONE,
	OFPQT_MIN_RATE,
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
