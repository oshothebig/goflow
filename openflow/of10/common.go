package of10

import . "github.com/oshothebig/goflow/openflow"

const WireProtocolVersion = 0x01

func NewHeader(typ MessageType) *Header {
	h := new(Header)
	h.Version = WireProtocolVersion
	h.Type = typ
	h.Length = MinimumHeaderLength
	h.Xid = generateXid()

	return h
}

var generateXid func() uint32 = NewXidGenerator()

const (
	// Immutable messages
	OFPT_HELLO = iota
	OFPT_ERROR
	OFPT_ECHO_REQUEST
	OFPT_ECHO_REPLY
	OFPT_VENDOR

	// Switch configuration messages
	OFPT_FEATURES_REQUEST
	OFPT_FEATURES_REPLY
	OFPT_GET_CONFIG_REQUEST
	OFPT_GET_CONFIG_REPLY
	OFPT_SET_CONFIG

	// Asynchronous messages
	OFPT_PACKET_IN
	OFPT_FLOW_REMOVED
	OFPT_PORT_STATUS

	// Controller command messages
	OFPT_PACKET_OUT
	OFPT_FLOW_MOD
	OFPT_PORT_MOD

	// Statistics messages
	OFPT_STATS_REQUEST
	OFPT_STATS_REPLY

	// Barrier messages
	OFPT_BARRIER_REQUEST
	OFPT_BARRIER_REPLY

	// Queue configuration messages
	OFPT_QUEUE_GET_CONFIG_REQUEST
	OFPT_QUEUE_GET_CONFIG_REPLY
)

var MessageTypes = struct {
	Hello                 MessageType
	Error                 MessageType
	EchoRequest           MessageType
	EchoReply             MessageType
	Vendor                MessageType
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
	PortMod               MessageType
	StatsRequest          MessageType
	StatsReply            MessageType
	BarrierRequest        MessageType
	BarrierReply          MessageType
	QueueGetConfigRequest MessageType
	QueueGetConfigReply   MessageType
}{
	// Immutable messages
	OFPT_HELLO,
	OFPT_ERROR,
	OFPT_ECHO_REQUEST,
	OFPT_ECHO_REPLY,
	OFPT_VENDOR,

	// Switch configuration messages
	OFPT_FEATURES_REQUEST,
	OFPT_FEATURES_REPLY,
	OFPT_GET_CONFIG_REQUEST,
	OFPT_GET_CONFIG_REPLY,
	OFPT_SET_CONFIG,

	// Asynchronous messages
	OFPT_PACKET_IN,
	OFPT_FLOW_REMOVED,
	OFPT_PORT_STATUS,

	// Controller command messages
	OFPT_PACKET_OUT,
	OFPT_FLOW_MOD,
	OFPT_PORT_MOD,

	// Statistics messages
	OFPT_STATS_REQUEST,
	OFPT_STATS_REPLY,

	// Barrier messages
	OFPT_BARRIER_REQUEST,
	OFPT_BARRIER_REPLY,

	// Queue configuration messages
	OFPT_QUEUE_GET_CONFIG_REQUEST,
	OFPT_QUEUE_GET_CONFIG_REPLY,
}
