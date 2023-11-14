package session

import (
	"encoding/json"
	"github.com/mcsymiv/go-gecko/element"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// WebDriver
// https://w3c.github.io/webdriver/
type WebDriver interface {
	Open(u string) error
	GetUrl() (string, error)
	Quit()
	Init(b, v string) element.WebElement
	FindElement(b, v string) (element.WebElement, error)
	FindElements(b, v string) (element.WebElements, error)
	ExecuteScriptSync(s string, args ...interface{}) error
	PageSource() (string, error)
}

type BrowserCapabilities interface {
	ImplilcitWait(w float32)
}

// Session
// Represents WebDriver
// Holds session Id
// Driver port
type Session struct {
	Id   string
	Port string
}

// Status response
// W3C type
type Status struct {
	Message string `json:"message"`
	Ready   bool   `json:"ready"`
}

// NewSessionResponse
// W3C type
type NewSessionResponse struct {
	SessionId    string                 `json:"sessionId"`
	Capabilities map[string]interface{} `json:"-"`
}

const GeckoDriverPath = "/Users/mcs/Development/tools/geckodriver"

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
			log.Println("Error getting driver status:", err)
			log.Println("Killing cmd:", cmd)
			cmd.Process.Kill()
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

	data, err := json.Marshal(c)
	if err != nil {
		log.Printf("New driver marshall error: %+v", err)
	}
	url := path.Url(path.Session)
	rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("New driver error request: %+v", err)
	}

	res := new(struct{ Value NewSessionResponse })
	err = json.Unmarshal(rr, &res)
	if err != nil {
		log.Printf("Unmarshal capabilities: %+v", err)
		return &Session{}, cmd
	}

	return &Session{
		Id: res.Value.SessionId,
	}, cmd
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

func (s *Session) Quit() {
	request.Do(http.MethodDelete, path.UrlArgs(path.Session, s.Id), nil)
}
