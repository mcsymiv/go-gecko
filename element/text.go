package element

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/request"
)

// Text
// Returns an element’s text “as rendered”
func (e *Element) Text() (string, error) {
	url := request.UrlArgs(request.Session, e.SessionId, request.Element, e.Id, request.Text)
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
