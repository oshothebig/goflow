package of10

import "bytes"

type FeaturesRequest struct {
	Header
}

func (m *FeaturesRequest) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
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

func (m *FeaturesReply) UnmarshalBinary(data []byte) error {
	size := len(data)
	fields := []interface{}{&m.Header, &m.DatapathId, &m.Buffers, &m.Tables, &m.pad, &m.Capabilities, &m.Actions}
	reader := bytes.NewReader(data)
	if err := unmarshalFields(reader, fields...); err != nil {
		return err
	}
	read := size - reader.Len()
	ports, err := readPhysicalPort(data[read:])
	if err != nil {
		return err
	}
	m.Ports = ports
	return nil
}

type DatapathId uint64
type Capability uint32

type GetConfigRequest struct {
	Header
}

func (m *GetConfigRequest) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type SwitchConfig struct {
	Header
	Flags          ConfigFlag
	MissSendLength uint16
}

type GetConfigReply struct {
	Header
	Flags          ConfigFlag
	MissSendLength uint16
}

func (m *GetConfigReply) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type SetConfig struct {
	Header
	Flags          ConfigFlag
	MissSendLength uint16
}

func (m *SetConfig) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type ConfigFlag uint16

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

func (m *FlowMod) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	fields := []interface{}{
		&m.Header, &m.Match, &m.Cookie, &m.Command, &m.IdleTimeout, &m.HardTimeout,
		&m.Priority, &m.BufferId, &m.OutPort, &m.Flags,
	}
	if err := unmarshalFields(reader, fields...); err != nil {
		return err
	}
	m.Actions = readActions(reader, reader.Len())
	return nil
}

type Cookie uint64
type FlowModCommand uint16
type BufferId uint32
type FlowModFlag uint16

type PortMod struct {
	Header
	PortNumber      PortNumber
	HardwareAddress [EthernetAddressLength]uint8
	Config          PortConfig
	Mask            PortConfig
	Advertise       PortFeature
	pad             [4]uint8
}

func (m *PortMod) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type QueueGetConfigRequest struct {
	Header
	Port PortNumber
	pad  [2]uint8
}

func (m *QueueGetConfigRequest) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type QueueGetConfigReply struct {
	Header
	Port   PortNumber
	pad    [6]uint8
	Queues []PacketQueue
}

func (m *QueueGetConfigReply) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	if err := unmarshalFields(reader, &m.Header, &m.Port, &m.pad); err != nil {
		return err
	}
	m.Queues = readPacketQueues(reader)
	return nil
}

type StatsRequest struct {
	Header
	Type  StatsType
	Flags uint16
	Body  []uint8
}

func (m *StatsRequest) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	fields := []interface{}{&m.Header, &m.Type, &m.Flags}
	if err := unmarshalFields(reader, fields...); err != nil {
		return err
	}
	m.Body = data[len(data)-reader.Len():]
	return nil
}

type StatsReply struct {
	Header
	Type  StatsType
	Flags uint16
	Body  []uint8
}

func (m *StatsReply) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	fields := []interface{}{&m.Header, &m.Type, &m.Flags}
	if err := unmarshalFields(reader, fields...); err != nil {
		return err
	}
	m.Body = data[len(data)-reader.Len():]
	return nil
}

type StatsType uint16

type DescriptionStats struct {
	Manufacturer []uint8
	Hardware     []uint8
	Software     []uint8
	SerialNumber []uint8
	Datapath     []uint8
}

type FlowStatsRequest struct {
	Match   Match
	TalbeId uint8
	pad     uint8
	OutPort PortNumber
}

type FlowStatsReply struct {
	Length          uint16
	TableId         uint8
	pad             [1]uint8
	Match           Match
	DurationSec     uint32
	DurationNanoSec uint32
	Priority        uint16
	IdleTimeout     uint16
	HardTimeout     uint16
	pad2            [6]uint8
	Cookie          Cookie
	PacketCount     uint64
	ByteCount       uint64
	Actions         []ActionHeader
}

type AggregateStatsRequest struct {
	Match   Match
	TableId uint8
	pad     [1]uint8
	OutPort PortNumber
}

type AggregateStatsReply struct {
	PacketCount uint64
	ByteCount   uint64
	FlowCount   uint64
	pad         [4]uint8
}

type TableStatsReply struct {
	TableId      uint8
	pad          [3]uint8
	Name         []uint8
	Wildcards    Wildcard
	MaxEntries   uint32
	ActiveCount  uint32
	LookupCount  uint32
	MatchedCount uint32
}

