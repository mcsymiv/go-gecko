package drivertest

import (
	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/hooks"
	"log"
	"testing"
)

func TestOpenUrl(t *testing.T) {

	d, tear := hooks.Driver(
		capabilities.ImplicitWait(3000),
	)
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")
	res, err := d.GetUrl()
	if err != nil {
		log.Println("Get url error", err)
	}

	log.Println(res)

}
