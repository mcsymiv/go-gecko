package session

import (
	"encoding/json"
	"fmt"

	"github.com/mcsymiv/go-gecko/config"
	"github.com/mcsymiv/go-gecko/models"
	"github.com/mcsymiv/go-gecko/request"
)

func New() *config.SessionConfig {

	params := &models.Capabilities{
		AlwaysMatch: models.AlwaysMatch{
			AcceptInsecureCerts: true,
		},
	}
	data, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}

	rr, err := request.Do("POST", fmt.Sprintf("%s%s", config.DriverUrl, "/session"), data)
	if err != nil {
		fmt.Println(err)
	}

	var res config.RemoteResponse

	err = json.Unmarshal(rr, &res)
	if err != nil {
		fmt.Println("error:", err)
	}

	return &config.SessionConfig{
		Id: res.Value.SessionId,
	}
}
