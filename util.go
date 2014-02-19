package goflow

import (
	"bytes"
	"encoding/binary"
	"io"
)

const defaultAlign = 8

func alignedSize(size, alignment uint) uint {
	return (size + alignment - 1) / alignment * alignment
}

func marshalFixedSizeData(data interface{}, b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	if err = binary.Write(buf, binary.BigEndian, data); err != nil {
		return
	}

	if n, err = buf.Read(b); err != nil {
		return
	}
	return n, io.EOF
}

func unmarshalFixedSizeData(data interface{}, b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, data); err != nil {
		return
	}
	n += binary.Size(data)

	return
}
