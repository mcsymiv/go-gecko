package util

import (
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/element"
	"github.com/mcsymiv/go-gecko/hooks"
	"github.com/mcsymiv/go-gecko/step"
)

func TestPortForward(t *testing.T) {

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("http://192.168.0.1")
	d.IsPageLoaded()

	time.Sleep(2 * time.Second)
	passInput, _ := d.FindElement(element.ByCssSelector, "[id='local-pwd-tb'] [type='password']")
	passInput.SendKeys("OrderOfThePhoniex1")

	time.Sleep(2 * time.Second)
	inputBtn, _ := d.FindElement(element.ByCssSelector, "[id='local-login-button'] [title='LOG IN']")
	inputBtn.Click()

	time.Sleep(2 * time.Second)
	advanced, _ := d.FindElement(element.ByXPath, ".//span[contains(text(),'Advanced')]")
	advanced.Click()

	time.Sleep(2 * time.Second)
	natForward, _ := d.FindElement(element.ByXPath, ".//span[text()='NAT Forwarding']")
	natForward.Click()

	time.Sleep(2 * time.Second)
	portForward, _ := d.FindElement(element.ByXPath, ".//span[text()='Port Forwarding']")
	portForward.Click()

	time.Sleep(2 * time.Second)
	editIcon, _ := d.FindElement(element.ByXPath, ".//a[contains(@class,'btn-edit')]")
	editIcon.Click()

	time.Sleep(2 * time.Second)
	connDevices, _ := d.FindElement(element.ByXPath, ".//span[text()='VIEW CONNECTED DEVICES']")
	connDevices.Click()

	time.Sleep(2 * time.Second)
	myPcDevice, _ := d.FindElement(element.ByXPath, ".//span[contains(text(), 'mcs-pc')]/..")
	myPcDevice.Click()

	time.Sleep(2 * time.Second)
	saveBtn, _ := d.FindElement(element.ByXPath, ".//div[contains(@class,'msg-content-wrap')]//span[text()='SAVE']")
	saveBtn.Click()
}

func TestDownload(t *testing.T) {
	d, tear := hooks.Driver(
    capabilities.ImplicitWait(3000),
    capabilities.BrowserName("chrome"),
  )
	defer tear()

	d.Open("https://bgdt-teamcity.elateral-dev.io/login.html")
	d.IsPageLoaded()

	st := step.New(d)
	st.FindX(".//a[text()='Log in using Azure Active Directory']").Element().Click()

	st.FindCss("[id='i0116']").SendAndSubmit("serhii.maksymiv@elateralazure.onmicrosoft.com")

	time.Sleep(3 * time.Second)
	st.FindCss("[id='i0118']").SendAndSubmit("!Mcsymivqamadness29$")

	time.Sleep(2 * time.Second)
	st.FindCss("[id='idSIButton9']").Element().Click()

	time.Sleep(2 * time.Second)
	st.FindX(".//span[text()='Projects']").Element().Click()
	st.FindCss("[id='search-projects']").SendAndSubmit("dev01")

	st.FindX(".//aside//span[text()='1 - Smoke (Concurrent tests)']").Element().Click()

	time.Sleep(2 * time.Second)
	st.FindCss("[data-grid-root] [data-test-build-number-link]").Element().Click()

	time.Sleep(2 * time.Second)
	st.FindCss("[data-tab-title='Allure Report']").Element().Click()

	time.Sleep(2 * time.Second)
}
