package session

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
	"github.com/mcsymiv/go-gecko/strategy"
)

type DriverRequest struct {
	DriverUrl string
}

func (dr *DriverRequest) Url() string {
	return dr.DriverUrl
}

const GeckoDriverPath = "/Users/mcs/Development/tools/geckodriver"

// NewDriver
func NewDriver(capsFn ...capabilities.CapabilitiesFunc) (WebDriver, *exec.Cmd) {

	// Start Firefox webdriver proxy - GeckoDriver
	// Redirects gecko proxy output to stdout and stderr
	// Into projects logs directory
	cmd := exec.Command("zsh", "-c", GeckoDriverPath, "--port", "4444", ">", "logs/gecko.session.logs", "2>&1", "&")
	err := cmd.Start()
	if err != nil {
		log.Println("Failed to start driver:", err)
		return &Session{}, cmd
	}

	// Tries to get webdriver process status
	// Once driver isReady, returns command for deferred kill
	for i := 0; i < 30; i++ {
		time.Sleep(50 * time.Millisecond)
		stat, err := GetStatus()
		if err != nil {
			log.Println("Error status:", err)
			return &Session{}, cmd
		}

		if stat.Ready {
			log.Println("Driver ready:", err)
			break
		}
	}

	c := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

	st := strategy.NewRequester(&DriverRequest{
		DriverUrl: path.Url(path.Session),
	})

	r := st.Post(c)

	res := new(struct{ Value NewSessionResponse })
	err = json.Unmarshal(r, &res)
	if err != nil {
		log.Printf("Unmarshal capabilities: %+v", err)
		return &Session{}, cmd
	}

	return &Session{
		Id: res.Value.SessionId,
	}, cmd
}

// New
// Connect to the WebDriver instance running locally
func New(capsFn ...capabilities.CapabilitiesFunc) (WebDriver, error) {

	c := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

	st := strategy.NewRequester(&DriverRequest{
		DriverUrl: path.Url(path.Session),
	})

	r := st.Post(c)

	res := new(struct{ Value NewSessionResponse })
	err := json.Unmarshal(r, &res)
	if err != nil {
		log.Printf("Unmarshal capabilities: %+v", err)
		return nil, err
	}

	return &Session{
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
func (s *Session) Quit() {
	request.Do(http.MethodDelete, path.UrlArgs(path.Session, s.Id), nil)
}
