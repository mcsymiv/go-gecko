package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mcsymiv/go-stripe/config"
)

const JsonContentType = "application/json"

var SessionRepo *SessionRepository

type SessionRepository struct {
	Config *config.SessionConfig
}

func CreateSessionRepository() {
	SessionRepo = &SessionRepository{
		Config: &config.SessionConfig{
			Path: "/session",
		},
	}
}

func (sr *SessionRepository) NewSession() {

	var jsonStr = []byte(`{"capabilities": {"alwaysMatch": {"acceptInsecureCerts": true}}}`)

	rr, err := DoRequest("POST", fmt.Sprintf("%s%s", config.DriverUrl, sr.Config.Path), jsonStr)
	if err != nil {
		fmt.Println(err)
	}

	var res config.RemoteResponse
	err = json.Unmarshal(rr, &res)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("%+v", res.Value.SessionId)
	sr.Config.Id = res.Value.SessionId
}

func DoRequest(method, url string, data []byte) (json.RawMessage, error) {
	req, err := NewRequest(strings.ToUpper(method), url, data)
	if err != nil {
		return nil, err
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// NewRequest creates and returns http.Request
// Separetes request logic into func as convenience method
func NewRequest(method, url string, data []byte) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Accept", JsonContentType)

	return request, nil
}
