package element

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

func (e *Element) Click() error {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Click)
	data, err := json.Marshal(&Empty{})
	if err != nil {
		log.Printf("Error on empty click marshal: %+v", err)
	}
	rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		log.Printf("Error on click: %+v", err)
		return err
	}

	res := new(struct{ Value map[string]string })
	err = json.Unmarshal(rr, res)
	if res.Value["error"] != "" || err != nil {
		return err
	}

	return nil
}
