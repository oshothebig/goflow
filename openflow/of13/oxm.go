package of13

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

type EtherType uint16
type VlanId uint16      // Lower 12 + 1 bits
type VlanPriority uint8 // Lower 3 bits
type Dscp uint8         // Lower 6 bits
type Ecn uint8          // Lower 2 bits
type ProtocolNumber uint8
type TcpPort uint16
type UdpPort uint16
type SctpPort uint16
type Icmpv4Type uint8
type Icmpv4Code uint8
type ArpOpcode uint16
type Ipv6FlowLabel uint32
type Ipv6FlowLabelMask uint32
type Icmpv6Type uint8
type Icmpv6Code uint8
type MplsLabel uint32   // Lower 20 bits
type MplsTc uint8       // Lower 3 bits
type MplsBos bool       // Lower 1 bit
type PbbIsid uint32     // Lower 24 bits
type PbbIsidMask uint32 // Lower 24 bits
type TunnelId uint64
type TunnelIdMask uint64
type Ipv6ExtHeader uint16 // Lower 9 bits

var oxmBitsMask = map[OxmField]struct {
	Bits        uint
	CanHaveMask bool
}{
	OxmFields.InPort:         {32, false},
	OxmFields.InPhysicalPort: {32, false},
	OxmFields.Metadata:       {64, true},
	OxmFields.EthDst:         {48, true},
	OxmFields.EthSrc:         {48, true},
	OxmFields.EthType:        {16, false},
	OxmFields.VlanId:         {13, true},
	OxmFields.VlanPriority:   {3, false},
	OxmFields.IpDscp:         {6, false},
	OxmFields.IpEcn:          {2, false},
	OxmFields.IpProtocol:     {8, false},
	OxmFields.Ipv4Src:        {32, true},
	OxmFields.Ipv4Dst:        {32, true},
	OxmFields.TcpSrc:         {16, false},
	OxmFields.TcpDst:         {16, false},
	OxmFields.UdpSrc:         {16, false},
	OxmFields.UdpDst:         {16, false},
	OxmFields.SctpSrc:        {16, false},
	OxmFields.SctpDst:        {16, false},
	OxmFields.Icmpv4Type:     {8, false},
	OxmFields.Icmpv4Code:     {8, false},
	OxmFields.ArpOpcode:      {16, false},
	OxmFields.ArpSpa:         {32, true},
	OxmFields.ArpTpa:         {32, true},
	OxmFields.ArpSha:         {48, true},
	OxmFields.ArpTha:         {48, true},
	OxmFields.Ipv6Src:        {128, true},
	OxmFields.Ipv6Dst:        {128, true},
	OxmFields.Ipv6FlowLabel:  {20, true},
	OxmFields.Icmpv6Type:     {8, false},
	OxmFields.Icmpv6Code:     {8, false},
	OxmFields.Ipv6NdTarget:   {128, false},
	OxmFields.Ipv6NdSll:      {48, false},
	OxmFields.Ipv6NdTll:      {48, false},
	OxmFields.MplsLabel:      {20, false},
	OxmFields.MplsTc:         {3, false},
	OxmFields.MplsBos:        {1, false},
	OxmFields.PbbIsid:        {24, true},
	OxmFields.TunnelId:       {64, true},
	OxmFields.Ipv6ExtHeader:  {9, true},
}

type OxmClass uint16

const OxmHeaderLength = 4

const (
	OFPXMC_NXM_0          OxmClass = 0x0000
	OFPXMC_NXM_1          OxmClass = 0x0001
	OFPXMC_OPENFLOW_BASIC OxmClass = 0x8000
	OFPXMC_EXPERIMENTER   OxmClass = 0xffff
)

var OxmClasses = struct {
	Nxm0          OxmClass
	Nxm1          OxmClass
	OpenFlowBasic OxmClass
	Experimenter  OxmClass
}{
	OFPXMC_NXM_0,
	OFPXMC_NXM_1,
	OFPXMC_OPENFLOW_BASIC,
	OFPXMC_EXPERIMENTER,
}

type OxmField uint8

