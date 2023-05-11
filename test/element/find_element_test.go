package element

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/request"
	"github.com/mcsymiv/go-gecko/session"
)

var st SessionTest

type SessionTest struct {
	Id string
}

func TestFindElement(t *testing.T) {

	// Starts firefox browser
	sc := session.New()
	st := &SessionTest{
		Id: sc.Id,
	}
	defer st.CloseSession(t)

	// test steps
	st.GetSessionStatus(t)
	st.OpenUrl(t)
	st.FindElement(t)
}

func (st *SessionTest) FindElement(t *testing.T) {
	element.Find(st.Id)
}

// TestGetStatus
// Prints info on remote to stdout
func (st *SessionTest) GetSessionStatus(t *testing.T) {
	rr, _ := request.Do(http.MethodGet, request.Url(request.Status), nil)

	if string(rr) == "" {
		t.Errorf("Session status error")
	}
}

// OpenUrl
// Goes to url
func (st *SessionTest) OpenUrl(t *testing.T) {
	url := []byte(`{"url": "https://google.com"}`)

	_, _ = request.Do(http.MethodPost, request.UrlArgs(request.Session, st.Id, request.UrlPath), url)
	rr, _ := request.Do(http.MethodGet, request.UrlArgs(request.Session, st.Id, request.UrlPath), url)

	fmt.Println(string(rr))
}

// CloseSession
func (st *SessionTest) CloseSession(t *testing.T) {
	request.Do(http.MethodDelete, request.UrlArgs(request.Session, st.Id), nil)
}
