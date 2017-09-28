package models

import (
	"time"
	"fmt"
	"strings"
)

const dateTimeLayout = "2017-04-25T15:08:43.687Z"
var nilTime = (time.Time{}).UnixNano()

type DateTime struct {
	time.Time `bson:",inline"`
}

func NewDateTime() DateTime {
	res := DateTime{Time: time.Now()}
	return res
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format(dateTimeLayout))), nil
}

func (d *DateTime) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return
	}
	d.Time, err = time.Parse(dateTimeLayout, s)
	return
}

func (d *DateTime) IsSet() bool {
	return d.Time.UnixNano() != nilTime
}