package of10

import "net"

type Match struct {
	Wildcards    uint32
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
	NetworkSrc   NetworkPort
	NetworkDst   NetworkPort
}

type VlanId uint16
type VlanPriority uint8
type EtherType uint16
type Dscp uint8
type ProtocolNumber uint8
type NetworkPort uint16
