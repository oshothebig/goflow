package of13

type MultipartHeader struct {
	Header
	Type  MultipartType
	Flags MultipartRequestFlag
	pad   [4]uint8
}

type MultipartType uint16

const (
	OFPMP_DESC MultipartType = iota
	OFPMP_FLOW
	OFPMP_AGGREGATE
	OFPMP_TABLE
	OFPMP_PORT_STATS
	OFPMP_QUEUE
	OFPMP_GROUP
	OFPMP_GROUP_DESC
	OFPMP_GROUP_FEATURES
	OFPMP_METER
	OFPMP_METER_CONFIG
	OFPMP_METER_FEATURES
	OFPMP_TABLE_FEATURES
	OFPMP_PORT_DESC
	OFPMP_EXPERIMENTER MultipartType = 0xffff
)

var MultipartTypes = struct {
	Description      MultipartType
	Flow             MultipartType
	Aggregate        MultipartType
	Table            MultipartType
	PortStats        MultipartType
	Queue            MultipartType
	Group            MultipartType
	GroupDescription MultipartType
	GroupFeatures    MultipartType
	Meter            MultipartType
	MeterConfig      MultipartType
	MeterFeatures    MultipartType
	TableFeatures    MultipartType
	PortDescription  MultipartType
	Experimenter     MultipartType
}{
	OFPMP_DESC,
	OFPMP_FLOW,
	OFPMP_AGGREGATE,
	OFPMP_TABLE,
	OFPMP_PORT_STATS,
	OFPMP_QUEUE,
	OFPMP_GROUP,
	OFPMP_GROUP_DESC,
	OFPMP_GROUP_FEATURES,
	OFPMP_METER,
	OFPMP_METER_CONFIG,
	OFPMP_METER_FEATURES,
	OFPMP_TABLE_FEATURES,
	OFPMP_PORT_DESC,
	OFPMP_EXPERIMENTER,
}

type MultipartRequestFlag uint16

const (
	OFPMPF_REQ_MORE MultipartRequestFlag = 1 << 0
)

const (
	DescriptionLength  = 256
	SerialNumberLength = 32
)

type Description struct {
	Manufacturer [DescriptionLength]uint8
	Hardware     [DescriptionLength]uint8
	Software     [DescriptionLength]uint8
	SerialNumber [SerialNumberLength]uint8
	Datapath     [DescriptionLength]uint8
}

type MultipartDescriptionRequest struct {
	MultipartHeader
}

type MultipartDescriptionReply struct {
	MultipartHeader
	Description
}

type FlowStatsRequest struct {
	TableId      uint8
	pad          [3]uint8
	OutPort      PortNumber
	OutGroup     uint32
	pad2         [2]uint8
	Cookie       Cookie
	CookieMask   CookieMask
	Match        Match
	Instructions []Instruction
}

type FlowStats struct {
	Length          uint16
	TableId         uint8
	pad             uint8
	DurationSec     uint32
	DurationNanoSec uint32
	Priority        uint32
	IdleTimeout     uint16
	HardTimeout     uint16
	Flags           FlowModFlag
	pad2            [4]uint8
	Cookie          Cookie
	PacketCount     uint64
	ByteCount       uint64
	Match           Match
	Instructions    []Instruction
}

type MultipartFlowStatsRequest struct {
	MultipartHeader
	FlowStatsRequest
}

type MultipartFlowStatsReply struct {
	MultipartHeader
	FlowStats
}

type AggregateStatsRequest struct {
	TableId    uint8
	pad        [3]uint8
	OutPort    PortNumber
	OutGroup   uint32
	pad2       [4]uint8
	Cookie     Cookie
	CookieMask CookieMask
	Match      Match
}

type AggregateStatsReply struct {
	PacketCount uint64
	ByteCount   uint64
	FlowCount   uint64
	pad         [4]uint8
}

type MultipartAggregateStatsRequest struct {
	MultipartHeader
	AggregateStatsRequest
}

type MultipartAggregateStatsReply struct {
	MultipartHeader
	AggregateStatsReply
}

type TableStats struct {
	TableId      uint8
	pad          [3]uint32
	ActiveCount  uint32
	LookupCount  uint64
	MatchedCount uint64
}

type MultipartTableRequest struct {
	MultipartHeader
}

type MultipartTableReply struct {
	MultipartHeader
	TableStats
}

type MultipartPortStatsRequest struct {
	MultipartHeader
	PortStatsRequest
}

type MultipartPortStatsReply struct {
	MultipartHeader
	Stats []PortStats
}

type PortStatsRequest struct {
	PortNumber PortNumber
	pad        [4]uint8
}

