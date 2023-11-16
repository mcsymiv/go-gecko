package element

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/request"
)

// Attribute
// Returns elements attribute value
func (e *Element) Attribute(a string) (string, error) {
	url := request.UrlArgs(request.Session, e.SessionId, request.Element, e.Id, request.Attribute, a)
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
