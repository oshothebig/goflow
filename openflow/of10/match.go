package of10

import "net"

type Match struct {
	Wildcards    Wildcard
	InPort       PortNumber
	EthSrc       net.HardwareAddr
	EthDst       net.HardwareAddr
	VlanId       VlanId
	VlanPriority VlanPriority
	pad1         [1]uint8
	EtherType    EtherType
	IpTos        Dscp
	IpProtocol   ProtocolNumber
	pad2         [2]uint8
	IpSrc        net.IP
	IpDst        net.IP
	TransportSrc TransportPort
	TransportDst TransportPort
}

type Wildcard uint32
type VlanId uint16
type VlanPriority uint8
type EtherType uint16
type Dscp uint8
type ProtocolNumber uint8
type TransportPort uint16

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