// corresponds to oxm_ofb_match_fields
const (
	OFPXMT_OFB_IN_PORT OxmField = iota
	OFPXMT_OFB_IN_PHY_PORT
	OFPXMT_OFB_METADATA
	OFPXMT_OFB_ETH_DST
	OFPXMT_OFB_ETH_SRC
	OFPXMT_OFB_ETH_TYPE
	OFPXMT_OFB_VLAN_ID
	OFPXMT_OFB_VLAN_PCP
	OFPXMT_OFB_IP_DSCP
	OFPXMT_OFB_IP_ECN
	OFPXMT_OFB_IP_PROTO
	OFPXMT_OFB_IPV4_SRC
	OFPXMT_OFB_IPV4_DST
	OFPXMT_OFB_TCP_SRC
	OFPXMT_OFB_TCP_DST
	OFPXMT_OFB_UDP_SRC
	OFPXMT_OFB_UDP_DST
	OFPXMT_OFB_SCTP_SRC
	OFPXMT_OFB_SCTP_DST
	OFPXMT_OFB_ICMPV4_TYPE
	OFPXMT_OFB_ICMPV4_CODE
	OFPXMT_OFB_ARP_OP
	OFPXMT_OFB_ARP_SPA
	OFPXMT_OFB_ARP_TPA
	OFPXMT_OFB_ARP_SHA
	OFPXMT_OFB_ARP_THA
	OFPXMT_OFB_IPV6_SRC
	OFPXMT_OFB_IPV6_DST
	OFPXMT_OFB_IPV6_FLABEL
	OFPXMT_OFB_ICMPV6_TYPE
	OFPXMT_OFB_ICMPV6_CODE
	OFPXMT_OFB_IPV6_ND_TARGET
	OFPXMT_OFB_IPV6_ND_SLL
	OFPXMT_OFB_IPV6_ND_TLL
	OFPXMT_OFB_MPLS_LABEL
	OFPXMT_OFB_MPLS_TC
	OFPXMT_OFB_MPLS_BOS
	OFPXMT_OFB_PBB_ISID
	OFPXMT_OFB_TUNNEL_ID
	OFPXMT_OFB_IPV6_EXTHDR
)

var OxmFields = struct {
	InPort         OxmField
	InPhysicalPort OxmField
	Metadata       OxmField
	EthDst         OxmField
	EthSrc         OxmField
	EthType        OxmField
	VlanId         OxmField
	VlanPriority   OxmField
	IpDscp         OxmField
	IpEcn          OxmField
	IpProtocol     OxmField
	Ipv4Src        OxmField
	Ipv4Dst        OxmField
	TcpSrc         OxmField
	TcpDst         OxmField
	UdpSrc         OxmField
	UdpDst         OxmField
	SctpSrc        OxmField
	SctpDst        OxmField
	Icmpv4Type     OxmField
	Icmpv4Code     OxmField
	ArpOpcode      OxmField
	ArpSpa         OxmField
	ArpTpa         OxmField
	ArpSha         OxmField
	ArpTha         OxmField
	Ipv6Src        OxmField
	Ipv6Dst        OxmField
	Ipv6FlowLabel  OxmField
	Icmpv6Type     OxmField
	Icmpv6Code     OxmField
	Ipv6NdTarget   OxmField
	Ipv6NdSll      OxmField
	Ipv6NdTll      OxmField
	MplsLabel      OxmField
	MplsTc         OxmField
	MplsBos        OxmField
	PbbIsid        OxmField
	TunnelId       OxmField
	Ipv6ExtHeader  OxmField
}{
	OFPXMT_OFB_IN_PORT,
	OFPXMT_OFB_IN_PHY_PORT,
	OFPXMT_OFB_METADATA,
	OFPXMT_OFB_ETH_DST,
	OFPXMT_OFB_ETH_SRC,
	OFPXMT_OFB_ETH_TYPE,
	OFPXMT_OFB_VLAN_ID,
	OFPXMT_OFB_VLAN_PCP,
	OFPXMT_OFB_IP_DSCP,
	OFPXMT_OFB_IP_ECN,
	OFPXMT_OFB_IP_PROTO,
	OFPXMT_OFB_IPV4_SRC,
	OFPXMT_OFB_IPV4_DST,
	OFPXMT_OFB_TCP_SRC,
	OFPXMT_OFB_TCP_DST,
	OFPXMT_OFB_UDP_SRC,
	OFPXMT_OFB_UDP_DST,
	OFPXMT_OFB_SCTP_SRC,
	OFPXMT_OFB_SCTP_DST,
	OFPXMT_OFB_ICMPV4_TYPE,
	OFPXMT_OFB_ICMPV4_CODE,
	OFPXMT_OFB_ARP_OP,
	OFPXMT_OFB_ARP_SPA,
	OFPXMT_OFB_ARP_TPA,
	OFPXMT_OFB_ARP_SHA,
	OFPXMT_OFB_ARP_THA,
	OFPXMT_OFB_IPV6_SRC,
	OFPXMT_OFB_IPV6_DST,
	OFPXMT_OFB_IPV6_FLABEL,
	OFPXMT_OFB_ICMPV6_TYPE,
	OFPXMT_OFB_ICMPV6_CODE,
	OFPXMT_OFB_IPV6_ND_TARGET,
	OFPXMT_OFB_IPV6_ND_SLL,
	OFPXMT_OFB_IPV6_ND_TLL,
	OFPXMT_OFB_MPLS_LABEL,
	OFPXMT_OFB_MPLS_TC,
	OFPXMT_OFB_MPLS_BOS,
	OFPXMT_OFB_PBB_ISID,
	OFPXMT_OFB_TUNNEL_ID,
	OFPXMT_OFB_IPV6_EXTHDR,
}

