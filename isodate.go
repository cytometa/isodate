package isodate

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type IsoDate struct {
	Time time.Time
}

func New(t time.Time) IsoDate {
	i := IsoDate{
		Time: time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()),
	}
	return i
}

// Parse to create a date from a string
func Parse(value string) (IsoDate, error) {
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return IsoDate{}, err
	}
	return IsoDate{Time: t}, nil
}

// UnmarshalJSON to read a JSON
func (d *IsoDate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	t, _ := time.Parse("2006-01-02", s)
	d.Time = t
	return nil
}

// MarshalJSON to create a JSON
func (d IsoDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

// Scan to read from SQL into variable, implements sql.Scan interface
func (d *IsoDate) Scan(value interface{}) error {
	d.Time = value.(time.Time)
	return nil
}

// Value to write IsoDate value to SQL, implements sql.driver.Value interface
func (d IsoDate) Value() (driver.Value, error) {
	return d.Time.Format("2006-01-02"), nil
}

// String to return string for IsoDate value, for example in HTML templates
func (d IsoDate) String() string {
	return d.Time.Format("2006-01-02")
}

// Format to format dates in a custom way
func (d IsoDate) Format(layout string) string {
	return d.Time.Format(layout)
}