type PortStatsRequest struct {
	PortNumber PortNumber
	pad        [6]uint8
}

type PortStatsReply struct {
	PortNumber      PortNumber
	pad             [6]uint8
	RxPackets       uint64
	TxPackets       uint64
	RxBytes         uint64
	TxBytes         uint64
	RxDropped       uint64
	TxDropped       uint64
	RxErrors        uint64
	TxErrors        uint64
	RxFrameErrors   uint64
	RxOverrunErrors uint64
	RxCrcErrors     uint64
	Collisions      uint64
}

type QueueStatsRequest struct {
	PortNumber PortNumber
	pad        [2]uint8
	QueueId    uint32
}

type QueueStatsReply struct {
	PortNumber PortNumber
	pad        [2]uint8
	QueueId    uint32
	TxBytes    uint64
	TxPackets  uint64
	TxErrors   uint64
}

type PacketOut struct {
	Header
	BufferId      uint32
	InPort        PortNumber
	ActionsLength uint16
	Actions       []Action
	Data          []uint8
}

func (m *PacketOut) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	fields := []interface{}{&m.Header, &m.BufferId, &m.InPort, &m.ActionsLength}
	if err := unmarshalFields(reader, fields...); err != nil {
		return err
	}
	m.Actions = readActions(reader, int(m.ActionsLength))
	return nil
}

type BarrierRequest struct {
	Header
}

func (m *BarrierRequest) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type BarrierReply struct {
	Header
}

func (m *BarrierReply) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type PacketIn struct {
	Header
	BufferId    uint32
	TotalLength uint16
	InPort      PortNumber
	Reason      PacketInReason
	pad         [1]uint8
	Data        []uint8
}

func (m *PacketIn) UnmarshalBinary(data []byte) error {
	fields := []interface{}{&m.Header, &m.BufferId, &m.TotalLength, &m.InPort, &m.Reason, &m.pad}
	reader := bytes.NewReader(data)
	if err := unmarshalFields(reader, fields...); err != nil {
		return err
	}
	read := len(data) - reader.Len()
	m.Data = data[read:]
	return nil
}

type PacketInReason uint8

type FlowRemoved struct {
	Header
	Match           Match
	Cookie          Cookie
	Priority        uint16
	Reason          FlowRemovedReason
	_               [1]uint8
	DurationSec     uint32
	DurationNanoSec uint32
	IdleTimeout     uint16
	_               [2]uint8
	PacketCount     uint64
	ByteCount       uint64
}

func (m *FlowRemoved) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type FlowRemovedReason uint8

type PortStatus struct {
	Header
	Reason      PortStatusReason
	_           [7]uint8
	Description PhysicalPort
}

func (m *PortStatus) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type PortStatusReason uint8

type ErrorMessage struct {
	Header
	Type ErrorType
	Code ErrorCode
	Data []uint8
}

func (m *ErrorMessage) UnmarshalBinary(data []byte) error {
	if err := unmarshalFields(bytes.NewReader(data), &m.Header, &m.Type); err != nil {
		return err
	}
	const fixedLength = HeaderLength + 4
	m.Data = data[fixedLength:]
	return nil
}

type ErrorType uint16

type ErrorCode uint16

var HelloFailedCodes = struct {
	Incompatible    ErrorCode
	PermissionError ErrorCode
}{
	OFPHFC_INCOMPATIBLE,
	OFPHFC_EPERM,
}

type Hello struct {
	Header
}

func (m *Hello) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type EchoRequest struct {
	Header
	Body []uint8
}

func (m *EchoRequest) UnmarshalBinary(data []byte) error {
	if err := unmarshalFields(bytes.NewReader(data), &m.Header); err != nil {
		return err
	}
	const fixedLength = HeaderLength
	m.Body = data[fixedLength:]
	return nil
}

type EchoReply struct {
	Header
	Body []uint8
}

func (m *EchoReply) UnmarshalBinary(data []byte) error {
	if err := unmarshalFields(bytes.NewReader(data), &m.Header); err != nil {
		return err
	}
	const fixedLength = HeaderLength
	m.Body = data[fixedLength:]
	return nil
}

type VendorMessage struct {
	Header
	Vendor VendorId
	Body   []uint8
}

func (m *VendorMessage) UnmarshalBinary(data []byte) error {
	if err := unmarshalFields(bytes.NewBuffer(data), &m.Header, &m.Vendor); err != nil {
		return err
	}
	const fixedLength = HeaderLength + 4
	m.Body = data[fixedLength:]
	return nil
}
