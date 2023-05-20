package capabilities

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

func ImplicitWait(w float32) CapabilitiesFunc {
	return func(cap *NewSessionCapabilities) {
		cap.Capabilities.AlwaysMatch.Timeouts.Implicit = w
	}
}
