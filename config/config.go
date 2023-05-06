package config

import (
	"log"

	"github.com/mcsymiv/go-gecko/models"
)

const DriverUrl = "http://localhost:4444"

// SessionConfig
// Global config struct
type SessionConfig struct {
	Id string
}

type LogginConfig struct {
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

type RemoteResponse struct {
	Value Value
}

type Value struct {
	SessionId    string                 `json:"sessionId"`
	Capabilities map[string]interface{} `json:"-"`
}

// Functional Options for gecko remote Capabilities
type CapabilitiesFunc func(*models.Capabilities)

// DefaultCapabilities
// Sets default Capabilities for session
// If none is provided in session.New
func DefaultCapabilities() models.Capabilities {
	return models.Capabilities{
		AlwaysMatch: models.AlwaysMatch{
			AcceptInsecureCerts: true,
			BrowserName:         "firefox",
		},
	}
}
