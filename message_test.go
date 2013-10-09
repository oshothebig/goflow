package goflow

import (
	"encoding/binary"
	"testing"
)

func TestXidGeneration(t *testing.T) {
	var xid uint32

	for expected := uint32(1); expected < 11; expected++ {
		xid = generateXid()

		if xid != expected {
			t.Errorf("generateXid() = %d, want %d", xid, expected)
		}
	}
}

func BenchmarkNewBarrierRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBarrierRequest()
	}
}

func BenchmarkNewBarrierReply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewBarrierReply()
	}
}

func TestRoleRequestSize(t *testing.T) {
	m := NewRoleRequest(ControllerRoles.Equal, 0)
	if binary.Size(m) != int(m.Length) {
		t.Errorf("size = %d, want: %d", m.Length, binary.Size(m))
	}
}

func TestRoleReplySize(t *testing.T) {
	m := NewRoleReply(ControllerRoles.Equal, 0)
	if binary.Size(m) != int(m.Length) {
		t.Errorf("size = %d, want: %d", m.Length, binary.Size(m))
	}
}

func BenchmarkNewRoleRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRoleRequest(ControllerRoles.Equal, 0)
	}
}

func BenchmarkNewRoleReply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewRoleReply(ControllerRoles.Equal, 0)
	}
}

func BenchmarkNewGetAsyncRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewGetAsyncRequest()
	}
}
