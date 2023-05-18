package driver

import (
	"github.com/mcsymiv/go-gecko/element"
)

type WebDriver interface {
	// Status()
	Open(u string)
	Quit()

	FindElement(b, v string) element.WebElement
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

type Capabilities struct {
	AlwaysMatch
}

type AlwaysMatch struct {
	AcceptInsecureCerts bool
	BrowserName         string
}
