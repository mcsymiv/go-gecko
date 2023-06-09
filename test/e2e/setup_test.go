package e2e

import (
	"os"
	"testing"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/service"
)

var d driver.WebDriver
var home HomePage

func TestMain(m *testing.M) {
	cmd, _ := service.Start()
	d, _ = driver.New(capabilities.ImplicitWait(300))

	t := m.Run()

	d.Quit()
	cmd.Process.Kill()

	os.Exit(t)

}

func TestHomePage(t *testing.T) {

	h := NewHomePage(d)
	ab := h.ClickOnAbTestingLink()

	attr := ab.title.Attribute("href")
	if attr != "" {
		t.Errorf("found attr")
	}

	text := ab.title.Text()
	if text == "" {
		t.Errorf("unable to get text: %+v", text)
	}
}
