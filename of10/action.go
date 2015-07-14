package of10

import (
	"bytes"
	"encoding"
	"encoding/binary"
	"errors"
	"io"
)

var (
	errUnsupportedAction = errors.New("Unsupported action type")
)

type ActionType uint16

type Action interface {
	GetType() ActionType
	encoding.BinaryUnmarshaler
}

type actionDecoder struct {
	rd  *bytes.Reader
	end int
}

func readActions(rd *bytes.Reader, length int) []Action {
	actions := make([]Action, 0, 8)
	decoder := &actionDecoder{rd, rd.Len() - length}
	for decoder.canDecode() {
		action, err := decoder.decode()
		if err != nil {
			break
		}
		actions = append(actions, action)
	}
	return actions
}

func (d *actionDecoder) canDecode() bool {
	return d.rd.Len() > d.end
}

func (d *actionDecoder) header() (*ActionHeader, error) {
	start := d.rd.Len()
	var header ActionHeader
	if err := binary.Read(d.rd, binary.BigEndian, &header); err != nil {
		return nil, err
	}

	// unread length of header
	offset := d.rd.Len() - start
	d.rd.Seek(int64(offset), 1)

	return &header, nil
}

func (d *actionDecoder) decode() (Action, error) {
	header, err := d.header()
	if err != nil {
		return nil, err
	}

	action := newAction(header.Type)
	if action == nil {
		return nil, errUnsupportedAction
	}

	data := make([]byte, header.Length)
	if _, err := io.ReadFull(d.rd, data); err != nil {
		return nil, err
	}

	if err := action.UnmarshalBinary(data); err != nil {
		return nil, err
	}

	return action, nil
}

func newAction(t ActionType) Action {
	switch t {
	case ActionTypes.Output:
		return new(SendOutPort)
	case ActionTypes.SetVlanId:
		return new(SetVlanVid)
	case ActionTypes.SetVlanPcp:
		return new(SetVlanPcp)
	case ActionTypes.StripVlan:
		return new(StripVlan)
	case ActionTypes.SetEtherSrc:
		return new(SetEtherSrc)
	case ActionTypes.SetEtherDst:
		return new(SetEtherDst)
	case ActionTypes.SetIpSrc:
		return new(SetIpSrc)
	case ActionTypes.SetIpDst:
		return new(SetIpDst)
	case ActionTypes.SetIpTos:
		return new(SetIpTos)
	case ActionTypes.SetNetworkSrc:
		return new(SetTransportSrc)
	case ActionTypes.SetNetworkDst:
		return new(SetTransportDst)
	case ActionTypes.Enqueue:
		return new(Enqueue)
	case ActionTypes.Vendor:
		return new(VendorActionHeader)
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

func (m *SendOutPort) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type Enqueue struct {
	ActionHeader
	Port    PortNumber
	_       [6]uint8
	QueueId uint32
}

func (m *Enqueue) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type SetVlanVid struct {
	ActionHeader
	Id VlanId
	_  [2]uint32
}

func (m *SetVlanVid) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type SetVlanPcp struct {
	ActionHeader
	Priority VlanPriority
	_        [3]uint8
}

func (m *SetVlanPcp) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type StripVlan struct {
	ActionHeader
}

func (m *StripVlan) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), m)
}

type SetEtherSrc struct {
	ActionHeader
	Address [EthernetAddressLength]uint8
	pad     [6]uint8
}

func (a *SetEtherSrc) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type SetEtherDst struct {
	ActionHeader
	Address [EthernetAddressLength]uint8
	pad     [6]uint8
}

func (a *SetEtherDst) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type SetIpSrc struct {
	ActionHeader
	Address [4]uint8
}

func (a *SetIpSrc) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type SetIpDst struct {
	ActionHeader
	Address [4]uint8
}

func (a *SetIpDst) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type SetIpTos struct {
	ActionHeader
	Tos Dscp
	pad [3]uint8
}

func (a *SetIpTos) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type SetTransportSrc struct {
	ActionHeader
	Port TransportPort
	pad  [2]uint8
}

func (a *SetTransportSrc) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type SetTransportDst struct {
	ActionHeader
	Port TransportPort
	pad  [2]uint8
}

func (a *SetTransportDst) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type VendorActionHeader struct {
	ActionHeader
	Vendor VendorId
}

func (a *VendorActionHeader) UnmarshalBinary(data []byte) error {
	return unmarshalFields(bytes.NewReader(data), a)
}

type VendorId uint32
