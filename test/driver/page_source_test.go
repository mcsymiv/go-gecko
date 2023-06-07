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

	_, err = d.Open("https://the-internet.herokuapp.com/")
	if err != nil {
		t.Errorf("Failed open url: %+v", err)
	}

	ps, err := d.PageSource()
	if err != nil || ps == "" {
		t.Errorf("Failed to get page source: %+v", err)
	}
}
