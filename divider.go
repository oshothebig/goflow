package goflow

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

type Divider interface {
	Divide(reader io.Reader) ([]byte, error)
}

type lengthBasedDivider struct {
	maxLength         int
	lengthFieldOffset int
	lengthFieldLength int
	lengthAdjustment  int
}

func (divider *lengthBasedDivider) minReadLength() int {
	return divider.lengthFieldOffset + divider.lengthFieldLength
}

func (divider *lengthBasedDivider) Divide(reader io.Reader) ([]byte, error) {
	buf := wrapReader(reader)
	header, err := buf.Peek(divider.minReadLength())
	if err != nil {
		return nil, err
	}

	length := divider.bytesToRead(header)
	data := make([]byte, length)
	n, err := buf.Read(data)
	if err != nil {
		return data[:n], err
	}

	return data, nil
}

func (divider *lengthBasedDivider) getLength(header []byte) int64 {
	b := header[divider.lengthFieldOffset:]
	return bigEndianValue(b)
}

func (divider *lengthBasedDivider) bytesToRead(header []byte) uint64 {
	return uint64(int64(divider.minReadLength()) + divider.getLength(header) + int64(divider.lengthAdjustment))
}

func bigEndianValue(b []byte) int64 {
	var length uint64
	if len(b) > 8 {
		binary.Read(bytes.NewReader(b[len(b)-8:]), binary.BigEndian, &length)
		return int64(length)
	}

	for i, v := range b {
		shift := uint(len(b)*8) - uint((i+1)*8)
		length |= uint64(v) << shift
	}

	return int64(length)
}

func wrapReader(r io.Reader) *bufio.Reader {
	buf, ok := r.(*bufio.Reader)
	if !ok {
		buf = bufio.NewReader(r)
	}

	return buf
}