type Oxm interface {
	Packetizable
	Header() *OxmHeader
}

type OxmHeader struct {
	Class   OxmClass
	Field   OxmField // only lower 7 bits are valid
	HasMask bool
	Length  uint8
}

func newOxmHeader(field OxmField, hasMask bool, bits uint) *OxmHeader {
	return &OxmHeader{
		OxmClasses.OpenFlowBasic, field, hasMask, uint8(calcOxmLength(bits, hasMask)),
	}
}

func (h *OxmHeader) Header() *OxmHeader {
	return h
}

func (h *OxmHeader) Len() uint {
	return OxmHeaderLength
}

func (h *OxmHeader) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, h.asUint32()); err != nil {
		return buf.Len(), err
	}

	if n, err = buf.Read(b); err != nil {
		return
	}
	return n, io.EOF
}

func (h *OxmHeader) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &h.Class); err != nil {
		return
	}
	var fieldAndMask uint8
	if err = binary.Read(buf, binary.BigEndian, &fieldAndMask); err != nil {
		return
	}
	h.Field = OxmField(fieldAndMask >> 1)
	if (fieldAndMask & 0x1) == 0 {
		h.HasMask = false
	} else {
		h.HasMask = true
	}
	if err = binary.Read(buf, binary.BigEndian, &h.Length); err != nil {
		return
	}
	n += int(h.Len())

	return
}

func (h *OxmHeader) asUint32() uint32 {
	class := uint32(h.Class) << 16
	field := uint32(h.Field&0x7f) << 9
	hasMask := uint32(0)
	if h.HasMask {
		hasMask = 1 << 8
	}
	length := uint32(h.Length)

	return class | field | hasMask | length
}

type OxmPort struct {
	OxmHeader
	Value PortNumber
}

func newOxmPort(field OxmField, port PortNumber) *OxmPort {
	return &OxmPort{
		*newOxmHeader(field, false, oxmBitsMask[OxmFields.InPort].Bits),
		port,
	}
}

func (o *OxmPort) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(&o.OxmHeader); err != nil {
		return
	}
	if err = binary.Write(buf, binary.BigEndian, o.Value); err != nil {
		return
	}

	if n, err = buf.Read(b); err != nil {
		return
	}
	return n, io.EOF
}

func (o *OxmPort) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if _, err = buf.WriteTo(&o.OxmHeader); err != nil {
		return
	}
	if err = binary.Read(buf, binary.BigEndian, &o.Value); err != nil {
		return
	}
	n += int(o.Len())

	return
}

func (o *OxmPort) Len() uint {
	return o.OxmHeader.Len() + calcOxmLength(oxmBitsMask[OxmFields.InPort].Bits, false)
}

type OxmInPort struct {
	OxmPort
}

func NewOxmInPort(port PortNumber) *OxmInPort {
	return &OxmInPort{*newOxmPort(OxmFields.InPort, port)}
}

