package models

import (
	"time"
	"fmt"
)

type xormDate time.Time

func (this xormDate) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
