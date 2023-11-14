package session

import (
	"encoding/json"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
	"log"
	"net/http"
)

// Navigation
// ContextRequester for driver navigation actions
type NavigationRequest struct {
	NavigationUrl string
}

// Url
// Requester method
func (n *NavigationRequest) Url() string {
	return n.NavigationUrl
}

// Open
// Goes to url
func (s *Session) Open(u string) error {
	data, err := json.Marshal(map[string]string{
		"url": u,
	})
	if err != nil {
		log.Printf("Open marshal: %+v", err)
		return err
	}

	url := path.UrlArgs(path.Session, s.Id, path.UrlPath)
	rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Open request error: %+v", err)
		return err
	}
	res := new(struct{ Value []map[string]string })
	if err := json.Unmarshal(rr, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return err
	}

	return nil
}

// GetUrl
func (s *Session) GetUrl() (string, error) {
	url := path.UrlArgs(path.Session, s.Id, path.UrlPath)
	rr, err := request.Do(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Open request error: %+v", err)
		return "", err
	}
	val := new(struct{ Value string })
	err = json.Unmarshal(rr, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return "", nil
	}

	return val.Value, nil
}
