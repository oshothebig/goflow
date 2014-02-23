package of10

import . "github.com/oshothebig/goflow/openflow"

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

func (header *ActionHeader) GetType() ActionType {
	return header.Type
}

type SendOutPort struct {
	ActionHeader
	Port PortNumber
	MaxLength uint16
}