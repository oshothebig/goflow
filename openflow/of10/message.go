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

const (
	OFPC_FLOW_STATS Capability = 1 << iota
	OFPC_TABLE_STATS
	OFPC_PORT_STATS
	OFPC_STP
	OFPC_RESERVED
	OFPC_IP_REASM
	OFPC_QUEUE_STATS
	OFPC_ARP_MATCH_IP
)

var Capabilities = struct {
	FlowStats    Capability
	TableStats   Capability
	PortStats    Capability
	Stp          Capability
	Reserved     Capability
	IpReassemble Capability
	QueueStats   Capability
	ArpMatchIp   Capability
}{
	OFPC_FLOW_STATS,
	OFPC_TABLE_STATS,
	OFPC_PORT_STATS,
	OFPC_STP,
	OFPC_RESERVED,
	OFPC_IP_REASM,
	OFPC_QUEUE_STATS,
	OFPC_ARP_MATCH_IP,
}
