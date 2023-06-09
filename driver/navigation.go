package driver

import (
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
)

type Navigation struct {
	NavigationUrl string
}

func (n *Navigation) Url() string {
	return n.NavigationUrl
}

// Open
// Goes to url
func (d *Driver) Open(u string) {
	st := strategy.NewRequester(&Navigation{
		NavigationUrl: path.UrlArgs(path.Session, d.Id, path.UrlPath),
	})

	st.Post(map[string]string{
		"url": u,
	})
}

// GetUrl
func (d *Driver) GetUrl() string {
	st := strategy.NewRequester(&Navigation{
		NavigationUrl: path.UrlArgs(path.Session, d.Id, path.UrlPath),
	})
	return st.Get()
}
