package element

import (
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
)

// State
// Used for ContextRequester
type State struct {
	StateUrl string
}

// Url
// Requester method
func (s *State) Url() string {
	return s.StateUrl
}

// Attribute
// Returns elements attribute value
func (e *Element) Attribute(a string) string {

	st := strategy.NewRequester(&State{
		StateUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Attribute, a),
	})
	return st.GetDefault()
}

// Text
// Returns an element’s text “as rendered”
func (e *Element) Text() string {

	st := strategy.NewRequester(&State{
		StateUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Text),
	})
	return st.GetDefault()
}
