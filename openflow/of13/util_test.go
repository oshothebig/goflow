package of13

import "testing"

func checkAlignedSize(t *testing.T, size, align, expected uint) {
	if alignedSize(size, align) != expected {
		t.Errorf("%d bytes fitted to %d byte alignment: %d, want %d", size, align, expected)
	}
}

func TestAlignedSize(t *testing.T) {
	checkAlignedSize(t, 1, 8, 8)
	checkAlignedSize(t, 2, 8, 8)
	checkAlignedSize(t, 3, 8, 8)
	checkAlignedSize(t, 4, 8, 8)
	checkAlignedSize(t, 5, 8, 8)
	checkAlignedSize(t, 6, 8, 8)
	checkAlignedSize(t, 7, 8, 8)
	checkAlignedSize(t, 8, 8, 8)
	checkAlignedSize(t, 9, 8, 16)
	checkAlignedSize(t, 10, 8, 16)
}
