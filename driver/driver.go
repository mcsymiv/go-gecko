package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/service"
)

type GoDriver struct {
	WebDriver
	Config
}

type Config struct {
	Timeout time.Duration
}

// WebDriver
// https://w3c.github.io/webdriver/
type WebDriver interface {
	Open(u string) error
	// GetUrl() (string, error)
	Quit()
	FindElement(b, v string) (WebElement, error)
	// FindElements(b, v string) (WebElements, error)

	// Service util function
	// To stop/kill local driver process
	Service() *exec.Cmd

	// MakeRequest
	// Performs API request on driver
	// TODO: Can be adjusted to make custom API calls if exposed correctly
	// MakeRequest(options ...service.RequestOptionFunc) ([]byte, error)
}

type Driver struct {
	Client       *service.WebClient
	Session      *service.Session
	ServiceCmd   *exec.Cmd
	Capabilities *capabilities.Capabilities
}

func NewDriver(capsFn ...capabilities.CapabilitiesFunc) WebDriver {
	caps := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&caps)
	}

	log.Printf("%+v", caps)

	cmd, err := service.NewService(&caps)
	if err != nil {
		log.Fatal("unable to start driver service", err)
	}

	// Tries to get driver status for 2 seconds
	// Once driver isReady, returns command for deferred kill
	start := time.Now()
	end := start.Add(2 * time.Second)
	for stat, err := service.GetStatus(&caps); err != nil || !stat.Ready; stat, err = service.GetStatus(&caps) {
		time.Sleep(200 * time.Millisecond)
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

	client := service.NewClient()

	return &Driver{
		Client:       client,
		Session:      s,
		ServiceCmd:   cmd,
		Capabilities: &caps,
	}
}

func (d *Driver) Service() *exec.Cmd {
	return d.ServiceCmd
}

func (d Driver) Quit() {
	url := fmt.Sprintf("http://localhost:4444/session/%s", d.Session.SessionId)
	req, _ := http.NewRequest(http.MethodDelete, url, nil)
	req.Header.Add("Accept", "json/application")
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Printf("Error quit request: %+v", err)
	}

	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error quit request: %+v", err)
	}
}

// Open
// Goes to url
func (d Driver) Open(u string) error {
	url := fmt.Sprintf("http://localhost:4444/session/%s/url", d.Session.SessionId)
	data, _ := json.Marshal(map[string]string{
		"url": u,
	})
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	req.Header.Add("Accept", "json/application")
	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Printf("Error quit request: %+v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	reply := make(map[string]string)
	if err := json.Unmarshal(body, &reply); err != nil {
		log.Println("Status unmarshal error", err)
		return err
	}

	return nil
	// url := formatActiveSessionUrl(d, "url")
	// url := fmt.Sprintf("%s/session/%s/url", d.Client.RequestOptions.Url, d.Session.SessionId)
	// data, _ := json.Marshal(map[string]string{
	// 	"url": u,
	// })

	// _, err := d.MakeRequest(
	// 	d.Client.RequestOptions.WithMethod(http.MethodPost),
	// 	d.Client.RequestOptions.WithUrl(url),
	// 	d.Client.RequestOptions.WithPayload(data),
	// )
	// if err != nil {
	// 	log.Printf("Error make request: %+v", err)
	// 	return err
	// }

	return nil
}

/*
func (d *Driver) GetUrl() (string, error) {
	url := formatActiveSessionUrl(d, "url")
	rr, err := d.MakeRequest(
		d.Client.RequestOptions.WithMethod(http.MethodGet),
		d.Client.RequestOptions.WithUrl(url),
	)
	if err != nil {
		log.Printf("Error make request: %+v", err)
		return "", err
	}

	val := new(struct{ Value string })
	err = json.Unmarshal(rr, val)
	if err != nil {
		log.Printf("Error on unmarshal response: %+v", err)
		return "", nil
	}

	return val.Value, nil
}
*/

// MakeRequest
// Wrapper function exposed on WebDriver to make external API calls
// Uses private client.makeReq implementation
/*
func (d *Driver) MakeRequest(options ...service.RequestOptionFunc) ([]byte, error) {
	return d.Client.MakeRequest(options...)
}
*/

func formatActiveSessionUrl(d *Driver, args ...string) string {
	return ""
}
