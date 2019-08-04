package time

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	defaultDateTimeFormat = "2006-01-02 15:04:05"
)

// DatabaseTime format json time field by myself
type DatabaseTime struct {
	time.Time
}

// MarshalJSON on DatabaseTime format Time field with %Y-%m-%d %H:%M:%S
func (t DatabaseTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(defaultDateTimeFormat))
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t DatabaseTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *DatabaseTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DatabaseTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func ToDatabaseTime(str string) (DatabaseTime, error) {
	var jt DatabaseTime
	loc, _ := time.LoadLocation("Local")
	value, err := time.ParseInLocation(defaultDateTimeFormat, str, loc)
	if err != nil {
		return jt, err
	}
	return DatabaseTime{
		Time: value,
	}, nil
}

func Now() DatabaseTime {
	return DatabaseTime{
		Time: time.Now(),
	}
}
