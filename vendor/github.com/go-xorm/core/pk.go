package core

import (
	"bytes"
	"encoding/gob"
)

type PK []interface{}

func NewPK(pks ...interface{}) *PK {
	p := PK(pks)
	return &p
}

func (p *PK) ToString() (string, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(*p)
	return buf.String(), err
}

func (p *PK) FromString(content string) error {
	dec := gob.NewDecoder(bytes.NewBufferString(content))
	err := dec.Decode(p)
	return err
}
