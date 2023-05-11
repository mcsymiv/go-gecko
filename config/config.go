package config

import (
	"github.com/mcsymiv/go-gecko/models"
)

// Global config struct
type SessionConfig struct {
	Id string
}

type RemoteResponse struct {
	Value Value
}

type Value struct {
	SessionId    string                 `json:"sessionId"`
	Capabilities map[string]interface{} `json:"-"`
}

// Functional Options for gecko remote Capabilities
// Usage:
//
// For the capabilities set with argument:
//
//	func browserName(s string) CapabilitiesFunc {
//	 return func(cap *models.Capabilities) {
//	   cap.BrowserName = s
//	 }
//	}
//
// For the capabilities:
//
//	func acceptInsecure(cap *models.Capabilities) {
//	  cap.AcceptInsecureCerts = false
//	}
//
// Example:
// Create session.New(browserName("chrome")
type CapabilitiesFunc func(*models.Capabilities)

// DefaultCapabilities
func DefaultCapabilities() models.Capabilities {
	return models.Capabilities{
		AlwaysMatch: models.AlwaysMatch{
			AcceptInsecureCerts: true,
			BrowserName:         "firefox",
		},
	}
}
