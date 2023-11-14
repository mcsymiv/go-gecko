package element

import (
	"encoding/json"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
	"log"
	"net/http"
)

// Attribute
// Returns elements attribute value
func (e *Element) Attribute(a string) (string, error) {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Attribute, a)
	r, err := request.Do(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Get attribute: %+v", err)
		return "", err
	}

	attr := new(struct{ Value string })
	err = json.Unmarshal(r, attr)
	if err != nil {
		log.Printf("Marshal attribute: %+v", err)
		return "", nil
	}

	return attr.Value, nil
}

// Text
// Returns an element’s text “as rendered”
func (e *Element) Text() (string, error) {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Text)
	r, err := request.Do(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Get text: %+v", err)
		return "", err
	}

	text := new(struct{ Value string })
	err = json.Unmarshal(r, text)
	if err != nil {
		log.Printf("Marshal text: %+v", err)
		return "", nil
	}

	return text.Value, nil
}

