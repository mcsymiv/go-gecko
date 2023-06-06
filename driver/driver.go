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
