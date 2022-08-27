package types_test

import (
	"testing"
	"time"

	"github.com/michimani/jsc/slack/types"

	"github.com/stretchr/testify/assert"
)

func Test_TsString_String(t *testing.T) {
	testTs := types.TsString("42")

	cases := []struct {
		name   string
		ts     *types.TsString
		expect string
	}{
		{"ok", &testTs, "42"},
		{"ok: nil", nil, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			s := c.ts.String()
			asst.Equal(c.expect, s)
		})
	}
}

func Test_TsString_Float64(t *testing.T) {
	okTs := types.TsString("42.42")
	ngTs := types.TsString("not-float")

	cases := []struct {
		name    string
		ts      *types.TsString
		expect  float64
		wantErr bool
	}{
		{"ok", &okTs, 42.42, false},
		{"ok: nil", nil, 0, false},
		{"ng: not float string", &ngTs, 0, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			f, err := c.ts.Float64()

			if c.wantErr {
				asst.Error(err)
				asst.Equal(float64(0), f)
				return
			}

			asst.Nil(err)
			asst.Equal(c.expect, f)
		})
	}
}

func Test_TsString_ToID(t *testing.T) {
	okTs := types.TsString("42.42")

	cases := []struct {
		name   string
		ts     *types.TsString
		expect string
	}{
		{"ok", &okTs, "p4242"},
		{"ok: nil", nil, ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			f := c.ts.ToID()
			asst.Equal(c.expect, f)
		})
	}
}

func Test_TsString_ToTime(t *testing.T) {
	ngTs := types.TsString("not-float")
	okTs := types.TsString("1611802233.000400")
	okTime := time.Date(2021, time.January, 28, 2, 50, 33, 0, time.UTC).Local()

	cases := []struct {
		name    string
		ts      *types.TsString
		expect  *time.Time
		wantErr bool
	}{
		{"ok", &okTs, &okTime, false}, // 2021-01-28T02:50:33Z
		{"ok: nil", nil, nil, false},
		{"ng: not timestamp string", &ngTs, nil, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			t, err := c.ts.ToTime()

			if c.wantErr {
				asst.Error(err)
				asst.Nil(t)
				return
			}

			asst.NoError(err)
			asst.Equal(c.expect, t)
		})
	}
}
