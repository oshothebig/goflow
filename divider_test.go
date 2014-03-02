package goflow

import (
	"bytes"
	"testing"
)

var uint64Data = []struct {
	data     []byte
	expected uint64
}{
	{[]byte{0x01, 0x02, 0x03}, 0x010203},
	{[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}, 0x0102030405060708},
	{[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}, 0x0203040506070809},
}

func TestUint64(t *testing.T) {
	for _, v := range uint64Data {
		actual := Uint64(v.data)
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
		&lengthBasedDivider{32, 0, 2},
		[]byte{0x00, 0x02, 0x01, 0x02, 0x03},
		[]byte{0x00, 0x02, 0x01, 0x02},
	},
}

func TestDivide(t *testing.T) {
	for _, v := range divideData {
		divider := v.divider
		actual, _ := divider.Divide(bytes.NewReader(v.data))
		if bytes.Compare(actual, v.expected) != 0 {
			t.Errorf("Got %v, expected %v", actual, v.expected)
		}
	}
}
