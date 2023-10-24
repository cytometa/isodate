package isodate

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {

	got, _ := Parse("2023-10-01")
	want := IsoDate{time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
