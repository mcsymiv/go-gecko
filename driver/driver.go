package driver

import (
	"encoding/json"
	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/request"
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

	Service() *exec.Cmd
}

type Driver struct {
	Session      *service.Session
	ServiceCmd   *exec.Cmd
	Capabilities *capabilities.Capabilities
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

	return &Driver{
		Session:      s,
		ServiceCmd:   cmd,
		Capabilities: &caps,
	}
}

func (d *Driver) Service() *exec.Cmd {
	return d.ServiceCmd
}

func (d *Driver) Quit() {
	url := FormatActiveSessionUrl(d)
	res, err := MakeRequest(Url(url), Method(http.MethodDelete))
	if err != nil {
		log.Printf("Error quit request: %+v", err)
	}
	log.Printf("Quit response: %+v", string(res))
	//request.Do(http.MethodDelete, request.UrlArgs(request.Session, d.Session.SessionId), nil)
}

// Open
// Goes to url
func (d *Driver) Open(u string) error {
	url := FormatActiveSessionUrl(d, "url")

	data, _ := json.Marshal(map[string]string{
		"url": u,
	})

	response, err := MakeRequest(Url(url), Method(http.MethodPost), Payload(data))
	if err != nil {
		log.Printf("Error make request: %+v", err)
		return err
	}

	resData := new(struct{ Value []map[string]string })
	if err := json.Unmarshal(response, &resData); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
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
	url := request.UrlArgs(request.Session, d.Session.SessionId, request.UrlPath)
	rr, err := request.Do(http.MethodGet, url, nil)
	if err != nil {
		log.Printf("Open request error: %+v", err)
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
	url := request.UrlArgs(request.Session, d.Session.SessionId, request.SwitchFrame)
	param := map[string]int{
		"id": 0,
	}
	data, err := json.Marshal(param)
	if err != nil {
		log.Println("Switch frame marshal error", err)
		return err
	}

	rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Println("Switch frame request error", err)
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
	url := request.UrlArgs(request.Session, d.Session.SessionId, request.PageSource)

	rr, err := request.Do(http.MethodGet, url, nil)
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

	data, err := json.Marshal(map[string]interface{}{
		"script": script,
		"args":   args,
	})
	if err != nil {
		log.Println("Marshal execute script error", err)
		return nil, err
	}

	url := request.UrlArgs(request.Session, d.Session.SessionId, request.Execute, request.ScriptSync)
	res, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Println("Exec script request error", err)
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
