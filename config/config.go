package config

import (
	"log"
)

const DriverUrl = "http://localhost:4444"

// rename to gecko session related confir
type SessionConfig struct {
	Path string
	Id   string
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
