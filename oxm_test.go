package goflow

import (
	"reflect"
	"testing"
)

var oxms = map[OxmField]struct {
	Bytes []byte
	Oxm   Oxm
}{
	OxmFields.InPort: {[]byte{
		0x80, 0x00,
		0x00,
		0x04,
		0x12, 0x34, 0x56, 0x78,
	}, NewOxmInPort(0x12345678)},
	OxmFields.InPhysicalPort: {[]byte{
		0x80, 0x00,
		0x02,
		0x04,
		0x12, 0x34, 0x56, 0x78,
	}, NewOxmInPhysicalPort(0x12345678)},
}

func TestOxmInPortSize(t *testing.T) {
	o := NewOxmInPort(0x12345678)
	if o.Len() != 4+4 {
		t.Errorf("Toal length of %s: %d, want %d", reflect.TypeOf(o).Elem().Name(), o.Len(), 8)
	}

	if o.Length != 4 {
		t.Errorf("%s.Length = %d, want %d", reflect.TypeOf(o).Elem().Name(), o.Length, 4)
	}
}

func TestOxmRead(t *testing.T) {
	for _, v := range oxms {
		expected := v.Bytes
		oxm := v.Oxm
		checkMarshall(t, oxm, expected)
	}
}
