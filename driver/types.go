package driver

import (
	"github.com/mcsymiv/go-gecko/element"
)

type WebDriver interface {
	GetStatus() (*Status, error)
	Open(u string) (string, error)
	Quit()

	FindElement(b, v string) (element.WebElement, error)
}

type BrowserCapabilities interface {
	ImplilcitWait(w float32)
}

type Driver struct {
	Id string
}

// Status response
type Status struct {
	Message string `json:"message"`
	Ready   bool   `json:"ready"`
}

type NewSessionResponse struct {
	SessionId    string                 `json:"sessionId"`
	Capabilities map[string]interface{} `json:"-"`
}
