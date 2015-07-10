package of10

import "net"

type ActionType uint16

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
