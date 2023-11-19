package drivertest

import (
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/hooks"
	"log"
	"testing"
)

func TestElementAttribute(t *testing.T) {
	d, tear := hooks.Driver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	_, err := d.FindElement(driver.ByLinkText, "Typos")
	if err != nil {
		log.Fatal("find element", err)
	}

	//attr, _ := el.Attribute("href")
	//if attr == "" {
	//	t.Errorf("element attribute: %+v", attr)
	//}

}
