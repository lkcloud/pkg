package time

import (
	"encoding/json"
	"strings"
	"time"
)

type JsonTime time.Time

func (j *JsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*j = JsonTime(t)
	return nil
}

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

func (j JsonTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
