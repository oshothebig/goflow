package of10

import (
	"testing"
)

func TestNewHello(t *testing.T) {
	instance := NewHello(nil)
	actual := instance.GetHeader().Type
	if actual != OFPT_HELLO {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_HELLO)
	}
}

func TestNewEchoRequest(t *testing.T) {
	instance := NewEchoRequest(nil)
	actual := instance.GetHeader().Type
	if actual != OFPT_ECHO_REQUEST {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_ECHO_REQUEST)
	}
}

func TestNewEchoReply(t *testing.T) {
	instance := NewEchoReply(nil)
	actual := instance.GetHeader().Type
	if actual != OFPT_ECHO_REPLY {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_ECHO_REPLY)
	}
}

func TestNewFeaturesRequest(t *testing.T) {
	instance := NewFeaturesRequest()
	actual := instance.GetHeader().Type
	if actual != OFPT_FEATURES_REQUEST {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_FEATURES_REQUEST)
	}
}

func TestNewGetConfigRequest(t *testing.T) {
	instance := NewGetConfigRequest()
	actual := instance.GetHeader().Type
	if actual != OFPT_GET_CONFIG_REQUEST {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_GET_CONFIG_REQUEST)
	}
}

func TestNewGetConfigReply(t *testing.T) {
	instance := NewGetConfigReply(ConfigFlags.FragmentNormal, 0)
	actual := instance.GetHeader().Type
	if actual != OFPT_GET_CONFIG_REPLY {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_GET_CONFIG_REPLY)
	}
}

func TestNewSetConfig(t *testing.T) {
	instance := NewSetConfig(ConfigFlags.FragmentNormal, 0)
	actual := instance.GetHeader().Type
	if actual != OFPT_SET_CONFIG {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_SET_CONFIG)
	}
}

func TestNewBarrierRequest(t *testing.T) {
	instance := NewBarrierRequest()
	actual := instance.GetHeader().Type
	if actual != OFPT_BARRIER_REQUEST {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_BARRIER_REQUEST)
	}
}

func TestNewBarrierReply(t *testing.T) {
	instance := NewBarrierReply()
	actual := instance.GetHeader().Type
	if actual != OFPT_BARRIER_REPLY {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_BARRIER_REPLY)
	}
}
