package of10

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

var (
	errUnsupportedAction error = errors.New("Unsupported action type")
)

type ActionType uint16

const actionHeaderLength = 4

type Action interface {
	GetType() ActionType
}

func readAction(buf *bytes.Buffer) (Action, error) {
	var header ActionHeader
	if err := binary.Read(buf, binary.BigEndian, &header); err != nil {
		return nil, err
	}
	action := newAction(&header)
	if action == nil {
		return nil, errUnsupportedAction
	}
	return action, nil
}

func newAction(h *ActionHeader) Action {
	switch h.Type {
	case ActionTypes.Output:
		return &SendOutPort{ActionHeader: *h}
	case ActionTypes.SetVlanId:
		return &SetVlanVid{ActionHeader: *h}
	case ActionTypes.SetVlanPcp:
		return &SetVlanPcp{ActionHeader: *h}
	case ActionTypes.StripVlan:
		return &StripVlan{ActionHeader: *h}
	case ActionTypes.SetEtherSrc:
		return &SetEtherSrc{ActionHeader: *h}
	case ActionTypes.SetEtherDst:
		return &SetEtherDst{ActionHeader: *h}
	case ActionTypes.SetIpSrc:
		return &SetIpSrc{ActionHeader: *h}
	case ActionTypes.SetIpDst:
		return &SetIpDst{ActionHeader: *h}
	case ActionTypes.SetIpTos:
		return &SetIpTos{ActionHeader: *h}
	case ActionTypes.SetNetworkSrc:
		return &SetTransportSrc{ActionHeader: *h}
	case ActionTypes.SetNetworkDst:
		return &SetTransportDst{ActionHeader: *h}
	case ActionTypes.Enqueue:
		return &Enqueue{ActionHeader: *h}
	case ActionTypes.Vendor:
		return &VendorActionHeader{ActionHeader: *h}
	default:
		return nil
	}
}

type ActionHeader struct {
	Type   ActionType
	Length uint16
}

func (header *ActionHeader) GetType() ActionType {
	return header.Type
}

type SendOutPort struct {
	ActionHeader
	Port      PortNumber
	MaxLength uint16
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

type StripVlan struct {
	ActionHeader
}

type SetEtherAddress struct {
	ActionHeader
	Address net.HardwareAddr
	pad     [6]uint8
}

type SetEtherSrc struct {
	ActionHeader
	Address [EthernetAddressLength]uint8
	pad     [6]uint8
}

type SetEtherDst struct {
	ActionHeader
	Address [EthernetAddressLength]uint8
	pad     [6]uint8
}

type SetIpAddress struct {
	ActionHeader
	Address net.IP
}

type SetIpSrc struct {
	ActionHeader
	Address [4]uint8
}

type SetIpDst struct {
	ActionHeader
	Address [4]uint8
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

type SetTransportSrc struct {
	ActionHeader
	Port TransportPort
	pad  [2]uint8
}

type SetTransportDst struct {
	ActionHeader
	Port TransportPort
	pad  [2]uint8
}

type VendorActionHeader struct {
	ActionHeader
	Vendor VendorId
}

type VendorId uint32
