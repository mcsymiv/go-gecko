package drivertest

import (
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/hooks"
	"testing"
)

func TestExecuteScriptSync(t *testing.T) {

	d, tear := hooks.Driver()
	defer tear()

	d.Open("https://the-internet.herokuapp.com/")

	_, err := d.ExecuteScriptSync("document.querySelector('#content li a').click()")
	if err != nil {
		t.Errorf("Failed script: %+v", err)
	}

	el, err := d.FindElement(driver.ByLinkText, "Elemental Selenium")
	if err != nil {
		t.Errorf("Failed find el: %+v:", err)
	}

	_, err = d.ExecuteScriptSync("arguments[0].click()", el.ElementIdentifier())
	if err != nil {
		t.Errorf("Failed to pass el as arg: %+v", err)
	}
}
