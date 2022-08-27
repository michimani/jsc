package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TsString is type of timestamp string.
type TsString string

func (ts *TsString) String() string {
	if ts == nil {
		return ""
	}

	return string(*ts)
}

func (ts *TsString) Float64() (float64, error) {
	if ts == nil {
		return 0, nil
	}

	f, err := strconv.ParseFloat(ts.String(), 64)
	if err != nil {
		return 0, err
	}

	return f, nil
}

func (ts *TsString) ToID() string {
	if ts == nil {
		return ""
	}

	return fmt.Sprintf("p%s", strings.ReplaceAll(ts.String(), ".", ""))
}

// ignore micro sec
func (ts *TsString) ToTime() (*time.Time, error) {
	if ts == nil {
		return nil, nil
	}

	f, err := ts.Float64()
	if err != nil {
		return nil, err
	}

	t := time.Unix(int64(f), 0)
	return &t, nil
}
