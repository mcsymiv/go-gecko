package service

import (
	"log"
	"os/exec"
	"time"

	"github.com/mcsymiv/go-gecko/session"
)

const GeckoDriverPath = "/Users/mcs/Development/tools/geckodriver"

func Start() (*exec.Cmd, error) {

	// Start Firefox webdriver proxy - GeckoDriver
	// Redirects gecko proxy output to stdout and stderr
	// Into projects logs directory
	cmd := exec.Command("zsh", "-c", GeckoDriverPath, ">", "logs/gecko.session.logs", "2>&1", "&")
	err := cmd.Start()
	if err != nil {
		log.Println("Failed to start Firefox browser:", err)
		return nil, err
	}

	// Tries to get webdriver process status
	// Once driver isReady, returns command for deferred kill
	for i := 0; i < 30; i++ {
		time.Sleep(50 * time.Millisecond)
		stat, err := session.GetStatus()

		if err == nil && stat.Ready {
			return cmd, nil
		}
	}

	return cmd, nil
}

func StartExisting() (*exec.Cmd, error) {

	// Start Firefox webdriver proxy - GeckoDriver
	// Redirects gecko proxy output to stdout and stderr
	// Into projects logs directory
	cmd := exec.Command("zsh", "-c", GeckoDriverPath, "--connect-existing", "--marionette-port", "62489", ">", "logs/gecko.session.logs", "2>&1", "&")
	err := cmd.Start()
	if err != nil {
		log.Println("Failed to start Firefox browser:", err)
		return nil, err
	}

	// Tries to get webdriver process status
	// Once driver isReady, returns command for deferred kill
	for i := 0; i < 30; i++ {
		time.Sleep(50 * time.Millisecond)
		stat, err := session.GetStatus()

		if err == nil && stat.Ready {
			return cmd, nil
		}
	}

	return cmd, nil
}
