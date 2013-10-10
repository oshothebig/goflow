package goflow

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/oshothebig/bingo"
)

const defaultAlign = 8

func alignedSize(size, alignment uint) uint {
	return (size + alignment - 1) / alignment * alignment
}

func unmarshalFixedSizeData(data interface{}, b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, data); err != nil {
		return
	}
	n += binary.Size(data)

	return
}

func marshal(data interface{}, b []byte) (n int, err error) {
	v, err := bingo.Marshal(data)
	if err != nil {
		return len(v), err
	}
	buf := bytes.NewBuffer(v)
	n, err = buf.Read(b)
	if err != nil {
		return
	}
	return n, io.EOF
}
