package service

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/mcsymiv/go-gecko/capabilities"
)

var GeckoDriverPath string = "/Users/mcs/Documents/tools/geckodriver"
var ChromeDriverPath string = "/Users/mcs/Documents/tools/chromedriver"

// NewService
// Use capabilities to specify driver executable args
func NewService(caps *capabilities.Capabilities) (*exec.Cmd, error) {
	// Returns command arguments for specified driver to start from shell
	var cmdArgs []string = driverCommand(caps)

	// Start local webdriver
	// Previously used line to start driver
	// cmd := exec.Command("zsh", "-c", GeckoDriverrequest, "--port", "4444", ">", "logs/gecko.session.logs", "2>&1", "&")
	cmd := exec.Command("/bin/zsh", cmdArgs...)
	err := cmd.Start()
	if err != nil {
		log.Println("Failed to start driver service:", err)
		return nil, err
	}

	return cmd, nil
}

// driverCommand
// Check for specified driver/browser name to pass to cmd to start the driver server
func driverCommand(cap *capabilities.Capabilities) []string {
	var cmdArgs []string = []string{
		"-c",
	}

	if cap.Capabilities.AlwaysMatch.BrowserName == "firefox" {
		cmdArgs = append(cmdArgs, GeckoDriverPath, "--port", cap.Port)
	} else {
		cmdArgs = append(cmdArgs, ChromeDriverPath, fmt.Sprintf("--port=%s", cap.Port))
	}

	cmdArgs = append(cmdArgs, ">", "logs/session.log", "2>&1", "&")
	return cmdArgs
}
