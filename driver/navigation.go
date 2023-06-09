package driver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

// Open
// Goes to url
func (d *Driver) Open(u string) (string, error) {
	url := map[string]string{
		"url": u,
	}
	param, err := json.Marshal(url)
	if err != nil {
		log.Println("Url marshal error", err)
		return "", err
	}

	// POST /url method returns null as value
	_, err = request.Do(http.MethodPost, path.UrlArgs(path.Session, d.Id, path.UrlPath), param)
	if err != nil {
		log.Println("Open url POST Error", err)
		return "", err
	}

	r, err := request.Do(http.MethodGet, path.UrlArgs(path.Session, d.Id, path.UrlPath), nil)
	if err != nil {
		log.Printf("Open url GET error: %+v", err)
		return "", err
	}

	ur := new(struct{ Value string })
	err = json.Unmarshal(r, ur)
	if err != nil {
		log.Printf("Url GET error: %+v", err)
		return "", err
	}

	return ur.Value, nil
}

// GetUrl
func (d *Driver) GetUrl() (string, error) {

	r, err := request.Do(http.MethodGet, path.UrlArgs(path.Session, d.Id, path.UrlPath), nil)
	if err != nil {
		log.Printf("Open url GET error: %+v", err)
		return "", err
	}

	ur := new(struct{ Value string })
	err = json.Unmarshal(r, ur)
	if err != nil {
		log.Printf("Url GET error: %+v", err)
		return "", err
	}

	return ur.Value, nil
}
