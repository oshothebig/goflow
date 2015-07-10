package of10

type Match struct {
	Wildcards    Wildcard
	InPort       PortNumber
	EthSrc       [EthernetAddressLength]uint8
	EthDst       [EthernetAddressLength]uint8
	VlanId       VlanId
	VlanPriority VlanPriority
	pad1         [1]uint8
	EtherType    EtherType
	IpTos        Dscp
	IpProtocol   ProtocolNumber
	pad2         [2]uint8
	IpSrc        [4]uint8
	IpDst        [4]uint8
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
