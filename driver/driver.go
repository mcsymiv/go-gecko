package driver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// New
// Connect to the WebDriver instance running locally
func New(capsFn ...CapabilitiesFunc) WebDriver {

	c := DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

	data, err := json.Marshal(c)
	log.Println(string(data))
	if err != nil {
		fmt.Println(err)
	}

	rr, err := request.Do(http.MethodPost, path.Url(path.Session), data)
	if err != nil {
		fmt.Println(err)
	}

	// var res RemoteResponse
	res := new(struct{ Value NewSessionResponse })

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
type CapabilitiesFunc func(*NewSessionCapabilities)

// DefaultCapabilities
func DefaultCapabilities() NewSessionCapabilities {
	return NewSessionCapabilities{
		Capabilities{
			AlwaysMatch{
				AcceptInsecureCerts: true,
				BrowserName:         "firefox",
			},
		},
	}
}
