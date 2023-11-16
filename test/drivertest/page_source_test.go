package drivertest

import (
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"
)

func TestPageSource(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	ps, err := d.PageSource()
	if ps == "" || err != nil {
		t.Errorf("Failed to get page source: %+v", err)
	}
}
