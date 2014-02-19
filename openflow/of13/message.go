package of13

import "net"

type Hello struct {
	Header
	Elements []HelloElement
}

type HelloElement interface {
	Packetizable
	Header() *HelloElemenetHeader
}

type HelloElemenetHeader struct {
	Type   HelloElementType
	Length uint16
}

func (h *HelloElemenetHeader) Header() *HelloElemenetHeader {
	return h
}

type HelloElementVersionBitmap struct {
	Header
	Bitmaps []uint32
}

type HelloElementType uint16

const (
	OFPHET_VERSIONBITMAP HelloElementType = 1
)

type EchoMessage struct {
	Header
	Data []uint8
}

type EchoRequest struct {
	EchoMessage
}

func NewEchoRequest(data []uint8) *EchoRequest {
	m := &EchoMessage{*NewHeader(MessageTypes.EchoRequest), newData(data)}
	m.Length = m.Length + uint16(len(data))

	return &EchoRequest{*m}
}

type EchoReply struct {
	EchoMessage
}

func NewEchoReply(data []uint8) *EchoReply {
	m := &EchoMessage{*NewHeader(MessageTypes.EchoReply), newData(data)}
	m.Length = m.Length + uint16(len(data))

	return &EchoReply{*m}
}

func newData(data []uint8) []uint8 {
	var result []uint8
	if data == nil {
		result = make([]uint8, 0)
	} else {
		result = make([]uint8, len(data))
		copy(result, data)
	}

	return result
}

type FeaturesRequest struct {
	Header
}

func NewFeatureRequest() *FeaturesRequest {
	return &FeaturesRequest{
		*NewHeader(MessageTypes.FeaturesRequest),
	}
}

type Capability uint32

type FeaturesReply struct {
	Header
	DatapathId   uint64
	Buffers      uint32
	Tables       uint8
	AuxiliaryId  uint8
	pad          [2]uint8
	Capabilities Capability
	Reserved     uint32
}

// corresponds to ofp_capabilities
const (
	OFPC_FLOW_STATS   Capability = 1 << 0
	OFPC_TABLE_STATS  Capability = 1 << 1
	OFPC_PORT_STATS   Capability = 1 << 2
	OFPC_GROUP_STATS  Capability = 1 << 3
	OFPC_IP_REASM     Capability = 1 << 5
	OFPC_QUEUE_STATS  Capability = 1 << 6
	OFPC_PORT_BLOCKED Capability = 1 << 8
)

var Capabilities = struct {
	FlowStats    Capability
	TableStats   Capability
	PortStats    Capability
	GroupStats   Capability
	IpReassemble Capability
	QueueStats   Capability
	PortBlocked  Capability
}{
	OFPC_FLOW_STATS,
	OFPC_TABLE_STATS,
	OFPC_PORT_STATS,
	OFPC_GROUP_STATS,
	OFPC_IP_REASM,
	OFPC_QUEUE_STATS,
	OFPC_PORT_BLOCKED,
}

type ConfigFlag uint16

type GetConfigRequest struct {
	Header
}

func NewGetConfigRequest() *GetConfigRequest {
	return &GetConfigRequest{*NewHeader(MessageTypes.GetConfigRequest)}
}

type SwitchConfigMessage struct {
	Header
	Flags          ConfigFlag
	MissSendLength uint16
}

func NewSwitchConfigMessage(typ MessageType, flags ConfigFlag, length uint16) *SwitchConfigMessage {
	m := &SwitchConfigMessage{
		*NewHeader(typ),
		flags,
		length,
	}
	m.Length = m.Length + 4

	return m
}

type GetConfigReply struct {
	SwitchConfigMessage
}

func NewGetConfigReply(flags ConfigFlag, length uint16) *GetConfigReply {
	return &GetConfigReply{*NewSwitchConfigMessage(MessageTypes.GetConfigReply, flags, length)}
}

type SetConfig struct {
	SwitchConfigMessage
}

func NewSetConfig(flags ConfigFlag, length uint16) *SetConfig {
	return &SetConfig{*NewSwitchConfigMessage(MessageTypes.SetConfig, flags, length)}
}

