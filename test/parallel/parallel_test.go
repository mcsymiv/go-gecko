package parallel

import (
	"testing"

	"github.com/mcsymiv/go-gecko/driver"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/step"
)

// TestParallel
// TODO: add pool of service and driver for parallel execution
func TestParallel(t *testing.T) {

	s, d := step.StartDriver(t)
	defer s.Process.Kill()
	defer d.Quit()

	st := step.New(d)

	t.Run("search golang", func(t *testing.T) {
		d.Open("https://www.google.com")
		el := st.FindAndClick(element.ByCssSelector, "#APjFqb", t)
		el.SendKeys("golang")
		el.SendKeys(driver.EnterKey)

		el, _ = d.FindElement(element.ByCssSelector, "#rso h3")
		title := el.Text()
		if title != "The Go Programming Language" {
			t.Fatalf("Invalid search title")
		}
	})

	t.Run("open checkboxes", func(t *testing.T) {
		d.Open("https://the-internet.herokuapp.com/")
		el, _ := d.FindElement(element.ByLinkText, "Checkboxes")
		el.Click()
	})

}