type PortStats struct {
	PortNumber      PortNumber
	pad             [4]uint8
	RxPackets       uint64
	TxPackets       uint64
	RxBytes         uint64
	TxBytes         uint64
	RxDropped       uint64
	TxDropped       uint64
	RxErrors        uint64
	TxErrors        uint64
	RxFrameErrors   uint64
	TxFrameErrors   uint64
	Collisions      uint64
	DurationSec     uint32
	DurationNanoSec uint32
}

type MultipartQueueStatsRequest struct {
	MultipartHeader
	QueueStatsRequest
}

type MultipartQueueStatsReply struct {
	MultipartHeader
	Stats []QueueStats
}

type QueueStatsRequest struct {
	PortNumber PortNumber
	QueueId    uint32
}

type QueueStats struct {
	PortNumber      PortNumber
	QueueId         uint32
	TxBytes         uint64
	TxPackets       uint64
	TxErrors        uint64
	DurationSec     uint32
	DurationNanoSec uint32
}

type MultipartGroupStatsRequest struct {
	MultipartHeader
	GroupStatsRequest
}

type MultipartGroupStatsReply struct {
	MultipartHeader
	Stats []GroupStats
}

type GroupStatsRequest struct {
	GroupId uint32
	pad     [4]uint8
}

type GroupStats struct {
	Length          uint16
	pad             [2]uint8
	GroupId         uint32
	RefCount        uint32
	pad2            [4]uint8
	PacketCount     uint64
	ByteCount       uint64
	DurationSec     uint32
	DurationNanoSec uint32
	BucketStats     []BucketCounter
}

type BucketCounter struct {
	PacketCount uint64
	ByteCount   uint64
}

type MultipartGroupDescriptionRequest struct {
	MultipartHeader
}

type MultipartGroupDescriptionReply struct {
	MultipartHeader
	Descriptions []GroupDescription
}

type GroupDescription struct {
	Length  uint16
	Type    GroupType
	pad     uint8
	GroupId uint32
	Buckets []Bucket
}

type MultipartGroupFeaturesRequest struct {
	MultipartHeader
}

type MultipartGroupFeaturesReply struct {
	MultipartHeader
	GroupFeatures
}

type GroupFeatures struct {
	Types        GroupType
	Capabilities GroupCapability
	MaxGroups    [4]uint32
	Actions      [4]uint32
}

type GroupCapability uint32

const (
	OFPGFC_SELECT_WEIGHT GroupCapability = 1 << iota
	OFPGFC_SELECT_LIVENESS
	OFPGFC_CHAINING
	OFPGFC_CHAINING_CHECK
)

var GroupCapabilities = struct {
	SelectWeight   GroupCapability
	SelectLiveness GroupCapability
	Chaining       GroupCapability
	ChainingCheck  GroupCapability
}{
	OFPGFC_SELECT_WEIGHT,
	OFPGFC_SELECT_LIVENESS,
	OFPGFC_CHAINING,
	OFPGFC_CHAINING_CHECK,
}

type MultipartMeterStatsRequest struct {
	MultipartHeader
	MeterRequest
}

type MultipartMeterStatsReply struct {
	MultipartHeader
	Stats []MeterStats
}

type MeterRequest struct {
	MeterId uint32
	pad     [4]uint8
}

type MeterStats struct {
	MeterId         uint32
	Length          uint16
	pad             [6]uint8
	FlowCount       uint32
	PacketInCount   uint64
	ByteInCount     uint64
	DurationSec     uint32
	DurationNanoSec uint32
	BandStats       []MeterBandStats
}

type MeterBandStats struct {
	PacketBandCount uint64
	ByteBandCount   uint64
}

type MultipartMeterConfigRequest struct {
	MultipartHeader
	MeterRequest
}

type MultipartMeterConfigReply struct {
	MultipartHeader
	Configs []MeterConfig
}

type MeterConfig struct {
	Length  uint16
	Flags   MeterModCommand
	MeterId uint32
	Bands   []MeterBand
}

type MultipartMeterFeaturesRequest struct {
	MultipartHeader
}

type MultipartMeterFeaturesReply struct {
	MultipartHeader
	MeterFeatures
}

type MeterFeatures struct {
	MaxMeter     uint32
	BandTypes    uint32
	Capabilities uint32
	MaxBands     uint8
	MaxColor     uint8
	pad          [2]uint8
}

const MaxTableNameLength = 32

type TableFeatures struct {
	Length        uint16
	TableId       uint8
	pad           [5]uint8
	Name          [MaxTableNameLength]uint8
	MetadataMatch Metadata
	MetadataWrite Metadata
	Config        TableConfig
	MaxEntries    uint32
	Properties    []TableFeatureProperty
}

type TableFeatureProperty interface {
	Packetizable
	Header() *TableFeatureProperty
}