// corresponds to ofp_config_flags
const (
	OFPC_FRAG_NORMAL ConfigFlag = 0
	OFPC_FRAG_DROP   ConfigFlag = 1 << 0
	OFPC_FRAG_REASM  ConfigFlag = 1 << 1
	OFPC_FRAG_MASK   ConfigFlag = 3
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

type PacketIn struct {
	Header
	BufferId    BufferId
	TotalLength uint16
	Reason      PacketInReason
	TableId     uint8
	Cookie      Cookie
	Match       Match
	pad         [2]uint8
	Data        []uint8
}

type PacketInReason uint8

const (
	OFPR_NO_MATCH PacketInReason = iota
	OFPR_ACTION
	OFPR_INVALID_TTL
)

var PacketInReasons = struct {
	NoMatch    PacketInReason
	Action     PacketInReason
	InvalidTtl PacketInReason
}{
	OFPR_NO_MATCH,
	OFPR_ACTION,
	OFPR_INVALID_TTL,
}

type BufferId uint32

const (
	OFP_NO_BUFFER BufferId = 0xffffffff
)

var BufferIds = struct {
	NoBuffer BufferId
}{
	OFP_NO_BUFFER,
}

type FlowRemoved struct {
	Header
	Cookie          Cookie
	Priority        uint16
	Reason          FlowRemovedReason
	TableId         uint8
	DurationSec     uint32
	DurationNanoSec uint32
	IdleTimeout     uint16
	HardTimeout     uint16
	PacketCount     uint64
	ByteCount       uint64
	Match           Match
}

type FlowRemovedReason uint8

const (
	OFPRR_IDLE_TIMEOUT FlowRemovedReason = iota
	OFPRR_HARD_TIMEOUT
	OFPRR_DELETE
	OFPRR_GROUP_DELETE
)

var FlowRemovedReasons = struct {
	IdleTimeout FlowRemovedReason
	HardTimeout FlowRemovedReason
	Delete      FlowRemovedReason
	GroupDelete FlowRemovedReason
}{
	OFPRR_IDLE_TIMEOUT,
	OFPRR_HARD_TIMEOUT,
	OFPRR_DELETE,
	OFPRR_GROUP_DELETE,
}

type PortStatus struct {
	Header
	Reason      PortReason
	pad         [7]uint8
	Description Port
}

type PortReason uint8

const (
	OFPPR_ADD PortReason = iota
	OFPPR_DELETE
	OFPRR_MODIFY
)

var PortReasons = struct {
	Add    PortReason
	Delete PortReason
	Modify PortReason
}{
	OFPPR_ADD,
	OFPPR_DELETE,
	OFPRR_MODIFY,
}

type PacketOut struct {
	Header
	BufferId      BufferId
	InPort        PortNumber
	ActionsLength uint16
	pad           [6]uint8
	Actions       []Action
	Data          []uint8
}

type FlowMod struct {
	Header
	Cookie       Cookie
	CookieMask   CookieMask
	TableId      uint8
	Command      FlowModCommand
	IdleTimeout  uint16
	HardTimeout  uint16
	Priority     uint16
	BufferId     BufferId
	OutPort      PortNumber
	OutGroup     uint32
	Flags        FlowModFlag
	pad          [2]uint8
	Match        Match
	Instructions []Instruction
}

type FlowModCommand uint8
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
	OFPFF_RESET_COUNTS
	OFPFF_NO_PKT_COUNTS
	OFPFF_NO_BYT_COUNTS
)

var FlowModFlags = struct {
	SendFlowRemoved FlowModFlag
	CheckOverlap    FlowModFlag
	ResetCounts     FlowModFlag
	NoPacketCounts  FlowModFlag
	NoByteCounts    FlowModFlag
}{
	OFPFF_SEND_FLOW_REM,
	OFPFF_CHECK_OVERLAP,
	OFPFF_RESET_COUNTS,
	OFPFF_NO_PKT_COUNTS,
	OFPFF_NO_BYT_COUNTS,
}

type GroupModCommand uint16
type GroupType uint8

type GroupMod struct {
	Header
	Command GroupModCommand
	Type    GroupType
	pad     [1]uint8
	GroupId uint32
	Buckets []Bucket
}

const (
	OFPGC_ADD GroupModCommand = iota
	OFPGC_MODIFY
	OFPGC_DELETE
)

var GroupModCommands = struct {
	Add    GroupModCommand
	Modify GroupModCommand
	Delete GroupModCommand
}{
	OFPGC_ADD,
	OFPGC_MODIFY,
	OFPGC_DELETE,
}

const (
	OFPGT_ALL GroupType = iota
	OFPGT_SELECT
	OFPGT_INDIRECT
	OFPGT_FF
)

var GroupTypes = struct {
	All          GroupType
	Select       GroupType
	Indirect     GroupType
	FastFailover GroupType
}{
	OFPGT_ALL,
	OFPGT_SELECT,
	OFPGT_INDIRECT,
	OFPGT_FF,
}

const (
	OFPG_MAX = 0xffffff00
	OFPG_ALL = 0xfffffffc
	OFPG_ANY = 0xffffffff
)

type Bucket struct {
	Length    uint16
	Weight    uint16
	WatchPort uint16
	pad       [4]uint8
	Actions   []Action
}

type PortMod struct {
	Header
	Number       PortNumber
	pad          [4]uint8
	HardwareAddr net.HardwareAddr
	pad2         [2]uint8
	Config       PortConfig
	Mask         PortConfigMask
	Advertise    PortFeature
	pad3         [4]uint8
}

