package of13

import "io"

const WireProtocolVersion = 0x04
const MinimumHeaderLength = 8

type Packetizable interface {
	io.ReadWriter
	Len() uint
}

type Message interface {
	Packetizable
	Header() *Header
}

type Header struct {
	Version uint8
	Type    MessageType
	Length  uint16
	Xid     uint32
}

func (h *Header) Header() *Header {
	return h
}

func NewHeader(typ MessageType) *Header {
	h := new(Header)
	h.Version = WireProtocolVersion
	h.Type = typ
	h.Length = MinimumHeaderLength
	h.Xid = generateXid()

	return h
}

func newXidGenerator() func() uint32 {
	var xid uint32 = 0
	return func() uint32 {
		xid += 1
		return xid
	}
}

var generateXid func() uint32 = newXidGenerator()

type MessageType uint8

const (
	OFPT_HELLO MessageType = iota
	OFPT_ERROR
	OFPT_ECHO_REQUEST
	OFPT_ECHO_REPLY
	OFPT_EXPERIMENTER
	OFPT_FEATURES_REQUEST
	OFPT_FEATURES_REPLY
	OFPT_GET_CONFIG_REQUEST
	OFPT_GET_CONFIG_REPLY
	OFPT_SET_CONFIG
	OFPT_PACKET_IN
	OFPT_FLOW_REMOVED
	OFPT_PORT_STATUS
	OFPT_PACKET_OUT
	OFPT_FLOW_MOD
	OFPT_GROUP_MOD
	OFPT_PORT_MOD
	OFPT_TABLE_MOD
	OFPT_MULTIPART_REQUEST
	OFPT_MULTIPART_REPLY
	OFPT_BARRIER_REQUEST
	OFPT_BARRIER_REPLY
	OFPT_QUEUE_GET_CONFIG_REQUEST
	OFPT_QUEUE_GET_CONFIG_REPLY
	OFPT_ROLE_REQUEST
	OFPT_ROLE_REPLY
	OFPT_GET_ASYNC_REQUEST
	OFPT_GET_ASYNC_REPLY
	OFPT_SET_ASYNC
	OFPT_METER_MOD
)

var MessageTypes = struct {
	Hello                 MessageType
	Error                 MessageType
	EchoRequest           MessageType
	EchoReply             MessageType
	Experimenter          MessageType
	FeaturesRequest       MessageType
	FeaturesReply         MessageType
	GetConfigRequest      MessageType
	GetConfigReply        MessageType
	SetConfig             MessageType
	PacketIn              MessageType
	FlowRemoved           MessageType
	PortStatus            MessageType
	PacketOut             MessageType
	FlowMod               MessageType
	GroupMod              MessageType
	PortMod               MessageType
	TableMod              MessageType
	MultipartRequest      MessageType
	MultipartReply        MessageType
	BarrierRequest        MessageType
	BarrierReply          MessageType
	QueueGetConfigRequest MessageType
	QueueGetConfigReply   MessageType
	RoleRequest           MessageType
	RoleReply             MessageType
	GetAsyncRequest       MessageType
	GetAsyncReply         MessageType
	SetAsync              MessageType
	MeterMod              MessageType
}{
	OFPT_HELLO,
	OFPT_ERROR,
	OFPT_ECHO_REQUEST,
	OFPT_ECHO_REPLY,
	OFPT_EXPERIMENTER,
	OFPT_FEATURES_REQUEST,
	OFPT_FEATURES_REPLY,
	OFPT_GET_CONFIG_REQUEST,
	OFPT_GET_CONFIG_REPLY,
	OFPT_SET_CONFIG,
	OFPT_PACKET_IN,
	OFPT_FLOW_REMOVED,
	OFPT_PORT_STATUS,
	OFPT_PACKET_OUT,
	OFPT_FLOW_MOD,
	OFPT_GROUP_MOD,
	OFPT_PORT_MOD,
	OFPT_TABLE_MOD,
	OFPT_MULTIPART_REQUEST,
	OFPT_MULTIPART_REPLY,
	OFPT_BARRIER_REQUEST,
	OFPT_BARRIER_REPLY,
	OFPT_QUEUE_GET_CONFIG_REQUEST,
	OFPT_QUEUE_GET_CONFIG_REPLY,
	OFPT_ROLE_REQUEST,
	OFPT_ROLE_REPLY,
	OFPT_GET_ASYNC_REQUEST,
	OFPT_GET_ASYNC_REPLY,
	OFPT_SET_ASYNC,
	OFPT_METER_MOD,
}

type Cookie uint64
type CookieMask uint64
type Metadata uint64
type MetadataMask uint64
