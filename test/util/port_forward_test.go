package util

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/hooks"
	"github.com/mcsymiv/go-gecko/step"
)

func findFile(r, n string) (string, error) {
	var f string

	err := filepath.WalkDir(r, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println("err")
			return err
		}
		if !info.IsDir() && info.Name() == n {
			f = path
		}
		return nil
	})

	if err != nil {
		return f, err
	}

	return f, nil
}

// ExecReplace
func dotenv(filepath string) {
	// read file into memory
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer f.Close()

	// var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		env := scanner.Text()
		key := strings.Split(env, "=")[0]
		value := strings.Split(env, "=")[1]
		os.Setenv(key, value)
	}
}

func loadEnv() {
	f, err := findFile("../../config", ".env")
	log.Println(f)
	if err != nil {
		log.Fatal("file not found", err)
	}
	dotenv(f)
}

func TestPortForward(t *testing.T) {
	loadEnv()

	d, tear := hooks.StartDriver()
	defer tear()

	d.Open("http://192.168.0.1")
	d.IsPageLoaded()
	st := step.New(d)

	time.Sleep(2 * time.Second)
	st.FindCss("[id='local-pwd-tb'] [type='password']").SendAndSubmit(os.Getenv("PORT_PASS"))
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

func TestDownload(t *testing.T) {
	loadEnv()

	d, tear := hooks.Driver(
		capabilities.ImplicitWait(10000),
		capabilities.BrowserName("firefox"),
	)
	defer tear()

	d.Open(os.Getenv("DOWNLOAD_URL"))
	d.IsPageLoaded()

	st := step.New(d)
	st.FindX(".//a[text()='Log in using Azure Active Directory']").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindCss("[id='i0116']").SendAndSubmit(os.Getenv("DOWNLOAD_LOGIN"))
	time.Sleep(2 * time.Second)
	st.FindCss("[id='i0118']").SendAndSubmit(os.Getenv("DOWNLOAD_PASS"))
	time.Sleep(2 * time.Second)
	st.FindCss("[id='idSIButton9']").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindX(".//span[text()='Projects']").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindCss("[id='search-projects']").Element().SendKeys("dev01")
	time.Sleep(2 * time.Second)
	st.FindX(".//aside//span[contains(text(),'Smoke (Concurrent tests')]").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindCss("[data-grid-root] [data-test-build-number-link]").Element().Click()
	time.Sleep(2 * time.Second)
	st.FindCss("[data-tab-title='Allure Report']").Element().Click()
	time.Sleep(2 * time.Second)
}
