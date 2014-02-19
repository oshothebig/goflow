package of10

import "net"

const MaxPortNameLength = 16

type PortNumber uint16
type PortConfig uint16
type PortState uint16
type PortFeature uint16

type PhysicalPort struct {
	PortNumber         PortNumber
	HardwareAddress    net.HardwareAddr
	Name               [MaxPortNameLength]uint8
	Config             PortConfig
	State              PortState
	CurrentFeatures    PortFeature
	AdvertisedFeatures PortFeature
	SupportedFeatures  PortFeature
	PeerFeatures       PortFeature
}

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
