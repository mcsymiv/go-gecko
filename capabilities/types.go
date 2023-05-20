package capabilities

type Capabilities interface {
	ImplicitWait(w float32) CapabilitiesFunc
}

type NewSessionCapabilities struct {
	Capabilities BrowserCapabilities `json:"capabilities"`
}

type BrowserCapabilities struct {
	AlwaysMatch AlwaysMatch `json:"alwaysMatch"`
}

type AlwaysMatch struct {
	AcceptInsecureCerts bool     `json:"acceptInsecureCerts"`
	BrowserName         string   `json:"browserName"`
	Timeouts            Timeouts `json:"timeouts,omitempty,-"`
}

type Timeouts struct {
	Implicit float32 `json:"implicit,omitempty,-"`
	PageLoad float32 `json:"pageLoad,omitempty,-"`
	Script   float32 `json:"script,omitempty,-"`
}
