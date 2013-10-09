package goflow

import (
	"reflect"
	"testing"
)

var samples = map[ActionType]struct {
	Bytes  []byte
	Action Action
}{
	ActionTypes.SendOutPort: {
		[]byte{
			0x00, 0x00, // Type: 0
			0x00, 0x10, // Length: 16
			0x00, 0x00, 0x00, 0x01, // Port: 1
			0x12, 0x34, // MaxLength: 0x1234
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // pad
		}, NewSendOutPort(1, 0x1234),
	},
	ActionTypes.CopyTtlOut: {
		[]byte{
			0x00, 0x0b, // Type: 11
			0x00, 0x08, // Length: 8
			0x00, 0x00, 0x00, 0x00,
		}, NewCopyTtlOut(),
	},
	ActionTypes.CopyTtlIn: {
		[]byte{
			0x00, 0x0c, // Type: 12
			0x00, 0x08, // Type: 16
			0x00, 0x00, 0x00, 0x00,
		}, NewCopyTtlIn(),
	},
	ActionTypes.SetMplsTtl: {
		[]byte{
			0x00, 0x0f, // Type: 15
			0x00, 0x08, // Length: 8
			0x12,             // TTL: 18
			0x00, 0x00, 0x00, // pad
		}, NewSetMplsTtl(18),
	},
	ActionTypes.DecrMplsTtl: {
		[]byte{
			0x00, 0x10, // Type: 16
			0x00, 0x08, // Length: 8
			0x00, 0x00, 0x00, 0x00,
		}, NewDecrMplsTtl(),
	},
	ActionTypes.PushVlan: {
		[]byte{
			0x00, 0x11, // Type: 17
			0x00, 0x08, // Length: 8
			0x08, 0x00, // EtherType: 0x0800 (IPv4)
			0x00, 0x00,
		}, NewPushVlan(0x0800),
	},
	ActionTypes.PopVlan: {
		[]byte{
			0x00, 0x12, // Type: 18
			0x00, 0x08, // Length: 8
			0x00, 0x00, 0x00, 0x00,
		}, NewPopVlan(),
	},
	ActionTypes.PushMpls: {
		[]byte{
			0x00, 0x13, // Type: 19
			0x00, 0x08, // Length: 8
			0x08, 0x00, // EtherType: 0x0800 (IPv4)
			0x00, 0x00,
		}, NewPushMpls(0x0800),
	},
	ActionTypes.PopMpls: {
		[]byte{
			0x00, 0x14, // Type: 20
			0x00, 0x08, // Length: 8
			0x08, 0x00, // EtherType: 0x0800 (IPv4)
			0x00, 0x00,
		}, NewPopMpls(0x0800),
	},
	ActionTypes.SetQueue: {
		[]byte{
			0x00, 0x15, // Type: 21
			0x00, 0x08, // Length: 8
			0x12, 0x34, 0x56, 0x78, // QueueId: 0x12345678
		}, NewSetQueue(0x12345678),
	},
	ActionTypes.GroupAction: {
		[]byte{
			0x00, 0x16, // Type: 22
			0x00, 0x08, // Length: 8
			0x12, 0x34, 0x56, 0x78,
		}, NewGroupAction(0x12345678),
	},
	ActionTypes.SetIpTtl: {
		[]byte{
			0x00, 0x17, // Type: 23
			0x00, 0x08, // Length: 8
			0x12, // TTL: 0x12
			0x00, 0x00, 0x00,
		}, NewSetIpTtl(0x12),
	},
	ActionTypes.DecrIpTtl: {
		[]byte{
			0x00, 0x18, // Type: 24
			0x00, 0x08, // Length: 8
			0x00, 0x00, 0x00, 0x00,
		}, NewDecrIpTtl(),
	},
	ActionTypes.PushPbb: {
		[]byte{
			0x00, 0x1a, // Type: 26
			0x00, 0x08, // Length: 8
			0x08, 0x00, // EtherType: 0x0800 (IPv4)
			0x00, 0x00,
		}, NewPushPbb(0x0800),
	},
	ActionTypes.PopPbb: {
		[]byte{
			0x00, 0x1b, // Type: 27
			0x00, 0x08, // Length: 8
			0x00, 0x00, 0x00, 0x00,
		}, NewPopPbb(),
	},
}

func checkActionType(t *testing.T, actual, expected ActionType) {
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}

func TestActionType(t *testing.T) {
	checkActionType(t, OFPAT_SET_FIELD, 25)
	checkActionType(t, OFPAT_EXPERIMENTER, 0xffff)

	checkActionType(t, ActionTypes.SetField, OFPAT_SET_FIELD)
	checkActionType(t, ActionTypes.Experimenter, OFPAT_EXPERIMENTER)
}

func TestActionRead(t *testing.T) {
	for _, v := range samples {
		checkMarshall(t, v.Action, v.Bytes)
	}
}

func TestActionWrite(t *testing.T) {
	for _, v := range samples {
		action := reflect.New(reflect.TypeOf(v.Action).Elem()).Interface()
		empty := action.(Action)
		checkUnmarshal(t, v.Bytes, empty, v.Action)
	}
}

func TestActionSize(t *testing.T) {
	for _, v := range samples {
		checkPacketizableSize(t, v.Action, len(v.Bytes))
	}
}

// TODO: add test for SetField, which is variable size struct
