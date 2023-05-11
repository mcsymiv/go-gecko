package element

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/request"
)

const (
	ByID              = "id"
	ByXPATH           = "xpath"
	ByLinkText        = "link text"
	ByPartialLinkText = "partial link text"
	ByName            = "name"
	ByTagName         = "tag name"
	ByClassName       = "class name"
	ByCSSSelector     = "css selector"
)

func Find(sid string) {
	params := map[string]string{
		"using": ByCSSSelector,
		"value": "#APjFqb",
	}

	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println("Find element error marshal", err)
	}

	url := fmt.Sprintf("%s/session/%s/element", request.BaseUrl, sid)
	rr, err := request.Do(http.MethodPost, url, data)
	if err != nil {
		fmt.Println("Find element request error", err)
	}

	fmt.Printf("%+v", string(rr))

}
