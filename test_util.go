package goflow

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func checkPacketizableSize(t *testing.T, p Packetizable, expected int) {
	name := reflect.TypeOf(p).Elem().Name()

	actual := p.Len()
	if actual != uint(expected) {
		t.Errorf("%s.Len(): %d, want %d", name, actual, expected)
	}

	data := make([]byte, expected)
	n, _ := p.Read(data)
	if n != expected {
		t.Errorf("Size of %s (serialized): %d, want %d", name, n, expected)
	}
}

func checkMarshall(t *testing.T, p Packetizable, expected []byte) {
	actual := make([]byte, p.Len())

	name := reflect.TypeOf(p).Elem().Name()

	n, err := p.Read(actual)
	if n != len(expected) {
		t.Errorf("Len(): %d, want %d", n, len(expected))
	}
	if err != io.EOF {
		t.Error("EOF not exist")
	}
	if !bytes.Equal(actual, expected) {
		t.Errorf("Bytes of %s not matched", name)
		t.Logf("Actual  : %x\n", actual)
		t.Logf("Expected: %x\n", expected)
	}
}

func checkUnmarshal(t *testing.T, b []byte, empty, expected Packetizable) {
	actual := empty
	actual.Write(b)

	name := reflect.TypeOf(actual).Elem().Name()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Fields of %s not matched", name)
		t.Logf("Actual  :%v\n", actual)
		t.Logf("Expected:%v\n", expected)
	}
}
