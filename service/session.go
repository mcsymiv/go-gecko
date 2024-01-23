package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/capabilities"
)

type Session struct {
	SessionId    string                 `json:"sessionId"`
	Capabilities map[string]interface{} `json:"-"`
}

type DriverStatus struct {
	Message string `json:"message"`
	Ready   bool   `json:"ready"`
}

// NewSession
// Return Session response with session Id
func NewSession(caps *capabilities.Capabilities) (*Session, error) {
	data, err := json.Marshal(caps)
	if err != nil {
		log.Printf("new driver marshall error: %+v", err)
		return nil, err
	}
	url := fmt.Sprintf("http://%s:%s/session", caps.Host, caps.Port)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "json/application")

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

	reply := new(struct{ Value Session })
	if err := json.Unmarshal(body, reply); err != nil {
		log.Println("Status unmarshal error", err)
		return &reply.Value, err
	}

	return &reply.Value, nil
}

// GetStatus
// Status returns information about whether a remote end is
// in a state in which it can create new sessions,
// but may additionally include arbitrary meta information
// that is specific to the implementation.
func GetStatus(caps *capabilities.Capabilities) (*DriverStatus, error) {
	url := fmt.Sprintf("http://%s:%s/status", caps.Host, caps.Port)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "json/application")
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

	reply := new(struct{ Value DriverStatus })
	if err := json.Unmarshal(body, reply); err != nil {
		log.Println("Status unmarshal error", err)
		return &reply.Value, err
	}

	return &reply.Value, nil
}
