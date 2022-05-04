package browser_ts_test

import (
	"proto2http/target/browser_ts"
	"testing"
)

func TestFileExtension(t *testing.T) {
	b := browser_ts.New()

	if b.FileExtension() != "ts" {
		t.Errorf("Expected file extension to be 'ts', got '%s'", b.FileExtension())
	}
}
