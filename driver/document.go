package driver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

func (d *Driver) PageSource() (string, error) {

	rr, err := request.Do(http.MethodGet, path.UrlArgs(path.Session, d.Id, path.PageSource), nil)
	if err != nil {
		log.Println("Status request error", err)
		return "", err
	}

	page := new(struct{ Value string })
	if err := json.Unmarshal(rr, page); err != nil {
		log.Println("Status unmarshal error", err)
		return "", err
	}

	return page.Value, err
}
