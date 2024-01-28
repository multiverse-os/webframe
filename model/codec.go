package model

import (
	"bytes"
	"encoding/gob"
)

type Instance interface{}

func Encode(obj interface{}) ([]byte, error) {
	byteBuffer := new(bytes.Buffer)
	if err := gob.NewEncoder(byteBuffer).Encode(obj); err != nil {
		return nil, err
	}
	return byteBuffer.Bytes(), nil
}

func Decode(data []byte, value interface{}) error {
	buf := bytes.NewBuffer(data)
	return gob.NewDecoder(buf).Decode(value)
}
