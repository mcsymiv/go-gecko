package config

import (
	"log"
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
