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

// CapabilitiesFunc Usage:
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
// Create driver.New(browserName("chrome"))
type CapabilitiesFunc func(*NewSessionCapabilities)

// DefaultCapabilities
func DefaultCapabilities() NewSessionCapabilities {
	return NewSessionCapabilities{
		BrowserCapabilities{
			AlwaysMatch{
				AcceptInsecureCerts: true,
				BrowserName:         "firefox",
			},
		},
	}
}

// Capabilities
func Capabilities(browsername string) NewSessionCapabilities {
	return NewSessionCapabilities{
		BrowserCapabilities{
			AlwaysMatch{
				AcceptInsecureCerts: true,
				BrowserName:         browsername,
			},
		},
	}
}

func ImplicitWait(w float32) CapabilitiesFunc {
	return func(cap *NewSessionCapabilities) {
		cap.Capabilities.AlwaysMatch.Timeouts.Implicit = w
	}
}

func Firefox(moz *MozOptions) CapabilitiesFunc {
	return func(cap *NewSessionCapabilities) {
		cap.Capabilities.AlwaysMatch.MozOptions = *moz
	}
}

func BrowserName(b string) CapabilitiesFunc {
  return func(cap *NewSessionCapabilities) {
    cap.Capabilities.AlwaysMatch.BrowserName = b
  }
}
