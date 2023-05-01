package handlers

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Creates session config
	// Sets /session path
	// Acts as before test setup
	os.Exit(m.Run())
}
