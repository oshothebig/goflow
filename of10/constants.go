package of10

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

const (
	OFPAT_OUTPUT ActionType = iota
	OFPAT_SET_VLAN_VID
	OFPAT_SET_VLAN_PCP
	OFPAT_SET_STRIP_VLAN
	OFPAT_SET_DL_SRC
	OFPAT_SET_DL_DST
	OFPAT_SET_NW_SRC
	OFPAT_SET_NW_DST
	OFPAT_SET_NW_TOS
	OFPAT_SET_TP_SRC
	OFPAT_SET_TP_DST
	OFPAT_ENQUEUE
	OFPAT_VENDOR ActionType = 0xffff
)

var ActionTypes = struct {
	Output        ActionType
	SetVlanId     ActionType
	SetVlanPcp    ActionType
	StripVlan     ActionType
	SetEtherSrc   ActionType
	SetEtherDst   ActionType
	SetIpSrc      ActionType
	SetIpDst      ActionType
	SetIpTos      ActionType
	SetNetworkSrc ActionType
	SetNetworkDst ActionType
	Enqueue       ActionType
	Vendor        ActionType
}{
	OFPAT_OUTPUT,
	OFPAT_SET_VLAN_VID,
	OFPAT_SET_VLAN_PCP,
	OFPAT_SET_STRIP_VLAN,
	OFPAT_SET_DL_SRC,
	OFPAT_SET_DL_DST,
	OFPAT_SET_NW_SRC,
	OFPAT_SET_NW_DST,
	OFPAT_SET_NW_TOS,
	OFPAT_SET_TP_SRC,
	OFPAT_SET_TP_DST,
	OFPAT_ENQUEUE,
	OFPAT_VENDOR,
}

const (
	OFPFW_IN_PORT Wildcard = 1 << iota
	OFPFW_DL_VLAN
	OFPFW_DL_SRC
	OFPFW_DL_DST
	OFPFW_DL_TYPE
	OFPFW_NW_PROTO
	OFPFW_TP_SRC
	OFPFW_TP_DST

	OFPFW_NW_SRC_SHIFT Wildcard = 8
	OFPFW_NW_SRC_BITS  Wildcard = 6
	OFPFW_NW_SRC_MASK  Wildcard = ((1 << OFPFW_NW_SRC_BITS) - 1) << OFPFW_NW_SRC_SHIFT

	OFPFW_NW_DST_SHIFT Wildcard = 16
	OFPFW_NW_DST_BITS  Wildcard = 6
	OFPFW_NW_DST_MASK  Wildcard = ((1 << OFPFW_NW_DST_BITS) - 1) << OFPFW_NW_DST_SHIFT
	OFPFW_NW_DST_ALL   Wildcard = 32 << OFPFW_NW_DST_SHIFT

	OFPFW_DL_VLAN_PCP Wildcard = 1 << 20
	OFPFW_NW_TOS      Wildcard = 1 << 21

	OFPFW_ALL Wildcard = ((1 << 22) - 1)
)

const (
	OFPPC_PORT_DOWN PortConfig = 1 << iota
	OFPPC_NO_STP
	OFPPC_NO_RECV
	OFPPC_NO_RECV_STP
	OFPPC_NO_FLOOD
	OFPPC_NO_FWD
	OFPPC_NO_PACKET_IN
)

var PortConfigs = struct {
	PortDown     PortConfig
	NoStp        PortConfig
	NoReceive    PortConfig
	NoReceiveStp PortConfig
	NoFlood      PortConfig
	NoForward    PortConfig
	NoPacketIn   PortConfig
}{
	OFPPC_PORT_DOWN,
	OFPPC_NO_STP,
	OFPPC_NO_RECV,
	OFPPC_NO_RECV_STP,
	OFPPC_NO_FLOOD,
	OFPPC_NO_FWD,
	OFPPC_NO_PACKET_IN,
}

const (
	OFPPS_LINK_DOWN   PortState = 1 << 0
	OFPPS_STP_LISTEN  PortState = 0 << 8
	OFPPS_STP_LEARN   PortState = 1 << 8
	OFPPS_STP_FORWARD PortState = 2 << 8
	OFPPS_STP_BLOCK   PortState = 3 << 8
	OFPPS_STP_MASK    PortState = 3 << 8
)

var PortStates = struct {
	LinkDown   PortState
	StpListen  PortState
	StpLearn   PortState
	StpForward PortState
	StpBlock   PortState
	StpMask    PortState
}{
	OFPPS_LINK_DOWN,
	OFPPS_STP_LISTEN,
	OFPPS_STP_LEARN,
	OFPPS_STP_FORWARD,
	OFPPS_STP_BLOCK,
	OFPPS_STP_MASK,
}

const (
	OFPP_MAX        PortNumber = 0xff00
	OFPP_IN_PORT    PortNumber = 0xfff8
	OFPP_TABLE      PortNumber = 0xfff9
	OFPP_NORMAL     PortNumber = 0xfffa
	OFPP_FLOOD      PortNumber = 0xfffb
	OFPP_ALL        PortNumber = 0xfffc
	OFPP_CONTROLLER PortNumber = 0xfffd
	OFPP_LOCAL      PortNumber = 0xfffe
	OFPP_NONE       PortNumber = 0xffff
)

var PortNumbers = struct {
	Max        PortNumber
	InPort     PortNumber
	Table      PortNumber
	Normal     PortNumber
	Flood      PortNumber
	All        PortNumber
	Controller PortNumber
	Local      PortNumber
	None       PortNumber
}{
	OFPP_MAX,
	OFPP_IN_PORT,
	OFPP_TABLE,
	OFPP_NORMAL,
	OFPP_FLOOD,
	OFPP_ALL,
	OFPP_CONTROLLER,
	OFPP_LOCAL,
	OFPP_NONE,
}

const (
	OFPPF_10MB_HD PortFeature = 1 << iota
	OFPPF_10MB_FD
	OFPPF_100MB_HD
	OFPPF_100MB_FD
	OFPPF_1GB_HD
	OFPPF_1GB_FD
	OFPPF_10GB_FD
	OFPPF_COPPER
	OFPPF_FIBER
	OFPPF_AUTONEG
	OFPPF_PAUSE
	OFPPF_PAUSE_ASYM
)

var PortFeatures = struct {
	HalfDuplex10M   PortFeature
	FullDuplex10M   PortFeature
	HalfDuplex100M  PortFeature
	FullDuplex100M  PortFeature
	HalfDuplex1G    PortFeature
	FullDuplex1G    PortFeature
	FullDuplex10G   PortFeature
	Copper          PortFeature
	Fiber           PortFeature
	AutoNegotiation PortFeature
	Pause           PortFeature
	AsymmetricPause PortFeature
}{
	OFPPF_10MB_HD,
	OFPPF_10MB_FD,
	OFPPF_100MB_HD,
	OFPPF_100MB_FD,
	OFPPF_1GB_HD,
	OFPPF_1GB_FD,
	OFPPF_10GB_FD,
	OFPPF_COPPER,
	OFPPF_FIBER,
	OFPPF_AUTONEG,
	OFPPF_PAUSE,
	OFPPF_PAUSE_ASYM,
}

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
