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
  SwitchFrame(element.WebElement) error
  SwitchFrameParent() error
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

var GeckoDriverPath string = "/Users/mcs/Documents/tools/geckodriver"
var ChromeDriverPath string = "/Users/mcs/Documents/tools/chromedriver"

func NewDriver(capsFn ...capabilities.CapabilitiesFunc) (WebDriver, *exec.Cmd) {

	c := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}

	// Returns command arguments for specified driver to start from shell
	var cmdArgs []string = driverCommand(c)
	log.Println(cmdArgs)

 // todo
 // create http client with set port

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

	// Tries to get driver status for 2 seconds
	// Once driver isReady, returns command for deferred kill
	start := time.Now()
	end := start.Add(2 * time.Second)
	for stat, err := GetStatus(); err != nil || !stat.Ready; stat, err = GetStatus() {
		time.Sleep(200 * time.Millisecond)
		log.Println("Error getting driver status:", err)

		if time.Now().After(end) {
			log.Println("Killing cmd:", cmd)
			cmd.Process.Kill()
			return nil, cmd
		}
	}

	s := startSession(&c)
	if s == nil {
		log.Fatal("Unable to start session", s)
	}

	return &Session{
		Id: s.SessionId,
	}, cmd
}

// initDriver
// Return NewSessionResponce with session Id
func startSession(c *capabilities.NewSessionCapabilities) *NewSessionResponse {
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

// driverCommand
// Check for specified driver/browser name to pass to cmd to start the driver server
func driverCommand(cap capabilities.NewSessionCapabilities) []string {
	var cmdArgs []string = []string{
		"-c",
	}

	if cap.Capabilities.AlwaysMatch.BrowserName == "firefox" {
		cmdArgs = append(cmdArgs, GeckoDriverPath, "--port", "4444")
	} else {
		cmdArgs = append(cmdArgs, ChromeDriverPath)
	}

	cmdArgs = append(cmdArgs, ">", "logs/session.log", "2>&1", "&")
	return cmdArgs
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
