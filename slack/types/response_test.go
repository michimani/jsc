package types_test

import (
	"testing"

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
