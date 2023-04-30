package handlers

import (
	"os"
	"testing"

	"github.com/mcsymiv/go-stripe/session"
)

func TestMain(m *testing.M) {
	// Creates session config
	// Sets /session path
	// Acts as before test setup
	session.CreateSessionRepository()
	os.Exit(m.Run())
}
