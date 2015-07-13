package of10

import (
	"bytes"
	"encoding/binary"
)

const MaxPortNameLength = 16
const EthernetAddressLength = 6

type PortNumber uint16
type PortConfig uint16
type PortState uint16
type PortFeature uint16

type PhysicalPort struct {
	PortNumber         PortNumber
	HardwareAddress    [EthernetAddressLength]uint8
	Name               [MaxPortNameLength]uint8
	Config             PortConfig
	State              PortState
	CurrentFeatures    PortFeature
	AdvertisedFeatures PortFeature
	SupportedFeatures  PortFeature
	PeerFeatures       PortFeature
}

func readPhysicalPort(b []byte) ([]PhysicalPort, error) {
	var port PhysicalPort
	count := len(b) / binary.Size(port)
	ports := make([]PhysicalPort, 0, count)

	buf := bytes.NewBuffer(b)
	if err := binary.Read(buf, binary.BigEndian, port); err != nil {
		return nil, err
	}

	return ports, nil
}
