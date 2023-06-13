package driver

import (
	"log"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/session"
)

func TestDriver(t *testing.T) {

	d, err := session.New(capabilities.ImplicitWait(3000))
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()
}
