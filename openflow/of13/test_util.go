package of13

import (
	"bytes"
	"encoding/binary"
	"io"
	"reflect"
	"testing"

	. "github.com/oshothebig/goflow/openflow"
)

func checkPacketizableSize(t *testing.T, p Packetizable, expected int) {
	data := make([]byte, expected)

	name := reflect.TypeOf(p).Elem().Name()
	actual := binary.Size(p)
	if actual != expected {
		t.Errorf("Size of %s: %d, want %d", name, actual, expected)
	}

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
