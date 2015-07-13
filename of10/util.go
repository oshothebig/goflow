package of10

import (
	"encoding/binary"
	"io"
)

func unmarshalFields(rd io.Reader, fields ...interface{}) error {
	for field := range fields {
		if err := binary.Read(rd, binary.BigEndian, field); err != nil {
			return err
		}
	}
	return nil
}
