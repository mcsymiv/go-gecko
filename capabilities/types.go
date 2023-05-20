package capabilities

type NewSessionCapabilities struct {
	Capabilities BrowserCapabilities `json:"capabilities"`
}

type BrowserCapabilities struct {
	AlwaysMatch AlwaysMatch `json:"alwaysMatch"`
}

type AlwaysMatch struct {
	AcceptInsecureCerts bool       `json:"acceptInsecureCerts"`
	BrowserName         string     `json:"browserName"`
	Timeouts            Timeouts   `json:"timeouts,omitempty"`
	MozOptions          MozOptions `json:"moz:firefoxOptions,omitempty"`
}

type Timeouts struct {
	Implicit float32 `json:"implicit,omitempty"`
	PageLoad float32 `json:"pageLoad,omitempty"`
	Script   float32 `json:"script,omitempty"`
}

type MozOptions struct {
	Profile string   `json:"profile,omitempty"`
	Binary  string   `json:"binary,omitempty"`
	Args    []string `json:"args,omitempty"`
}
