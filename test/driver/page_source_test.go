package driver

import (
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
)

func TestPageSource(t *testing.T) {

	caps := capabilities.ImplicitWait(3000)
	d, err := driver.New(caps)
	if err != nil {
		t.Errorf("Failed session: %+v", err)
	}

	defer d.Quit()

	d.Open("https://the-internet.herokuapp.com/")

	ps := d.PageSource()
	if ps == "" {
		t.Errorf("Failed to get page source: %+v", err)
	}
}
