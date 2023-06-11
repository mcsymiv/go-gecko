package element

import (
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
)

// Interaction
// ContextRequester for element interation actions
type Interaction struct {
	InteractionUrl string
}

// Url
// Requester method
func (i *Interaction) Url() string {
	return i.InteractionUrl
}

// Click
func (e *Element) Click() {

	st := strategy.NewRequester(&Interaction{
		InteractionUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Click),
	})

	// Empty struct is used to avoid gecko driver bug
	// That is thrown when POST is used without body
	// Click driver endpoint requires no body
	st.Post(&Empty{})
}

// SendKeys
func (e *Element) SendKeys(s string) {
	st := strategy.NewRequester(&Interaction{
		InteractionUrl: path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Value),
	})

	st.Post(&SendKeys{
		Text: s,
	})
}
