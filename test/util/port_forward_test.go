package util

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/hooks"
	"github.com/mcsymiv/go-gecko/step"
)

func TestPortForward(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("http://192.168.0.1")
	d.IsPageLoaded()
	st := step.New(d)

	time.Sleep(2 * time.Second)
	st.FindCss("[id='local-login-button'] [title='LOG IN']").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindX(".//span[contains(text(),'Advanced')]").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindX(".//span[text()='NAT Forwarding']").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindX(".//span[text()='Port Forwarding']").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindX(".//a[contains(@class,'btn-edit')]").Element().Click()
	time.Sleep(2 * time.Second)
  st.FindX(".//span[text()='VIEW CONNECTED DEVICES']").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindX(".//span[contains(text(), 'mcs-pc')]/..").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindX(".//div[contains(@class,'msg-content-wrap')]//span[text()='SAVE']").Element().Click()
}

