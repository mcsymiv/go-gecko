package driver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

func New(capsFn ...CapabilitiesFunc) WebDriver {

	c := DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	rr, err := request.Do(http.MethodPost, path.Url(path.Session), data)
	if err != nil {
		fmt.Println(err)
	}

	// var res RemoteResponse
	res := struct{ Value NewSessionResponse }{}

	err = json.Unmarshal(rr, &res)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &Driver{
		Id: res.Value.SessionId,
	}
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
type CapabilitiesFunc func(*Capabilities)

// DefaultCapabilities
func DefaultCapabilities() Capabilities {
	return Capabilities{
		AlwaysMatch: AlwaysMatch{
			AcceptInsecureCerts: true,
			BrowserName:         "firefox",
		},
	}
}
