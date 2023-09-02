package session

import (
	"github.com/mcsymiv/go-gecko/element"
)

// WebDriver
// https://w3c.github.io/webdriver/
type WebDriver interface {
	// Navigation
	Open(u string)
	GetUrl() string

	// Session
	// GetStatus() (*Status, error)
	Quit()

	// Elemnents
	Init(b, v string) element.WebElement
	FindElement(b, v string) (element.WebElement, error)
	FindElements(b, v string) (element.WebElements, error)

	// Document
	ExecuteScriptSync(s string, args ...interface{}) error
	PageSource() string
}

// BrowserCapabilities
type BrowserCapabilities interface {
	ImplilcitWait(w float32)
}

// Session
// Represents WebDriver
// Holds session Id
// Driver port
type Session struct {
	Id   string
	Port string
}

// Status response
// W3C type
type Status struct {
	Message string `json:"message"`
	Ready   bool   `json:"ready"`
}

// NewSessionResponse
// W3C type
type NewSessionResponse struct {
	SessionId    string                 `json:"sessionId"`
	Capabilities map[string]interface{} `json:"-"`
}
