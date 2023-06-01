package driver

import (
	"log"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
)

func TestOpenUrl(t *testing.T) {

	d, err := driver.New(capabilities.ImplicitWait(3000))
	if err != nil {
		log.Fatal("session start err", err)
	}
	defer d.Quit()

	u := "https://the-internet.herokuapp.com/"
	url, err := d.Open(u)
	if err != nil {
		t.Errorf("Url: %+v", err)
	}

	if url != u {
		t.Errorf("Invalid url. Expected: %s. Actual: %s", u, url)
	}
}
