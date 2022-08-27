package util_test

import (
	"testing"

	"github.com/michimani/jsc/util"
	"github.com/stretchr/testify/assert"
)

func Test_StringPtr(t *testing.T) {
	cases := []struct {
		name string
		v    string
	}{
		{"ok", "test"},
		{"ok: empty", ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			p := util.StringPtr(c.v)
			asst.Equal(c.v, *p)
		})
	}
}

func Test_IntPtr(t *testing.T) {
	cases := []struct {
		name string
		v    int
	}{
		{"ok", 100},
		{"ok: zero", 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			p := util.IntPtr(c.v)
			asst.Equal(c.v, *p)
		})
	}
}

func Test_BoolPtr(t *testing.T) {
	cases := []struct {
		name string
		v    bool
	}{
		{"ok: true", true},
		{"ok: false", false},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			p := util.BoolPtr(c.v)
			asst.Equal(c.v, *p)
		})
	}
}
