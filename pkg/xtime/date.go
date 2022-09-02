package xtime

import (
	"bytes"
	"encoding/json"
	"time"
)

const ISO8601Date = "2006-01-02"

type Date struct {
	time.Time
}

func (d Date) format() string {
	return d.Time.Format(ISO8601Date)
}

func (d Date) String() string {
	return d.format()
}

// UnmarshalJSON converts a byte array into a Date
func (d *Date) UnmarshalJSON(text []byte) error {
	if string(text) == "null" {
		// Nulls are converted to zero times
		var zero Date
		*d = zero
		return nil
	}
	b := bytes.NewBuffer(text)
	dec := json.NewDecoder(b)
	var s string
	if err := dec.Decode(&s); err != nil {
		return err
	}
	value, err := time.ParseInLocation(ISO8601Date, s, time.Local)
	if err != nil {
		return err
	}
	d.Time = value
	return nil
}

// MarshalJSON returns the JSON output of a Date.
// Null will return a zero value date.
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + d.format() + `"`), nil
}
