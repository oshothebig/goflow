package of10

import (
	"bytes"
	"encoding/binary"
)

type FeaturesRequest struct {
	Header
}

func (m *FeaturesRequest) FillBody(body []byte) error {
	return nil
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

func (m *FeaturesReply) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.DatapathId); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Buffers); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Tables); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.pad); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Capabilities); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Actions); err != nil {
		return err
	}
	portsBytes := buf.Bytes()
	ports, err := readPhysicalPort(portsBytes)
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

func (m *GetConfigRequest) FillBody(body []byte) error {
	return nil
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

func (m *GetConfigReply) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Flags); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.MissSendLength); err != nil {
		return err
	}
	return nil
}

type SetConfig struct {
	Header
	Flags          ConfigFlag
	MissSendLength uint16
}

func (m *SetConfig) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Flags); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.MissSendLength); err != nil {
		return err
	}
	return nil
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

func (m *FlowMod) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Match); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Cookie); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Command); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.IdleTimeout); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.HardTimeout); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Priority); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.BufferId); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.OutPort); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Flags); err != nil {
		return err
	}
	m.Actions = readActions(buf)
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

func (m *PortMod) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.PortNumber); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.HardwareAddress); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Config); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Mask); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Advertise); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.pad); err != nil {
		return err
	}
	return nil
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

func (m *StatsRequest) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Type); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Flags); err != nil {
		return err
	}
	m.Body = buf.Bytes()
	return nil
}

type StatsReply struct {
	Header
	Type  StatsType
	Flags uint16
	Body  []uint8
}

func (m *StatsReply) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Type); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Flags); err != nil {
		return err
	}
	m.Body = buf.Bytes()
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
	Actions       []ActionHeader
	Data          []uint8
}

type BarrierRequest struct {
	Header
}

func (m *BarrierRequest) FillBody(body []byte) error {
	return nil
}

type BarrierReply struct {
	Header
}

func (m *BarrierReply) FillBody(body []byte) error {
	return nil
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

func (m *PacketIn) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.BufferId); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.TotalLength); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.InPort); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Reason); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.pad); err != nil {
		return err
	}
	m.Data = buf.Bytes()
	return nil
}

type PacketInReason uint8

type FlowRemoved struct {
	Header
	Match           Match
	Cookie          Cookie
	Priority        uint16
	Reason          FlowRemovedReason
	pad             [1]uint8
	DurationSec     uint32
	DurationNanoSec uint32
	IdleTimeout     uint16
	pad2            [2]uint8
	PacketCount     uint64
	ByteCount       uint64
}

func (m *FlowRemoved) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Match); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Cookie); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Priority); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Reason); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.pad); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.DurationSec); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.DurationNanoSec); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.IdleTimeout); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.pad2); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.PacketCount); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.ByteCount); err != nil {
		return err
	}
	return nil
}

type FlowRemovedReason uint8

type PortStatus struct {
	Header
	Reason      PortStatusReason
	pad         [7]uint8
	Description PhysicalPort
}

func (m *PortStatus) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Reason); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.pad); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Description); err != nil {
		return err
	}
	return nil
}

type PortStatusReason uint8

type ErrorMessage struct {
	Header
	Type ErrorType
	Code ErrorCode
	Data []uint8
}

func (m *ErrorMessage) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	if err := binary.Read(buf, binary.BigEndian, &m.Type); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &m.Code); err != nil {
		return err
	}
	m.Data = buf.Bytes()
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

func (m *Hello) FillBody(body []byte) error {
	return nil
}

type EchoRequest struct {
	Header
	Body []uint8
}

func (m *EchoRequest) FillBody(body []byte) error {
	m.Body = body
	return nil
}

type EchoReply struct {
	Header
	Body []uint8
}

func (m *EchoReply) FillBody(body []byte) error {
	m.Body = body
	return nil
}

type VendorMessage struct {
	Header
	Vendor VendorId
	Body   []uint8
}

func (m *VendorMessage) FillBody(body []byte) error {
	buf := bytes.NewBuffer(body)
	err := binary.Read(buf, binary.BigEndian, &m.Vendor)
	if err != nil {
		return err
	}
	m.Body = buf.Bytes()
	return nil
}
