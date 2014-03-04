package goflow

import (
	"bytes"
	"testing"
)

var uint64Data = []struct {
	data     []byte
	expected int64
}{
	{[]byte{0x01, 0x02, 0x03}, 0x010203},
	{[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}, 0x0102030405060708},
	{[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}, 0x0203040506070809},
}

func TestBigEndianValue(t *testing.T) {
	for _, v := range uint64Data {
		actual := bigEndianValue(v.data)
		if actual != v.expected {
			t.Errorf("Got %d, expected %d", actual, v.expected)
		}
	}
}

var divideData = []struct {
	divider  Divider
	data     []byte
	expected []byte
}{
	{
		&lengthBasedDivider{32, 0, 2, 0},
		[]byte{0x00, 0x02, 0x01, 0x02, 0x03},
		[]byte{0x00, 0x02, 0x01, 0x02},
	},
	{
		&lengthBasedDivider{32, 0, 2, 2},
		[]byte{
			0x00, 0x02, // length: 2
			0x01, 0x02, 0x03, 0x04, 0x05,
		},
		[]byte{
			0x00, 0x02,
			0x01, 0x02, 0x03, 0x04,
		},
	},
	{
		&lengthBasedDivider{65535, 2, 2, -4},
		[]byte{
			0x01,       // version
			0x00,       // type: Hello
			0x00, 0x08, // length: 8
			0x00, 0x00, 0x00, 0x01, // xid: 1
			0x00, // byte not needed
		},
		[]byte{
			0x01,
			0x00,
			0x00, 0x08,
			0x00, 0x00, 0x00, 0x01,
		},
	},
}

func TestDivide(t *testing.T) {
	for _, v := range divideData {
		divider := v.divider
		actual, err := divider.Divide(bytes.NewReader(v.data))
		if err != nil {
			t.Errorf("Error: %v in %#v", err, v)
		}
		if bytes.Compare(actual, v.expected) != 0 {
			t.Errorf("Got %v, expected %v", actual, v.expected)
		}
	}
}
