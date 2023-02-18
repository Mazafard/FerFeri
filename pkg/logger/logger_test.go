package logger

import (
	"bytes"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestLogger(t *testing.T) {
	// Redirect log output to a buffer
	var buf bytes.Buffer
	output := zerolog.ConsoleWriter{Out: &buf}
	log = zerolog.New(output).With().Timestamp().Logger()

	// Test Infof
	Infof("This is an info message: %d", 42)
	expected := "This is an info message: 42\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Infof output does not contain expected message. Got %q, expected %q", buf.String(), expected)
	}
	buf.Reset()

	// Test Debugf
	Debugf("This is a debug message: %d", 42)
	expected = "This is a debug message: 42\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Debugf output does not contain expected message. Got %q, expected %q", buf.String(), expected)
	}
	buf.Reset()

	// Test Warnf
	Warnf("This is a warning message: %d", 42)
	expected = "This is a warning message: 42\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Warnf output does not contain expected message. Got %q, expected %q", buf.String(), expected)
	}
	buf.Reset()

	// Test Errorf
	Errorf("This is an error message: %d", 42)
	expected = "This is an error message: 42\n"
	if !strings.Contains(buf.String(), expected) {
		t.Errorf("Errorf output does not contain expected message. Got %q, expected %q", buf.String(), expected)
	}
	buf.Reset()

	// Test Panicf
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Panicf did not panic as expected")
		}
	}()
	Panicf("This is a panic message: %d", 42)
}
