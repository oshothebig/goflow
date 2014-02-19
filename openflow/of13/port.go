package of13

import "net"

const MaxPortNameLength = 16
const EthernetAddressLength = 6

type PortConfig uint32
type PortConfigMask uint32
type PortState uint32
type PortNumber uint32
type PortFeature uint32

type Port struct {
	PortNumber         PortNumber
	pad                [4]uint8
	HardwareAddress    net.HardwareAddr
	pad2               [2]uint8
	Name               [MaxPortNameLength]uint8
	Config             PortConfig
	State              PortState
	CurrentFeatures    PortFeature
	AdvertisedFeatures PortFeature
	SupportedFeatures  PortFeature
	PeerFeatures       PortFeature
	CurentSpeed        uint32
	MaxSpeed           uint32
}

// corresponding to ofp_port_config
const (
	OFPPC_PORT_DOWN    PortConfig = 1 << 0
	OFPPC_NO_RECV      PortConfig = 1 << 2
	OFPPC_NO_FWD       PortConfig = 1 << 5
	OFPPC_NO_PACKET_IN PortConfig = 1 << 6
)

var PortConfigs = struct {
	PortDown   PortConfig
	NoReceive  PortConfig
	NoForward  PortConfig
	NoPacketIn PortConfig
}{
	OFPPC_PORT_DOWN,
	OFPPC_NO_RECV,
	OFPPC_NO_FWD,
	OFPPC_NO_PACKET_IN,
}

// corresponding to ofp_port_state
const (
	OFPPS_LINK_DOWN PortState = 1 << iota
	OFPPS_BLOCKED
	OFPPS_LIVE
)

var PortStates = struct {
	LinkDown PortState
	Blocked  PortState
	Live     PortState
}{
	OFPPS_LINK_DOWN,
	OFPPS_BLOCKED,
	OFPPS_LIVE,
}

// corresponding to ofp_port_no
const (
	OFPP_MAX        PortNumber = 0xffffff00
	OFPP_IN_PORT    PortNumber = 0xfffffff8
	OFPP_TABLE      PortNumber = 0xfffffff9
	OFPP_NORMAL     PortNumber = 0xfffffffa
	OFPP_FLOOD      PortNumber = 0xfffffffb
	OFPP_ALL        PortNumber = 0xfffffffc
	OFPP_CONTROLLER PortNumber = 0xfffffffd
	OFPP_LOCAL      PortNumber = 0xfffffffe
	OFPP_ANY        PortNumber = 0xffffffff
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
	Any        PortNumber
}{
	OFPP_MAX,
	OFPP_IN_PORT,
	OFPP_TABLE,
	OFPP_NORMAL,
	OFPP_FLOOD,
	OFPP_ALL,
	OFPP_CONTROLLER,
	OFPP_LOCAL,
	OFPP_ANY,
}

// corresponding to ofp_port_features
const (
	OFPPF_10MB_HD PortFeature = 1 << iota
	OFPPF_10MB_FD
	OFPPF_100MB_HD
	OFPPF_100MB_FD
	OFPPF_1GB_HD
	OFPPF_1GB_FD
	OFPPF_10GB_FD
	OFPPF_40GB_FD
	OFPPF_100GB_FD
	OFPPF_1TB_FD
	OFPPF_OTHER
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
	FullDuplex40G   PortFeature
	FullDuplex100G  PortFeature
	FullDuplex1T    PortFeature
	OtherRate       PortFeature
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
	OFPPF_40GB_FD,
	OFPPF_100GB_FD,
	OFPPF_1TB_FD,
	OFPPF_OTHER,
	OFPPF_COPPER,
	OFPPF_FIBER,
	OFPPF_AUTONEG,
	OFPPF_PAUSE,
	OFPPF_PAUSE_ASYM,
}
