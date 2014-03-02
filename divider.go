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

	length := uint64(divider.minReadLength()) + divider.getLength(header)
	data := make([]byte, length)
	n, err := buf.Read(data)
	if err != nil {
		return data[:n], err
	}

	return data, nil
}

func (divider *lengthBasedDivider) getLength(header []byte) uint64 {
	b := header[divider.lengthFieldOffset:]
	return Uint64(b)
}

func Uint64(b []byte) uint64 {
	var length uint64
	if len(b) > 8 {
		binary.Read(bytes.NewReader(b[len(b)-8:]), binary.BigEndian, &length)
		return length
	}

	for i, v := range b {
		shift := uint(len(b)*8) - uint((i+1)*8)
		length |= uint64(v) << shift
	}

	return length
}

func wrapReader(r io.Reader) *bufio.Reader {
	buf, ok := r.(*bufio.Reader)
	if !ok {
		buf = bufio.NewReader(r)
	}

	return buf
}
