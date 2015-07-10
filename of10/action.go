package of10

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var (
	errUnsupportedAction error = errors.New("Unsupported action type")
)

type ActionType uint16

const actionHeaderLength = 4

type Action interface {
	GetType() ActionType
	FillBody(buf *bytes.Buffer) error
}

func readAction(buf *bytes.Buffer) (Action, error) {
	// read action header
	var header ActionHeader
	if err := binary.Read(buf, binary.BigEndian, &header); err != nil {
		return nil, err
	}

	// make an empty action
	action := newAction(&header)
	if action == nil {
		return nil, errUnsupportedAction
	}

	// read remaining body
	body := make([]byte, header.Length-actionHeaderLength)
	if _, err := io.ReadFull(buf, body); err != nil {
		return nil, err
	}

	// parse action body
	bodyBuf := bytes.NewBuffer(body)
	if err := action.FillBody(bodyBuf); err != nil {
		return nil, err
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

func (a *SendOutPort) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Port); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.MaxLength); err != nil {
		return err
	}
	return nil
}

type Enqueue struct {
	ActionHeader
	Port    PortNumber
	pad     [6]uint8
	QueueId uint32
}

func (a *Enqueue) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Port); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.QueueId); err != nil {
		return err
	}
	return nil
}

type SetVlanVid struct {
	ActionHeader
	Id  VlanId
	pad [2]uint32
}

func (a *SetVlanVid) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Id); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	return nil
}

type SetVlanPcp struct {
	ActionHeader
	Priority VlanPriority
	pad      [3]uint8
}

func (a *SetVlanPcp) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Priority); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	return nil
}

type StripVlan struct {
	ActionHeader
}

func (a *StripVlan) FillBody(buf *bytes.Buffer) error {
	return nil
}

type SetEtherSrc struct {
	ActionHeader
	Address [EthernetAddressLength]uint8
	pad     [6]uint8
}

func (a *SetEtherSrc) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Address); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	return nil
}

type SetEtherDst struct {
	ActionHeader
	Address [EthernetAddressLength]uint8
	pad     [6]uint8
}

func (a *SetEtherDst) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Address); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	return nil
}

type SetIpSrc struct {
	ActionHeader
	Address [4]uint8
}

func (a *SetIpSrc) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Address); err != nil {
		return err
	}
	return nil
}

type SetIpDst struct {
	ActionHeader
	Address [4]uint8
}

func (a *SetIpDst) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Address); err != nil {
		return err
	}
	return nil
}

type SetIpTos struct {
	ActionHeader
	Tos Dscp
	pad [3]uint8
}

func (a *SetIpTos) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Tos); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	return nil
}

type SetTransportSrc struct {
	ActionHeader
	Port TransportPort
	pad  [2]uint8
}

func (a *SetTransportSrc) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Port); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	return nil
}

type SetTransportDst struct {
	ActionHeader
	Port TransportPort
	pad  [2]uint8
}

func (a *SetTransportDst) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Port); err != nil {
		return err
	}
	if err := binary.Read(buf, binary.BigEndian, &a.pad); err != nil {
		return err
	}
	return nil
}

type VendorActionHeader struct {
	ActionHeader
	Vendor VendorId
}

func (a *VendorActionHeader) FillBody(buf *bytes.Buffer) error {
	if err := binary.Read(buf, binary.BigEndian, &a.Vendor); err != nil {
		return err
	}
	return nil
}

type VendorId uint32
