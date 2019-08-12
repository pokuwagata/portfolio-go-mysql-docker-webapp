package logger

import (
	"bytes"
	"testing"
	"regexp"
)

const (
	Rdate = `[0-9][0-9][0-9][0-9]/[0-9][0-9]/[0-9][0-9]`
	Rtime = `[0-9][0-9]:[0-9][0-9]:[0-9][0-9]`
)

func TestErrorf(t *testing.T) {
	expectedMessage := "error"
	buf := new(bytes.Buffer)
	logger := NewLogger(buf)
	logger.Errorf(expectedMessage)

	pattern := "^\\[Error]" + Rdate + " " + Rtime + " " + expectedMessage + "\n"
	matched, err := regexp.Match(pattern, buf.Bytes())

	if err != nil {
		t.Fatalf("pattern %q did not compile: %s", pattern, err)
	}

	if !matched {
		t.Errorf("message did not match pattern. expected=%q, got=%q", expectedMessage, buf.Bytes())
	}
}

// TODO: 他のログ出力も必要になったらテストを追加する
