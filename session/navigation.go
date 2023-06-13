package session

import (
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
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
func (s *Session) Open(u string) {
	st := strategy.NewRequester(&NavigationRequest{
		NavigationUrl: path.UrlArgs(path.Session, s.Id, path.UrlPath),
	})

	st.Post(map[string]string{
		"url": u,
	})
}

// GetUrl
func (s *Session) GetUrl() string {
	st := strategy.NewRequester(&NavigationRequest{
		NavigationUrl: path.UrlArgs(path.Session, s.Id, path.UrlPath),
	})
	return st.Get()
}
