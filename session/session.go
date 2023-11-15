package session

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os/exec"
	// "strings"
	"time"

	"github.com/mcsymiv/go-gecko/element"

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
	ExecuteScriptSync(s string, args ...interface{}) (interface{}, error)
	PageSource() (string, error)
	IsPageLoaded()
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

var GeckoDriverPath string = "/Users/mcs/Development/tools/geckodriver"
var ChromeDriverPath string = "/Users/mcs/Development/tools/chromedriver"

func NewDriver(capsFn ...capabilities.CapabilitiesFunc) (WebDriver, *exec.Cmd) {

  var driverPath string
	c := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

  if c.Capabilities.AlwaysMatch.BrowserName == "firefox" {
    driverPath = GeckoDriverPath
  } else {
    driverPath = ChromeDriverPath 
  }

	var cmdArgs []string = []string{
		"-c",    
		driverPath,
    // "--port",
    // "4444",
		">",
		"logs/session.log",
		"2>&1",
		"&",
	}
	// Start Firefox webdriver proxy - GeckoDriver
	// Redirects gecko proxy output to stdout and stderr
	// Into projects logs directory

  // Previously used line to start driver
	// cmd := exec.Command("zsh", "-c", GeckoDriverPath, "--port", "4444", ">", "logs/gecko.session.logs", "2>&1", "&")
  cmd := exec.Command("/bin/zsh", cmdArgs...)
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
			return nil, cmd
		}

		if stat.Ready {
			log.Println("Driver ready:", err)
			break
		}
	}


  res := initDriver(&c)
  if res == nil {
    log.Fatal("Unable to get capabilities", res)
  }

	return &Session{
		Id: res.SessionId,
	}, cmd
}

// initDriver
// Return NewSessionResponce with session Id
func initDriver(c *capabilities.NewSessionCapabilities) *NewSessionResponse {
	data, err := json.Marshal(c)
	if err != nil {
		log.Printf("New driver marshall error: %+v", err)
    return nil 
	}
	url := path.Url(path.Session)
	rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("New driver error request: %+v", err)
    return nil 
	}

	res := new(struct{ Value NewSessionResponse })
	err = json.Unmarshal(rr, &res)
	if err != nil {
		log.Printf("Unmarshal capabilities: %+v", err)
    return nil
	}

  return &res.Value
}

// GetStatus
//
// Status returns information about whether a remote end is
// in a state in which it can create new sessions,
// but may additionally include arbitrary meta information
// that is specific to the implementation.
func GetStatus() (*Status, error) {
	url := path.Url(path.Status)
	rr, err := request.Do(http.MethodGet, url, nil)
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
