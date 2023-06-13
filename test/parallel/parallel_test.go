package parallel

import (
	"testing"

	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/session"
	"github.com/mcsymiv/go-gecko/step"
)

// TestParallel
// TODO: add pool of service and driver for parallel execution
func TestParallel(t *testing.T) {

	d, s := step.StartDriver(t)
	defer d.Process.Kill()
	defer s.Quit()

	st := step.New(s)

	t.Run("search golang", func(t *testing.T) {
		s.Open("https://www.google.com")
		el := st.FindAndClick(element.ByCssSelector, "#APjFqb", t)
		el.SendKeys("golang")
		el.SendKeys(session.EnterKey)

		el, _ = s.FindElement(element.ByCssSelector, "#rso h3")
		title := el.Text()
		if title != "The Go Programming Language" {
			t.Fatalf("Invalid search title")
		}
	})

	t.Run("open checkboxes", func(t *testing.T) {
		s.Open("https://the-internet.herokuapp.com/")
		el, _ := s.FindElement(element.ByLinkText, "Checkboxes")
		el.Click()
	})

}
