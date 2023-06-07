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

// ExecuteScriptSync
func (d *Driver) ExecuteScriptSync(s string, args []interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	data, err := json.Marshal(map[string]interface{}{
		"script": s,
		"args":   args,
	})

	r, err := request.Do(http.MethodPost, path.UrlArgs(path.Session, d.Id, path.Execute, path.ScriptSync), data)
	if err != nil {
		log.Println("Status request error", err)
		return err
	}

	page := new(struct{ Value interface{} })
	err = json.Unmarshal(r, page)
	if page.Value != nil || err != nil {
		log.Println("Status unmarshal error", err, page.Value)
		return err
	}

	return nil
}
