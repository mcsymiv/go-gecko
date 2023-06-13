package session

import (
	"encoding/json"
	"log"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/strategy"
)

type DocumentRequest struct {
	DocumentUrl string
}

func (d *DocumentRequest) Url() string {
	return d.DocumentUrl
}

func (s *Session) PageSource() string {

	st := strategy.NewRequester(&DocumentRequest{
		DocumentUrl: path.UrlArgs(path.Session, s.Id, path.PageSource),
	})

	return st.Get()
}

// ExecuteScriptSync
func (s *Session) ExecuteScriptSync(script string, args ...interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	st := strategy.NewRequester(&DocumentRequest{
		DocumentUrl: path.UrlArgs(path.Session, s.Id, path.Execute, path.ScriptSync),
	})

	r := st.Post(map[string]interface{}{
		"script": script,
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
