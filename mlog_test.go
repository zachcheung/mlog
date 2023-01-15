package mlog

import (
	"bytes"
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	const testString = "test"
	var b bytes.Buffer
	l := New()
	l.SetOutput(&b)
	l.SetPrefix("")
	l.SetFlags(0)
	l.EnableDebug()
	l.Debug(testString)
	expect := fmt.Sprintf("[%s] %s\n", colorize(gray, "DEBUG"), testString)
	if b.String() != expect {
		t.Errorf("log output should match %q is %q", expect, b.String())
	}
}
