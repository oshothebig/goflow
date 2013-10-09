package goflow

import "testing"

func TestPortConfigs(t *testing.T) {
	if PortConfigs.PortDown != 1 {
		t.Errorf("PortConfigs.PortDown=%d, want %d", PortConfigs.PortDown, 1)
	}

	if PortConfigs.NoReceive != 4 {
		t.Errorf("PortConfigs.NoReceive=%d, want %d", PortConfigs.NoReceive, 4)
	}

	if PortConfigs.NoForward != 32 {
		t.Errorf("PortConfigs.NoForward=%d, want %d", PortConfigs.NoForward, 32)
	}

	if PortConfigs.NoPacketIn != 64 {
		t.Errorf("PortConfigs.NoPacketIn=%d, want %d", PortConfigs.NoPacketIn, 64)
	}
}

func TestPortStateValue(t *testing.T) {
	if OFPPS_LINK_DOWN != 1 {
		t.Errorf("PortStateLinkDown=%d, want %d", OFPPS_LINK_DOWN, 1)
	}

	if OFPPS_BLOCKED != 2 {
		t.Errorf("PortStateBlocked=%d, want %d", OFPPS_BLOCKED, 2)
	}

	if OFPPS_LIVE != 4 {
		t.Error("PortStateLive=%d, want %d", OFPPS_LIVE, 4)
	}
}
