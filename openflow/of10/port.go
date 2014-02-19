package of10

import (
	"net"

	. "github.com/oshothebig/goflow/openflow"
)

const MaxPortNameLength = 16

type PortNumber uint16
type PortConfig uint16
type PortState uint16
type PortFeature uint16

type PhysicalPort struct {
	PortNumber         PortNumber
	HardwareAddress    net.HardwareAddr
	Name               [MaxPortNameLength]uint8
	Config             PortConfig
	State              PortState
	CurrentFeatures    PortFeature
	AdvertisedFeatures PortFeature
	SupportedFeatures  PortFeature
	PeerFeatures       PortFeature
}
