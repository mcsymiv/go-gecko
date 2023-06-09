package element

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
	"github.com/mcsymiv/go-gecko/strategy"
)

type State struct {
	AttributeUrl string
}

func (e *Element) Attribute(a string) string {

	st := strategy.NewRequester(&State{
		AttributeUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Attribute, a),
	})
	return st.Get()
}

func (s *State) Url() string {
	return s.AttributeUrl
}

// Text
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
