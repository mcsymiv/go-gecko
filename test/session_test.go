package handlers

import (
	"testing"

	"github.com/mcsymiv/go-stripe/session"
)

// TestStartFirefox
// Open new firefox browser
func TestStartFirefox(t *testing.T) {

	// Starts firefox browser
	// Sets new session id to SessionRepo
	session.SessionRepo.NewSession()

	if session.SessionRepo.Config.Id == "" {
		t.Error("No session id")
	}
}
