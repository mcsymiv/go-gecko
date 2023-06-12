package driver

import (
	"encoding/json"
	"log"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
)

type Document struct {
	DocumentUrl string
}

func (d *Document) Url() string {
	return d.DocumentUrl
}

func (d *Driver) PageSource() string {

	st := strategy.NewRequester(&Document{
		DocumentUrl: path.UrlArgs(path.Session, d.Id, path.PageSource),
	})

	return st.Get()
}

// ExecuteScriptSync
func (d *Driver) ExecuteScriptSync(s string, args ...interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	st := strategy.NewRequester(&Document{
		DocumentUrl: path.UrlArgs(path.Session, d.Id, path.Execute, path.ScriptSync),
	})

	r := st.Post(map[string]interface{}{
		"script": s,
		"args":   args,
	})

	sr := new(struct{ Value interface{} })
	err := json.Unmarshal(r, sr)
	if sr.Value != nil || err != nil {
		log.Println("Status unmarshal error", err, sr.Value)
		return err
	}

	return nil
}
