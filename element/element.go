package element

import (
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/request"
)

func (e *Element) Click() {
	url := request.UrlArgs(request.Session, e.SessionId, request.Element, e.Id, request.Click)
	_, err := request.Do(http.MethodPost, url, nil)
	if err != nil {
		fmt.Println("Error on click", err)
	}
}
