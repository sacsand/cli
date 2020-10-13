package cmd

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Helpers_FormatLatency(t *testing.T) {
	t.Parallel()

	cases := []struct {
		d        time.Duration
		expected time.Duration
	}{
		{time.Millisecond * 123456, time.Millisecond * 123450},
		{time.Millisecond * 12340, time.Millisecond * 12340},
		{time.Microsecond * 123456, time.Microsecond * 123450},
		{time.Microsecond * 123450, time.Microsecond * 123450},
		{time.Nanosecond * 123456, time.Nanosecond * 123450},
		{time.Nanosecond * 123450, time.Nanosecond * 123450},
		{time.Nanosecond * 123, time.Nanosecond * 123},
	}

	for _, tc := range cases {
		t.Run(tc.d.String(), func(t *testing.T) {
			assert.Equal(t, formatLatency(tc.d), tc.expected)
		})
	}
}

func Test_Helper_Replace(t *testing.T) {
	at := assert.New(t)

	dir, err := ioutil.TempDir("", "test_helper_replace")
	at.Nil(err)
	defer func() {
		at.Nil(os.RemoveAll(dir))
	}()

	f, err := ioutil.TempFile(dir, "*.go")
	at.Nil(err)
	at.Nil(f.Close())

	at.Nil(replace(dir, "*.go", "old", "new"))
}