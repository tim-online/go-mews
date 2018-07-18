package customers

import (
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return err
	}

	d.Time, err = time.Parse(time.RFC3339, value)
	return err
}
