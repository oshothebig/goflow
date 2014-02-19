package of13

type MatchType uint16
type Ipv6ExtHeaderFlag uint16

type Match struct {
	Type      MatchType
	Length    uint16
	OxmFields []Oxm
}

// corresponds to ofp_match_type
const (
	OFPMT_STANDARD MatchType = iota
	OFPMT_OXM
)

var MatchTypes = struct {
	Standard MatchType
	Oxm      MatchType
}{
	OFPMT_STANDARD,
	OFPMT_OXM,
}

// corresponds to ofp_vlan_id
const (
	OFPVID_PRESENT VlanId = 0x1000
	OFPVID_NONE    VlanId = 0x0000
)

// corresponds to ofp_ipv6exthdr_flags
const (
	OFPIEH_NONEXT Ipv6ExtHeaderFlag = 1 << iota
	OFPIEH_ESP
	OFPIEH_AUTH
	OFPIEH_DEST
	OFPIEH_FRAG
	OFPIEH_ROUTER
	OFPIEH_HOP
	OFPIEH_UNREP
	OFPIEH_UNSEQ
)

// TODO: implement ofp_oxm_experimenter_header