type TableFeaturePropertyType uint16

const (
	OFPTFPT_INSTRUCTIONS        TableFeaturePropertyType = 0
	OFPTFPT_INSTRUCTIONS_MISS   TableFeaturePropertyType = 1
	OFPTFPT_NEXT_TABLES         TableFeaturePropertyType = 2
	OFPTFPT_NEXT_TABLES_MISS    TableFeaturePropertyType = 3
	OFPTFPT_WRITE_ACTIONS       TableFeaturePropertyType = 4
	OFPTFPT_WRITE_ACTIONS_MISS  TableFeaturePropertyType = 5
	OFPTFPT_APPLY_ACTIONS       TableFeaturePropertyType = 6
	OFPTFPT_APPLY_ACTIONS_MISS  TableFeaturePropertyType = 7
	OFPTFPT_MATCH               TableFeaturePropertyType = 8
	OFPTFPT_WILDCARD            TableFeaturePropertyType = 10
	OFPTFPT_WRITE_SETFIELD      TableFeaturePropertyType = 12
	OFPTFPT_WRITE_SETFIELD_MISS TableFeaturePropertyType = 13
	OFPTFPT_APPLY_SETFIELD      TableFeaturePropertyType = 14
	OFPTFPT_APPLY_SETFIELD_MISS TableFeaturePropertyType = 15
	OFPTFPT_EXPERIMENTER        TableFeaturePropertyType = 0xfffe
	OFPTFPT_EXPERIMENTER_MISS   TableFeaturePropertyType = 0xffff
)

var TableFeaturePropertyTypes = struct {
	Instructions      TableFeaturePropertyType
	InstructionsMiss  TableFeaturePropertyType
	NextTable         TableFeaturePropertyType
	NextTableMiss     TableFeaturePropertyType
	WriteActions      TableFeaturePropertyType
	WriteActionsMiss  TableFeaturePropertyType
	ApplyActions      TableFeaturePropertyType
	ApplyActionsMiss  TableFeaturePropertyType
	Match             TableFeaturePropertyType
	Wildcard          TableFeaturePropertyType
	WriteSetField     TableFeaturePropertyType
	WriteSetFieldMiss TableFeaturePropertyType
	ApplySetField     TableFeaturePropertyType
	ApplySetFieldMiss TableFeaturePropertyType
	Experimenter      TableFeaturePropertyType
	ExperimenterMiss  TableFeaturePropertyType
}{
	OFPTFPT_INSTRUCTIONS,
	OFPTFPT_INSTRUCTIONS_MISS,
	OFPTFPT_NEXT_TABLES,
	OFPTFPT_NEXT_TABLES_MISS,
	OFPTFPT_WRITE_ACTIONS,
	OFPTFPT_WRITE_ACTIONS_MISS,
	OFPTFPT_APPLY_ACTIONS,
	OFPTFPT_APPLY_ACTIONS_MISS,
	OFPTFPT_MATCH,
	OFPTFPT_WILDCARD,
	OFPTFPT_WRITE_SETFIELD,
	OFPTFPT_WRITE_SETFIELD_MISS,
	OFPTFPT_APPLY_SETFIELD,
	OFPTFPT_APPLY_SETFIELD_MISS,
	OFPTFPT_EXPERIMENTER,
	OFPTFPT_EXPERIMENTER_MISS,
}

type TableFeaturePropertyHeader struct {
	Type   TableFeaturePropertyType
	Length uint16
}

func (h *TableFeaturePropertyHeader) Header() *TableFeaturePropertyHeader {
	return h
}

type TableFeaturePropertyInstructions struct {
	TableFeaturePropertyHeader
	Instructions []Instruction
}

type TableFeaturePropertyNextTable struct {
	TableFeaturePropertyHeader
	NextTableIds []uint8
}

type TableFeaturePropertyActions struct {
	TableFeaturePropertyHeader
	Actions []Action
}

type TableFeaturePropertyOxm struct {
	TableFeaturePropertyHeader
	OxmIds []OxmHeader
}

type TableFeaturePropertyExperimenter struct {
	TableFeaturePropertyHeader
	Experimeter      uint32
	ExperimenterType uint32
	Data             []uint32
}

type MultipartTableFeaturesRequest struct {
	MultipartHeader
	Features []TableFeatures
}

type MultipartTableFeaturesReply struct {
	MultipartHeader
	Features []TableFeatures
}

type MultipartPortDescriptionRequest struct {
	MultipartHeader
}

type MultipartPortDescriptonReply struct {
	MultipartHeader
	Ports []Port
}

type MultipartExperimenterHeader struct {
	MultipartHeader
	Experimenter     uint32
	ExperimenterType uint32
}
