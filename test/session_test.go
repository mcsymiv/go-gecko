package test

import (
	"fmt"
	"testing"

	"github.com/mcsymiv/go-stripe/config"
	"github.com/mcsymiv/go-stripe/session"
)

var sc *config.SessionConfig

// TestStartFirefox
// Open new firefox browser
func TestCreateSession(t *testing.T) {

	// Starts firefox browser
	// Sets new session id to SessionRepo
	sc = session.NewSession()

	if sc.Id == "" {
		t.Error("No session id")
	}
}

// TestGetStatus
// Prints info on remote to stdout
func TestGetStatus(t *testing.T) {
	rr, _ := session.DoRequest("get", fmt.Sprintf("%s%s", config.DriverUrl, "/status"), nil)

	fmt.Println(string(rr))

}

// TestCloseSession
// Closes remote session
func TestCloseSession(t *testing.T) {
	rr, _ := session.DoRequest("delete", fmt.Sprintf("%s%s/%s", config.DriverUrl, "/session", sc.Id), nil)

	fmt.Println(string(rr))
}
