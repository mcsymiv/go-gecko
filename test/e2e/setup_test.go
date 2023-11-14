package e2e

import (
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"
)

func TestHomePage(t *testing.T) {
	d, tearDown := hooks.StartDriver()
	defer tearDown()

	h := NewHomePage(d)
	ab := h.ClickOnAbTestingLink()

	attr, _ := ab.title.Attribute("href")
	if attr != "" {
		t.Errorf("found attr")
	}

	text, _ := ab.title.Text()
	if text == "" {
		t.Errorf("unable to get text: %+v", text)
	}
}
