package element

import (
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
)

type State struct {
	StateUrl string
}

func (s *State) Url() string {
	return s.StateUrl
}

func (e *Element) Attribute(a string) string {

	st := strategy.NewRequester(&State{
		StateUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Attribute, a),
	})
	return st.Get()
}

// Text
func (e *Element) Text() string {

	st := strategy.NewRequester(&State{
		StateUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Text),
	})
	return st.Get()
}
