package of13

import (
	"bytes"
	"encoding/binary"
	"io"

	. "github.com/oshothebig/goflow/openflow"
)

const (
	ActionHeaderLength  = 4
	MinimumActionLength = 8
	SendOutPortLength   = 16
)

type ActionType uint16

// corresponds to ofp_action_type
const (
	OFPAT_OUTPUT       ActionType = 0
	OFPAT_COPY_TTL_OUT ActionType = 11
	OFPAT_COPY_TTL_IN  ActionType = 12
	OFPAT_SET_MPLS_TTL ActionType = 15
	OFPAT_DEC_MPLS_TTL ActionType = 16
	OFPAT_PUSH_VLAN    ActionType = 17
	OFPAT_POP_VLAN     ActionType = 18
	OFPAT_PUSH_MPLS    ActionType = 19
	OFPAT_POP_MPLS     ActionType = 20
	OFPAT_SET_QUEUE    ActionType = 21
	OFPAT_GROUP        ActionType = 22
	OFPAT_SET_NW_TTL   ActionType = 23
	OFPAT_DEC_NW_TTL   ActionType = 24
	OFPAT_SET_FIELD    ActionType = 25
	OFPAT_PUSH_PBB     ActionType = 26
	OFPAT_POP_PBB      ActionType = 27
	OFPAT_EXPERIMENTER ActionType = 0xffff
)

// corresponds to ofp_controller_max_len
const (
	OFPCML_MAX       = 0xffe5
	OFPCML_NO_BUFFER = 0xffff
)

var ActionTypes = struct {
	SendOutPort  ActionType
	CopyTtlOut   ActionType
	CopyTtlIn    ActionType
	SetMplsTtl   ActionType
	DecrMplsTtl  ActionType
	PushVlan     ActionType
	PopVlan      ActionType
	PushMpls     ActionType
	PopMpls      ActionType
	SetQueue     ActionType
	GroupAction  ActionType
	SetIpTtl     ActionType
	DecrIpTtl    ActionType
	SetField     ActionType
	PushPbb      ActionType
	PopPbb       ActionType
	Experimenter ActionType
}{
	OFPAT_OUTPUT,
	OFPAT_COPY_TTL_OUT,
	OFPAT_COPY_TTL_IN,
	OFPAT_SET_MPLS_TTL,
	OFPAT_DEC_MPLS_TTL,
	OFPAT_PUSH_VLAN,
	OFPAT_POP_VLAN,
	OFPAT_PUSH_MPLS,
	OFPAT_POP_MPLS,
	OFPAT_SET_QUEUE,
	OFPAT_GROUP,
	OFPAT_SET_NW_TTL,
	OFPAT_DEC_NW_TTL,
	OFPAT_SET_FIELD,
	OFPAT_PUSH_PBB,
	OFPAT_POP_PBB,
	OFPAT_EXPERIMENTER,
}

type Action interface {
	Packetizable
	GetType() ActionType
}

type SendOutPort struct {
	Type      ActionType
	Length    uint16
	Port      uint32
	MaxLength uint16
	_         [6]uint8
}

func NewSendOutPort(port uint32, length uint16) *SendOutPort {
	m := &SendOutPort{
		Type:      ActionTypes.SendOutPort,
		Length:    SendOutPortLength,
		Port:      port,
		MaxLength: length,
	}

	return m
}

func (a *SendOutPort) Len() uint {
	return SendOutPortLength
}

func (a *SendOutPort) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *SendOutPort) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *SendOutPort) GetType() ActionType {
	return a.Type
}

type MinimumActionStruct struct {
	Type   ActionType
	Length uint16
	_      [4]uint8
}

func newMinimumActionStruct(typ ActionType) *MinimumActionStruct {
	return &MinimumActionStruct{Type: typ, Length: MinimumActionLength}
}

func (a *MinimumActionStruct) Len() uint {
	return MinimumActionLength
}

func (a *MinimumActionStruct) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *MinimumActionStruct) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *MinimumActionStruct) GetType() ActionType {
	return a.Type
}

type CopyTtlOut struct {
	MinimumActionStruct
}

func NewCopyTtlOut() *CopyTtlOut {
	return &CopyTtlOut{*newMinimumActionStruct(ActionTypes.CopyTtlOut)}
}

type CopyTtlIn struct {
	MinimumActionStruct
}

func NewCopyTtlIn() *CopyTtlIn {
	return &CopyTtlIn{*newMinimumActionStruct(ActionTypes.CopyTtlIn)}
}

type SetMplsTtl struct {
	Type   ActionType
	Length uint16
	TTL    uint8
	_      [3]uint8
}

func (a *SetMplsTtl) Len() uint {
	return MinimumActionLength
}

func (a *SetMplsTtl) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *SetMplsTtl) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *SetMplsTtl) GetType() ActionType {
	return a.Type
}

func NewSetMplsTtl(ttl uint8) *SetMplsTtl {
	return &SetMplsTtl{
		Type:   ActionTypes.SetMplsTtl,
		Length: MinimumActionLength,
		TTL:    ttl,
	}
}

type DecrMplsTtl struct {
	MinimumActionStruct
}

