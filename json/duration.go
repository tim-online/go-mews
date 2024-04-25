package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"regexp"
	"strconv"
	"time"
)

// Duration represents an ISO8601 Duration
// https://en.wikipedia.org/wiki/ISO_8601#Durations
type Duration struct {
	Years  int
	Months int
	Weeks  int
	Days   int
	// Time Component
	Hours   int
	Minutes int
	Seconds int
}

var pattern = regexp.MustCompile(`P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<week>\d+)W)?((?P<day>\d+)D)?(T((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>\d+)S)?)?`)

// ParseISO8601 parses an ISO8601 duration string.
func ParseISO8601(from string) (Duration, error) {
	var match []string
	var d Duration

	if pattern.MatchString(from) {
		match = pattern.FindStringSubmatch(from)
	} else {
		return d, errors.New("could not parse duration string")
	}

	for i, name := range pattern.SubexpNames() {
		part := match[i]
		if i == 0 || name == "" || part == "" {
			continue
		}

		val, err := strconv.Atoi(part)
		if err != nil {
			return d, err
		}
		switch name {
		case "year":
			d.Years = val
		case "month":
			d.Months = val
		case "week":
			d.Weeks = val
		case "day":
			d.Days = val
		case "hour":
			d.Hours = val
		case "minute":
			d.Minutes = val
		case "second":
			d.Seconds = val
		default:
			return d, fmt.Errorf("unknown field %s", name)
		}
	}

	return d, nil
}

// IsZero reports whether d represents the zero duration, P0D.
func (d Duration) IsZero() bool {
	return d.Years == 0 && d.Months == 0 && d.Weeks == 0 && d.Days == 0 && d.Hours == 0 && d.Minutes == 0 && d.Seconds == 0
}

// HasTimePart returns true if the time part of the duration is non-zero.
func (d Duration) HasTimePart() bool {
	return d.Hours > 0 || d.Minutes > 0 || d.Seconds > 0
}

// Shift returns a time.Time, shifted by the duration from the given start.
//
// NB: Shift uses time.AddDate for years, months, weeks, and days, and so
// shares its limitations. In particular, shifting by months is not recommended
// unless the start date is before the 28th of the month. Otherwise, dates will
// roll over, e.g. Aug 31 + P1M = Oct 1.
//
// Week and Day values will be combined as W*7 + D.
func (d Duration) Shift(t time.Time) time.Time {
	if d.Years != 0 || d.Months != 0 || d.Weeks != 0 || d.Days != 0 {
		days := d.Weeks*7 + d.Days
		t = t.AddDate(d.Years, d.Months, days)
	}
	t = t.Add(d.timeDuration())
	return t
}

func (d Duration) timeDuration() time.Duration {
	var dur time.Duration
	dur = dur + (time.Duration(d.Hours) * time.Hour)
	dur = dur + (time.Duration(d.Minutes) * time.Minute)
	dur = dur + (time.Duration(d.Seconds) * time.Second)
	return dur
}

var tmpl = template.Must(template.New("duration").Parse(`P{{if .Y}}{{.Y}}Y{{end}}{{if .M}}{{.M}}M{{end}}{{if .W}}{{.W}}W{{end}}{{if .D}}{{.D}}D{{end}}{{if .HasTimePart}}T{{end }}{{if .TH}}{{.TH}}H{{end}}{{if .TM}}{{.TM}}M{{end}}{{if .TS}}{{.TS}}S{{end}}`))

// String returns an ISO8601-ish representation of the duration.
func (d Duration) String() string {
	var s bytes.Buffer

	if d.IsZero() {
		return "P0D"
	}

	err := tmpl.Execute(&s, d)
	if err != nil {
		panic(err)
	}

	return s.String()
}

// MarshalJSON satisfies json.Marshaler.
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (d *Duration) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	tmp, err := ParseISO8601(s)
	if err != nil {
		return err
	}
	*d = tmp

	return nil
}
