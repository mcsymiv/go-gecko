package element

import (
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

func (e *Element) Click() {
	url := path.UrlArgs(path.Session, e.SessionId, path.Element, e.Id, path.Click)
	_, err := request.Do(http.MethodPost, url, nil)
	if err != nil {
		fmt.Println("Error on click", err)
	}
}
