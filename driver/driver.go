package driver

import (
	"github.com/mcsymiv/go-gecko/capabilities"
	"github.com/mcsymiv/go-gecko/service"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// WebDriver
// https://w3c.github.io/webdriver/
type WebDriver interface {
	Open(u string) (WebDriverResponse, error)
	GetUrl() (WebDriverResponse, error)
	Quit()
}

type Driver struct {
	Client       *http.Client
	Session      *service.Session
	Service      *exec.Cmd
	Capabilities *capabilities.Capabilities
}

func NewDriver(capsFn ...capabilities.CapabilitiesFunc) WebDriver {
	caps := capabilities.DefaultCapabilities()
	for _, capFn := range capsFn {
		capFn(&caps)
	}

	cmd, err := service.NewService(&caps)
	if err != nil {
		log.Fatal("Unable to start driver service", err)
	}

	// Tries to get driver status for 2 seconds
	// Once driver isReady, returns command for deferred kill
	start := time.Now()
	end := start.Add(5 * time.Second)
	for stat, err := service.GetStatus(&caps); err != nil || !stat.Ready; stat, err = service.GetStatus(&caps) {
		time.Sleep(500 * time.Millisecond)
		log.Println("Error getting driver status:", err)

		if time.Now().After(end) {
			log.Println("Killing cmd:", cmd)
			cmd.Process.Kill()
			return nil
		}
	}

	s, err := service.NewSession(&caps)
	if err != nil || s == nil {
		log.Fatal("Unable to start session", s)
	}

	ro := DefaultRequestOptions()

	return &Driver{
		Client:       &http.Client{},
		Session:      s,
		ServiceCmd:   cmd,
		Capabilities: &caps,
	}
}
