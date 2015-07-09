package of10

import "net"

type ActionType uint16

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

type Action interface {
	Packetizable
	GetType() ActionType
}

type ActionHeader struct {
	Type   ActionType
	Length uint16
}

func NewActionHeader(typ ActionType, length uint16) *ActionHeader {
	return &ActionHeader{typ, length}
}

func (header *ActionHeader) GetType() ActionType {
	return header.Type
}

type SendOutPort struct {
	ActionHeader
	Port      PortNumber
	MaxLength uint16
}

func NewSendOutPort(port PortNumber, maxLength uint16) *SendOutPort {
	return &SendOutPort{
		*NewActionHeader(ActionTypes.Output, 8),
		port,
		maxLength,
	}
}

type Enqueue struct {
	ActionHeader
	Port    PortNumber
	pad     [6]uint8
	QueueId uint32
}

type SetVlanVid struct {
	ActionHeader
	Id  VlanId
	pad [2]uint32
}

type SetVlanPcp struct {
	ActionHeader
	Priority VlanPriority
	pad      [3]uint8
}

type SetEtherAddress struct {
	ActionHeader
	Address net.HardwareAddr
	pad     [6]uint8
}

type SetIpAddress struct {
	ActionHeader
	Address net.IP
}

type SetIpTos struct {
	ActionHeader
	Tos Dscp
	pad [3]uint8
}

type SetTransportPort struct {
	ActionHeader
	Port TransportPort
	pad  [2]uint8
}

type VendorActionHeader struct {
	ActionHeader
	Vendor VendorId
}

type VendorId uint32
