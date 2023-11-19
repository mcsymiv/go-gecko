package driver

import (
	"encoding/json"
	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/service"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// WebDriver
// https://w3c.github.io/webdriver/
type WebDriver interface {
	Open(u string) error
	GetUrl() (string, error)
	Quit()
	FindElement(b, v string) (WebElement, error)
	FindElements(b, v string) (WebElements, error)
	ExecuteScriptSync(s string, args ...interface{}) (interface{}, error)
	PageSource() (string, error)
	IsPageLoaded()
	SwitchFrame(WebElement) error

	// Service util function
	// To stop/kill local driver process
	Service() *exec.Cmd

	// MakeRequest
	// Performs API request on driver
	// TODO: Can be adjusted to make custom API calls if exposed correctly
	MakeRequest(options ...RequestOptionFunc) ([]byte, error)
}

type Driver struct {
	RequestOptions *RequestOptions
	Session        *service.Session
	ServiceCmd     *exec.Cmd
	Capabilities   *capabilities.Capabilities
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

func (d *Driver) Quit() {
	url := formatActiveSessionUrl(d)
	_, err := makeReq(d, WithMethod(http.MethodDelete), WithUrl(url))
	if err != nil {
		log.Printf("Error quit request: %+v", err)
	}
}

// Open
// Goes to url
func (d *Driver) Open(u string) error {
	url := formatActiveSessionUrl(d, "url")
	data, _ := json.Marshal(map[string]string{
		"url": u,
	})

	_, err := makeReq(d, WithMethod(http.MethodPost), WithUrl(url), WithPayload(data))
	if err != nil {
		log.Printf("Error make request: %+v", err)
		return err
	}

	return nil
}

// IsPageLoaded
// TODO
// Should validate if page is fully loaded
// And block test execution until true
func (d *Driver) IsPageLoaded() {
	load := `
    function load() {
      if (document.readyState === "complete") {
        return true
      } else if (document.readyState === "interactive") {
        // DOM ready! Images, frames, and other subresources are still downloading.
        return false
      } else {
        return false
      }
    }
    return load()
  `
	res, err := d.ExecuteScriptSync(load)
	if err != nil {
		log.Println("Page load script error", err)
	}

	if res.(bool) {
		return
	}

	for {
		if res, _ = d.ExecuteScriptSync(load); res != nil && res.(bool) {
			break
		}
	}
}

func (d *Driver) GetUrl() (string, error) {
	url := formatActiveSessionUrl(d, "url")
	rr, err := makeReq(d, WithMethod(http.MethodGet), WithUrl(url))
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

func (d *Driver) SwitchFrame(e WebElement) error {
	url := formatActiveSessionUrl(d, "frame")
	param := map[string]int{
		"id": 0,
	}
	data, err := json.Marshal(param)
	if err != nil {
		log.Println("Switch frame marshal error", err)
		return err
	}

	rr, err := makeReq(d, WithMethod(http.MethodPost), WithUrl(url), WithPayload(data))
	if err != nil {
		log.Printf("Error make request: %+v", err)
		return err
	}

	val := new(struct{ Value map[string]interface{} })
	err = json.Unmarshal(rr, val)
	if err != nil {
		log.Printf("Switch frame error on unmarshal: %+v", err)
		return nil
	}

	return nil
}

func (d *Driver) PageSource() (string, error) {
	url := formatActiveSessionUrl(d, "source")
	rr, err := makeReq(d, WithMethod(http.MethodGet), WithUrl(url))
	if err != nil {
		log.Println("Page source request error", err)
		return "", err
	}

	reply := new(struct{ Value string })
	if err := json.Unmarshal(rr, reply); err != nil {
		log.Println("Page source unmarshal error", err)
		return reply.Value, err
	}

	return reply.Value, nil
}

func (d *Driver) ExecuteScriptSync(script string, args ...interface{}) (interface{}, error) {
	if args == nil {
		args = make([]interface{}, 0)
	}

	url := formatActiveSessionUrl(d, "execute", "script")
	data, err := json.Marshal(map[string]interface{}{
		"script": script,
		"args":   args,
	})
	if err != nil {
		log.Println("Marshal execute script error", err)
		return nil, err
	}

	res, err := makeReq(d, WithMethod(http.MethodPost), WithUrl(url), WithPayload(data))
	if err != nil {
		log.Printf("Error make request: %+v", err)
		return nil, err
	}

	rr := new(struct{ Value interface{} })
	err = json.Unmarshal(res, rr)
	if err != nil {
		log.Println("Exec script unmarshal error", err, rr.Value)
		return nil, err
	}

	return rr.Value, nil
}

// MakeRequest
// Wrapper function exposed on WebDriver to make external API calls
// Uses private client.makeReq implementation
func (d *Driver) MakeRequest(options ...RequestOptionFunc) ([]byte, error) {
	return makeReq(d, options...)
}
