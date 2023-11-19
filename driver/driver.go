package driver

import (
	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/service"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// WebDriver
// https://w3c.github.io/webdriver/
type WebDriver interface {
	Open(u string) (WebDriverResponse, error)
	GetUrl() (WebDriverResponse, error)
	Quit()

	// FindElement
	// Finds element based upon w3c 'using' selectors
	// Not all using WebElement constants are supported
	// Pay attention to 'By...' comments
	FindElement(b, v string) (WebElementResponse, error)
	FindElements(b, v string) (WebElements, error)

	// Service util function
	// To stop/kill local driver process
	Service() *exec.Cmd

	// MakeRequest
	// Performs API request on driver
	// TODO: Can be adjusted to make custom API calls if exposed correctly
	MakeRequest(options ...RequestOptionFunc) ([]byte, error)
}

type WebDriverResponse interface {
	GetValue() interface{}
}

type WebElementResponse interface {
	//GetElementId() interface{}
}

// WebDriverStrategy interface
type WebDriverStrategy interface {
	Execute(*http.Client, *http.Request) (WebDriverResponse, error)
}

type WebElementStrategy interface {
	Execute(*http.Client, *http.Request) (WebElementResponse, error)
}

type Driver struct {
	RequestOptions *RequestOptions
	Session        *service.Session
	ServiceCmd     *exec.Cmd
	Capabilities   *capabilities.Capabilities

	WebDriverStrategy
	WebElementStrategy
}

func NewDriver(capsFn ...capabilities.CapabilitiesFunc) WebDriver {
	caps := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&caps)
	}

	cmd, err := service.NewService(&caps)
	if err != nil {
		log.Fatal("Unable to start driver service", err)
	}

	// Tries to get driver status for 2 seconds
	// Once driver isReady, returns command for deferred kill
	start := time.Now()
	end := start.Add(5 * time.Second)
	for stat, err := service.GetStatus(&caps); err != nil || !stat.Ready; stat, err = service.GetStatus(&caps) {
		time.Sleep(500 * time.Millisecond)
		log.Println("Error getting driver status:", err)

		if time.Now().After(end) {
			log.Println("Killing cmd:", cmd)
			cmd.Process.Kill()
			return nil
		}
	}

	s, err := service.NewSession(&caps)
	if err != nil || s == nil {
		log.Fatal("Unable to start session", s)
	}

	ro := DefaultRequestOptions()

	return &Driver{
		RequestOptions: &ro,
		Session:        s,
		ServiceCmd:     cmd,
		Capabilities:   &caps,
	}
}

func (d *Driver) Service() *exec.Cmd {
	return d.ServiceCmd
}

// MakeRequest
// Wrapper function exposed on WebDriver to make external API calls
// Uses private client.makeReq implementation
func (d *Driver) MakeRequest(options ...RequestOptionFunc) ([]byte, error) {
	return makeReq(d, options...)
}

func executeRequest(c *http.Client, req *http.Request) ([]byte, error) {
	response, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("Error read response body:", err)
		return nil, err
	}

	return body, nil
}