type TableConfig uint32

type TableMod struct {
	Header
	TableId uint8
	pad     [3]uint8
	Config  TableConfig
}

// corresponds to ofp_talbe
const (
	OFPTT_MAX = 0xfe
	OFPTT_ALL = 0xff
)

// corresponds to ofp_table_config
const (
	OFPTC_DEPRECATED_MASK TableConfig = 3
)

type BarrierRequest struct {
	Header
}

type BarrierReply struct {
	Header
}

func NewBarrierRequest() *BarrierRequest {
	m := &BarrierRequest{*NewHeader(MessageTypes.BarrierRequest)}

	return m
}

func NewBarrierReply() *BarrierReply {
	m := &BarrierReply{*NewHeader(MessageTypes.BarrierReply)}

	return m
}

type QueueGetConfigRequest struct {
	Header
	Port PortNumber
	pad  [4]uint8
}

type QueueGetConfigReply struct {
	Header
	Port   PortNumber
	pad    [4]uint8
	Queues []PacketQueue
}

const RoleMessageLength = 24

type ControllerRole uint32

type RoleMessage struct {
	Header
	Role         ControllerRole
	pad          [4]uint8
	GenerationId uint64
}

type RoleRequest struct {
	RoleMessage
}

func NewRoleRequest(r ControllerRole, id uint64) *RoleRequest {
	m := RoleMessage{
		Header:       *NewHeader(MessageTypes.RoleRequest),
		Role:         r,
		GenerationId: id,
	}
	m.Length = RoleMessageLength

	return &RoleRequest{m}
}

type RoleReply struct {
	RoleMessage
}

func NewRoleReply(r ControllerRole, id uint64) *RoleReply {
	m := RoleMessage{
		Header:       *NewHeader(MessageTypes.RoleReply),
		Role:         r,
		GenerationId: id,
	}
	m.Length = RoleMessageLength

	return &RoleReply{m}
}

const (
	OFPCR_ROLE_NOCHANGE ControllerRole = iota
	OFPCR_ROLE_EQUAL
	OFPCR_ROLE_MASTER
	OFPCR_ROLE_SLAVE
)

var ControllerRoles = struct {
	NoChange ControllerRole
	Equal    ControllerRole
	Master   ControllerRole
	Slave    ControllerRole
}{
	OFPCR_ROLE_NOCHANGE,
	OFPCR_ROLE_EQUAL,
	OFPCR_ROLE_MASTER,
	OFPCR_ROLE_SLAVE,
}

type GetAsyncRequest struct {
	Header
}

func NewGetAsyncRequest() *GetAsyncRequest {
	m := &GetAsyncRequest{*NewHeader(MessageTypes.GetAsyncRequest)}

	return m
}

type GetAsyncReply struct {
	AsyncConfig
}

type SetAsync struct {
	AsyncConfig
}

// TODO: read the specification to understand arrays of uint32
type AsyncConfig struct {
	Header
	PacketInMask    [2]uint32
	PortStatusMask  [2]uint32
	FlowRemovedMask [2]uint32
}

type MeterModCommand uint16
type MeterFlag uint16

type MeterMod struct {
	Header
	Command MeterModCommand
	Flags   MeterFlag
	MeterId uint32
	Bands   []MeterBand
}

type MeterBand interface {
	Packetizable
	Header() *MeterBandHeader
}

type MeterBandHeader struct {
	Type      uint16
	Length    uint16
	Rate      uint32
	BurstSize uint32
}

func (h *MeterBandHeader) Header() *MeterBandHeader {
	return h
}

type MeterBandDrop struct {
	MeterBandHeader
	pad [4]uint8
}

type MeterBandDscpRemark struct {
	MeterBandHeader
	PrecedenceLevel uint8
	pad             [3]uint8
}

type MeterBandExperimenter struct {
	MeterBandHeader
	Experimenter uint32
}

const (
	OFPMC_ADD MeterModCommand = iota
	OFPMC_MODIFY
	OFPMC_DELETE
)

var MeterModCommands = struct {
	Add    MeterModCommand
	Modify MeterModCommand
	Delete MeterModCommand
}{
	OFPMC_ADD,
	OFPMC_MODIFY,
	OFPMC_DELETE,
}

const (
	OFPMF_KBPS MeterFlag = 1 << iota
	OFPMF_PKTPS
	OFPMF_BURST
	OFPMF_STATS
)

const (
	OFPM_MAX        = 0xffff0000
	OFPM_SLOWPATH   = 0xfffffffd
	OFPM_CONTROLLER = 0xfffffffe
	OFPM_ALL        = 0xffffffff
)

const (
	OFPMBT_DROP         = 1
	OFPMBT_DSCP_REMARK  = 2
	OFPMBT_EXPERIMENTER = 0xffff
)
