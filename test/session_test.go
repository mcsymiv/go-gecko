package test

import (
	"fmt"
	"testing"

	"github.com/mcsymiv/go-stripe/config"
	"github.com/mcsymiv/go-stripe/session"
)

var st SessionTest

type SessionTest struct {
	Id string
}

func TestSession(t *testing.T) {

	// Starts firefox browser
	// Sets new session id to SessionRepo
	sc := session.NewSession()
	st := &SessionTest{
		Id: sc.Id,
	}

	// test steps
	st.GetSessionStatus(t)
	st.OpenUrl(t)
	st.CloseSession(t)
}

// TestGetStatus
// Prints info on remote to stdout
func (st *SessionTest) GetSessionStatus(t *testing.T) {
	rr, _ := session.DoRequest("get", fmt.Sprintf("%s%s", config.DriverUrl, "/status"), nil)

	if string(rr) == "" {
		t.Errorf("Session status error")
	}
}

// OpenUrl
// Goes to url
func (st *SessionTest) OpenUrl(t *testing.T) {
	url := []byte(`{"url": "https://google.com"}`)

	_, _ = session.DoRequest("post", fmt.Sprintf("%s/session/%s/url", config.DriverUrl, st.Id), url)
	rr, _ := session.DoRequest("get", fmt.Sprintf("%s/session/%s/url", config.DriverUrl, st.Id), url)

	fmt.Println(string(rr))
}

// CloseSession
func (st *SessionTest) CloseSession(t *testing.T) {
	session.DoRequest("delete", fmt.Sprintf("%s%s/%s", config.DriverUrl, "/session", st.Id), nil)
}