func calcOxmLength(bits uint, hasMask bool) uint {
	unit := alignedSize(bits, 8) / 8
	if hasMask {
		return unit * 2
	} else {
		return unit
	}
}

type OxmInPhysicalPort struct {
	OxmPort
}

func NewOxmInPhysicalPort(port PortNumber) *OxmInPhysicalPort {
	return &OxmInPhysicalPort{*newOxmPort(OxmFields.InPhysicalPort, port)}
}

type OxmMetadata struct {
	OxmHeader
	Value Metadata
	Mask  MetadataMask
}

type OxmEthDst struct {
	OxmHeader
	Value net.HardwareAddr
	Mask  net.HardwareAddr
}

type OxmEthSrc struct {
	OxmHeader
	Value net.HardwareAddr
	Mask  net.HardwareAddr
}

type OxmEthType struct {
	OxmHeader
	Value EtherType
}

type OxmVlanId struct {
	OxmHeader
	Value VlanId
	Mask  VlanId
}

type OxmVlanPriority struct {
	OxmHeader
	Value VlanPriority
}

type OxmIpDscp struct {
	OxmHeader
	Value Dscp
}

type OxmIpEcn struct {
	OxmHeader
	Value Ecn
}

type OxmIpProtocol struct {
	OxmHeader
	Value ProtocolNumber
}

type OxmIpv4Src struct {
	OxmHeader
	Value net.IP
	Mask  net.IPMask
}

type OxmIpv4Dst struct {
	OxmHeader
	Value net.IP
	Mask  net.IPMask
}

type OxmTcpSrc struct {
	OxmHeader
	Value TcpPort
}

type OxmTcpDst struct {
	OxmHeader
	Value TcpPort
}

type OxmUdpSrc struct {
	OxmHeader
	Value UdpPort
}

type OxmUdpDst struct {
	OxmHeader
	Value UdpPort
}

type OxmSctpSrc struct {
	OxmHeader
	Value SctpPort
}

type OxmSctpDst struct {
	OxmHeader
	Value SctpPort
}

type OxmIcmpv4Type struct {
	OxmHeader
	Value Icmpv4Type
}

type OxmIcmpv4Code struct {
	OxmHeader
	Value Icmpv4Code
}

type OxmArpOpcode struct {
	OxmHeader
	Value ArpOpcode
}

type OxmArpSpa struct {
	OxmHeader
	Value net.IP
	Mask  net.IPMask
}

type OxmArpTpa struct {
	OxmHeader
	Value net.IP
	Mask  net.IPMask
}

type OxmArpSha struct {
	OxmHeader
	Value net.HardwareAddr
	Mask  net.HardwareAddr
}

type OxmArpTha struct {
	OxmHeader
	Value net.HardwareAddr
	Mask  net.HardwareAddr
}

type OxmIpv6Src struct {
	OxmHeader
	Value net.IP
	Mask  net.IPMask
}

type OxmIpv6Dst struct {
	OxmHeader
	Value net.IP
	Mask  net.IPMask
}

type OxmIpv6FlowLabel struct {
	OxmHeader
	Value Ipv6FlowLabel
	Mask  Ipv6FlowLabelMask
}

type OxmIcmpv6Type struct {
	OxmHeader
	Value Icmpv6Type
}

type OxmIcmpv6Code struct {
	OxmHeader
	Value Icmpv6Code
}

type OxmIpv6NdTarget struct {
	OxmHeader
	Value net.IP
}

type OxmIpv6NdSll struct {
	OxmHeader
	Value net.HardwareAddr
}

type OxmIpv6NdTll struct {
	OxmHeader
	Value net.HardwareAddr
}

type OxmMplsLabel struct {
	OxmHeader
	Value MplsLabel
}

type OxmMplsTc struct {
	OxmHeader
	Value MplsTc
}

type OxmMplsBos struct {
	OxmHeader
	Value MplsBos
}

type OxmPbbIsid struct {
	OxmHeader
	Value PbbIsid
	Mask  PbbIsidMask
}

type OxmTunnelId struct {
	OxmHeader
	Value TunnelId
	Mask  TunnelIdMask
}

type OxmIpv6ExtHeader struct {
	OxmHeader
	Value Ipv6ExtHeader
}
