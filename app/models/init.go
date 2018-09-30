package models

import (
	"time"
	"fmt"
)

type XormDate time.Time

func (this XormDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}