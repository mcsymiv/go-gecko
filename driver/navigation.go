package driver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// Open
// Goes to url
func (d Driver) Open(u string) error {
	data, err := json.Marshal(map[string]string{
		"url": u,
	})
	if err != nil {
		log.Printf("Open marshal: %+v", err)
		return err
	}

	url := path.UrlArgs(path.Session, d.SessionId, path.UrlPath)
	log.Println("url", url)
	rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Open request error: %+v", err)
		return err
	}

	res := new(struct{ Value []map[string]string })
	if err := json.Unmarshal(rr, &res); err != nil {
		log.Printf("Find element unmarshal: %+v", err)
		return err
	}

	return nil
}

// IsPageLoaded
// TODO
// Should validate if page is fully loaded
// And block test execution until true
func (d Driver) IsPageLoaded() {
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

// GetUrl
func (d Driver) GetUrl() (string, error) {
	url := path.UrlArgs(path.Session, d.SessionId, path.UrlPath)
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

func (d Driver) SwitchFrame(e element.WebElement) error {
	url := path.UrlArgs(path.Session, d.SessionId, path.SwitchFrame)
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
