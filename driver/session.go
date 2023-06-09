package driver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// New
// Connect to the WebDriver instance running locally
func New(capsFn ...capabilities.CapabilitiesFunc) (WebDriver, error) {

	c := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

	data, err := json.Marshal(c)
	if err != nil {
		log.Printf("Marshal capabilities: %+v", err)
		return nil, err
	}
	log.Printf("caps 2: %+v", string(data))

	r, err := request.Do(http.MethodPost, path.Url(path.Session), data)
	if err != nil {
		log.Printf("Connect to driver instance with capabilities: %+v", err)
		return nil, err
	}

	res := new(struct{ Value NewSessionResponse })
	err = json.Unmarshal(r, &res)
	if err != nil {
		log.Printf("Unmarshal capabilities: %+v", err)
		return nil, err
	}

	return &Driver{
		Id: res.Value.SessionId,
	}, nil
}

// GetStatus
//
// Status returns information about whether a remote end is
// in a state in which it can create new sessions,
// but may additionally include arbitrary meta information
// that is specific to the implementation.
func GetStatus() (*Status, error) {
	rr, err := request.Do(http.MethodGet, path.Url(path.Status), nil)
	if err != nil {
		log.Println("Status request error", err)
		return nil, err
	}

	reply := new(struct{ Value Status })
	if err := json.Unmarshal(rr, reply); err != nil {
		log.Println("Status unmarshal error", err)
		return &reply.Value, err
	}

	return &reply.Value, nil
}

// Closes session
func (d *Driver) Quit() {
	request.Do(http.MethodDelete, path.UrlArgs(path.Session, d.Id), nil)
}