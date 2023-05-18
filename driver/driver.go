package driver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// New
// Connect to the WebDriver instance running locally
func New(capsFn ...CapabilitiesFunc) (WebDriver, error) {

	c := DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

	data, err := json.Marshal(c)
	if err != nil {
		log.Println("Marshal capabilities error", err)
		return nil, err
	}

	r, err := request.Do(http.MethodPost, path.Url(path.Session), data)
	if err != nil {
		log.Println("Connect to driver instance with capabilities error", err)
		return nil, err
	}

	res := new(struct{ Value NewSessionResponse })
	err = json.Unmarshal(r, &res)
	if err != nil {
		log.Println("Unmarshal capabilities error", err)
		return nil, err
	}

	return &Driver{
		Id: res.Value.SessionId,
	}, nil
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
