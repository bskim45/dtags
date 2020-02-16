package quay

import (
	"fmt"
	"strings"
	"time"
)

var nilTime = (time.Time{}).UnixNano()

type QuayTime struct {
	time.Time
}

func (ct *QuayTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(time.RFC1123Z, s)
	return
}

func (ct *QuayTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(time.RFC1123Z))), nil
}
