package element

import (
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
)

// StateRequest
// Used for ContextRequester
type StateRequest struct {
	StateRequestUrl string
}

// Url
// Requester method
func (s *StateRequest) Url() string {
	return s.StateRequestUrl
}

// Attribute
// Returns elements attribute value
func (e *Element) Attribute(a string) string {

	st := strategy.NewRequester(&StateRequest{
		StateRequestUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Attribute, a),
	})
	return st.Get()
}

// Text
// Returns an element’s text “as rendered”
func (e *Element) Text() string {

	st := strategy.NewRequester(&StateRequest{
		StateRequestUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Text),
	})
	return st.Get()
}
