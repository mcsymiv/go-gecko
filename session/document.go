package session

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

func (s *Session) PageSource() (string, error) {
	url := path.UrlArgs(path.Session, s.Id, path.PageSource)

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

func (s *Session) ExecuteScriptSync(script string, args ...interface{}) error {
	if args == nil {
		args = make([]interface{}, 0)
	}

	data, err := json.Marshal(map[string]interface{}{
		"script": script,
		"args":   args,
	})
	if err != nil {
		log.Println("Marshal execute script error", err)
		return err
	}

	url := path.UrlArgs(path.Session, s.Id, path.Execute, path.ScriptSync)
	res, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Println("Exec script request error", err)
		return err
	}

	rr := new(struct{ Value interface{} })
	err = json.Unmarshal(res, rr)
	if err != nil {
		log.Println("Exec script unmarshal error", err, rr.Value)
		return err
	}

	return nil
}
