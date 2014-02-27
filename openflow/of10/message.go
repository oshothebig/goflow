package of10

import (
	"net"

	. "github.com/oshothebig/goflow/openflow"
)

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

type SwitchConfig struct {
	Header
	Flags          ConfigFlag
	MissSendLength uint16
}

type ConfigFlag uint16

const (
	OFPC_FRAG_NORMAL ConfigFlag = iota
	OFPC_FRAG_DROP
	OFPC_FRAG_REASM
	OFPC_FRAG_MASK
)

var ConfigFlags = struct {
	FragmentNormal     ConfigFlag
	FragmentDrop       ConfigFlag
	FragmentReassemble ConfigFlag
	FragmentMask       ConfigFlag
}{
	OFPC_FRAG_NORMAL,
	OFPC_FRAG_DROP,
	OFPC_FRAG_REASM,
	OFPC_FRAG_MASK,
}

type FlowMod struct {
	Header
	Match       Match
	Cookie      Cookie
	Command     FlowModCommand
	IdleTimeout uint16
	HardTimeout uint16
	Priority    uint16
	BufferId    BufferId
	OutPort     PortNumber
	Flags       FlowModFlag
	Actions     []Action
}

type Cookie uint64
type FlowModCommand uint16
type BufferId uint32
type FlowModFlag uint16

const (
	OFPFC_ADD FlowModCommand = iota
	OFPFC_MODIFY
	OFPFC_MODIFY_STRICT
	OFPFC_DELETE
	OFPFC_DELETE_STRICT
)

var FlowModCommands = struct {
	Add          FlowModCommand
	Modify       FlowModCommand
	ModifyStrict FlowModCommand
	Delete       FlowModCommand
	DeleteStrict FlowModCommand
}{
	OFPFC_ADD,
	OFPFC_MODIFY,
	OFPFC_MODIFY_STRICT,
	OFPFC_DELETE,
	OFPFC_DELETE_STRICT,
}

const (
	OFPFF_SEND_FLOW_REM FlowModFlag = 1 << iota
	OFPFF_CHECK_OVERLAP
	OFPFF_EMERG
)

var FlowModFlags = struct {
	SendFlowRemoved FlowModFlag
	CheckOverlap    FlowModFlag
	Emergency       FlowModFlag
}{
	OFPFF_SEND_FLOW_REM,
	OFPFF_CHECK_OVERLAP,
	OFPFF_EMERG,
}

type PortMod struct {
	Header
	PortNumber      PortNumber
	HardwareAddress net.HardwareAddr
	Config          PortConfig
	Mask            PortConfig
	Advertise       PortFeature
	pad             [4]uint8
}

type QueueGetConfigRequest struct {
	Header
	Port PortNumber
	pad  [2]uint8
}

type QueueGetConfigReply struct {
	Header
	Port   PortNumber
	pad    [6]uint8
	Queues []PacketQueue
}

type StatsRequest struct {
	Header
	Type  StatsType
	Flags uint16
	Body  []uint8
}

type StatsReply struct {
	Header
	Type  StatsType
	Flags uint16
	Body  []uint8
}

type StatsType uint16

const (
	OFPST_DESC StatsType = iota
	OFPST_FLOW
	OFPST_AGGREGATE
	OFPST_TABLE
	OFPST_PORT
	OFPST_QUEUE
	OFPST_VENDOR StatsType = 0xffff
)

var StatsTypes = struct {
	Description StatsType
	Flow        StatsType
	Aggregate   StatsType
	Table       StatsType
	Port        StatsType
	Queue       StatsType
	Vendor      StatsType
}{
	OFPST_DESC,
	OFPST_FLOW,
	OFPST_AGGREGATE,
	OFPST_TABLE,
	OFPST_PORT,
	OFPST_QUEUE,
	OFPST_VENDOR,
}
