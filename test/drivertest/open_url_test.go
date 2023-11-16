package drivertest

import (
	"github.com/mcsymiv/go-gecko/hooks"
	"log"
	"testing"
)

func TestOpenUrl(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	err := d.Open("https://the-internet.herokuapp.com/")
	if err != nil {
		log.Fatal("Open error")
	}
}