func NewDecrMplsTtl() *DecrMplsTtl {
	return &DecrMplsTtl{*newMinimumActionStruct(ActionTypes.DecrMplsTtl)}
}

type PushAction struct {
	Type      ActionType
	Length    uint16
	EtherType uint16
	_         [2]uint8
}

func newPushAction(at ActionType, et uint16) *PushAction {
	return &PushAction{
		Type:      at,
		Length:    MinimumActionLength,
		EtherType: et,
	}
}

func (a *PushAction) Len() uint {
	return MinimumActionLength
}

func (a *PushAction) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *PushAction) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *PushAction) GetType() ActionType {
	return a.Type
}

type PushVlan struct {
	PushAction
}

func NewPushVlan(t uint16) *PushVlan {
	return &PushVlan{*newPushAction(ActionTypes.PushVlan, t)}
}

type PopVlan struct {
	MinimumActionStruct
}

func NewPopVlan() *PopVlan {
	return &PopVlan{*newMinimumActionStruct(ActionTypes.PopVlan)}
}

type PushMpls struct {
	PushAction
}

func NewPushMpls(t uint16) *PushMpls {
	return &PushMpls{*newPushAction(ActionTypes.PushMpls, t)}
}

type PopMpls struct {
	Type      ActionType
	Length    uint16
	EtherType uint16
	_         [2]uint8
}

func NewPopMpls(t uint16) *PopMpls {
	return &PopMpls{
		Type:      ActionTypes.PopMpls,
		Length:    MinimumActionLength,
		EtherType: t,
	}
}

func (a *PopMpls) Len() uint {
	return MinimumActionLength
}

func (a *PopMpls) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *PopMpls) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *PopMpls) GetType() ActionType {
	return a.Type
}

type SetQueue struct {
	Type    ActionType
	Length  uint16
	QueueId uint32
}

func NewSetQueue(id uint32) *SetQueue {
	return &SetQueue{
		Type:    ActionTypes.SetQueue,
		Length:  MinimumActionLength,
		QueueId: id,
	}
}

func (a *SetQueue) Len() uint {
	return MinimumActionLength
}

func (a *SetQueue) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *SetQueue) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *SetQueue) GetType() ActionType {
	return a.Type
}

type GroupAction struct {
	Type    ActionType
	Length  uint16
	GroupId uint32
}

func NewGroupAction(id uint32) *GroupAction {
	return &GroupAction{
		Type:    ActionTypes.GroupAction,
		Length:  8,
		GroupId: id,
	}
}

func (a *GroupAction) Len() uint {
	return MinimumActionLength
}

func (a *GroupAction) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *GroupAction) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *GroupAction) GetType() ActionType {
	return a.Type
}

type SetIpTtl struct {
	Type   ActionType
	Length uint16
	TTL    uint8
	_      [3]uint8
}

func NewSetIpTtl(ttl uint8) *SetIpTtl {
	return &SetIpTtl{
		Type:   ActionTypes.SetIpTtl,
		Length: MinimumActionLength,
		TTL:    ttl,
	}
}

func (a *SetIpTtl) Len() uint {
	return MinimumActionLength
}

func (a *SetIpTtl) Read(b []byte) (n int, err error) {
	return marshalFixedSizeData(a, b)
}

func (a *SetIpTtl) Write(b []byte) (n int, err error) {
	return unmarshalFixedSizeData(a, b)
}

func (a *SetIpTtl) GetType() ActionType {
	return a.Type
}

type DecrIpTtl struct {
	MinimumActionStruct
}

func NewDecrIpTtl() *DecrIpTtl {
	return &DecrIpTtl{*newMinimumActionStruct(ActionTypes.DecrIpTtl)}
}

type SetField struct {
	Type   ActionType
	Length uint16

	Field Oxm
	pad   []uint8
}

func NewSetField(o Oxm) *SetField {
	s := &SetField{
		Type:  ActionTypes.SetField,
		Field: o,
	}
	s.Length = uint16(s.Len())
	s.pad = make([]uint8, s.Len()-ActionHeaderLength-s.Field.Len())

	return s
}

func (a *SetField) Len() uint {
	size := ActionHeaderLength + a.Field.Len()
	return alignedSize(size, defaultAlign)
}

func (a *SetField) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, a.Type); err != nil {
		return
	}
	if err = binary.Write(buf, binary.BigEndian, a.Length); err != nil {
		return
	}
	if _, err = buf.ReadFrom(a.Field); err != nil {
		return
	}
	if err = binary.Write(buf, binary.BigEndian, a.pad); err != nil {
		return
	}

	if n, err = buf.Read(b); err != nil {
		return
	}
	return n, io.EOF
}

func (a *SetField) GetType() ActionType {
	return a.Type
}

type PushPbb struct {
	PushAction
}

func NewPushPbb(t uint16) *PushPbb {
	return &PushPbb{*newPushAction(ActionTypes.PushPbb, t)}
}

type PopPbb struct {
	MinimumActionStruct
}

func NewPopPbb() *PopPbb {
	return &PopPbb{*newMinimumActionStruct(ActionTypes.PopPbb)}
}

type ExperimenterActionHeader struct {
	Type         ActionType
	Length       uint16
	Experimenter uint32
}
