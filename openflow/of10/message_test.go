package of10

import (
	"testing"
)

func TestNewFeaturesRequest(t *testing.T) {
	instance := NewFeaturesRequest()
	actual := instance.GetHeader().Type
	if actual != OFPT_FEATURES_REQUEST {
		t.Errorf("Actual: %#v, Expected: %#v", actual, MessageTypes.FeaturesRequest)
	}
}
