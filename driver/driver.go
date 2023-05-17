package driver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcsymiv/go-gecko/config"
	"github.com/mcsymiv/go-gecko/path"
	"github.com/mcsymiv/go-gecko/request"
)

func New(capsFn ...config.CapabilitiesFunc) WebDriver {

	c := config.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&c)
	}
	fmt.Printf("%+v", c)

	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}

	rr, err := request.Do(http.MethodPost, path.Url(path.Session), data)
	if err != nil {
		fmt.Println(err)
	}

	var res RemoteResponse

	err = json.Unmarshal(rr, &res)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &Driver{
		Id: res.Value.SessionId,
	}
}
