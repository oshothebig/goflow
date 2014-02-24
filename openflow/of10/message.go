package of10

import . "github.com/oshothebig/goflow/openflow"

type FeaturesRequest struct {
	Header
}

type FeaturesReply struct {
	Header
	DatapathId   DatapathId
	Buffers      uint32
	Tables       uint8
	pad          [3]uint8
	Capabilities Capability
	Actions      ActionType
	Ports        []PhysicalPort
}

type DatapathId uint64
type Capability uint32
