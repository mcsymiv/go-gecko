package capabilities

type Capabilities struct {
	Port         string              `json:"-"`
	Host         string              `json:"-"`
	Capabilities BrowserCapabilities `json:"capabilities"`
}

type BrowserCapabilities struct {
	AlwaysMatch `json:"alwaysMatch"`
}

type AlwaysMatch struct {
	AcceptInsecureCerts bool   `json:"acceptInsecureCerts"`
	BrowserName         string `json:"browserName"`
	Timeouts            `json:"timeouts,omitempty"`
	MozOptions          `json:"moz:firefoxOptions,omitempty"`
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
	Log     `json:"log,omitempty"`
}

type Log struct {
	Level string `json:"level,omitempty"`
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
type CapabilitiesFunc func(*Capabilities)

// DefaultCapabilities
// Sets default firefox browser with local dev url
// With defined in service port, i.e. :4444
// Port and Host fields are used and passed to the WebDriver instance
// To reference and build current driver url
func DefaultCapabilities() Capabilities {
	return Capabilities{
		Port: "4444",
		Capabilities: BrowserCapabilities{
			AlwaysMatch{
				AcceptInsecureCerts: true,
				BrowserName:         "firefox",
			},
		},
	}
}

func ImplicitWait(w float32) CapabilitiesFunc {
	return func(cap *Capabilities) {
		cap.Capabilities.AlwaysMatch.Timeouts.Implicit = w
	}
}

func Firefox(moz *MozOptions) CapabilitiesFunc {
	return func(cap *Capabilities) {
		cap.Capabilities.AlwaysMatch.MozOptions = *moz
	}
}

func BrowserName(b string) CapabilitiesFunc {
	return func(cap *Capabilities) {
		cap.Capabilities.AlwaysMatch.BrowserName = b
	}
}

// func Port(p string) CapabilitiesFunc {
// 	return func(caps *Capabilities) {
// 		caps.Port = p
// 	}
// }

// func Host(h string) CapabilitiesFunc {
// 	return func(caps *Capabilities) {
// 		caps.Host = h
// 	}
// }
