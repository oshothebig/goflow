package of10

import (
	"testing"
)

func TestNewFeaturesRequest(t *testing.T) {
	instance := NewFeaturesRequest()
	actual := instance.GetHeader().Type
	if actual != OFPT_FEATURES_REQUEST {
		t.Errorf("Actual: %#v, Expected: %#v", actual, OFPT_FEATURES_REQUEST)
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
