package service

import (
	"github.com/mcsymiv/go-gecko/capabilities"
	"log"
	"os/exec"
)

var GeckoDriverPath string = "/Users/mcs/Development/tools/geckodriver"
var ChromeDriverPath string = "/Users/mcs/Development/tools/chromedriver"

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
		cmdArgs = append(cmdArgs, GeckoDriverPath, "--port", "4444")
	} else {
		cmdArgs = append(cmdArgs, ChromeDriverPath)
	}

	cmdArgs = append(cmdArgs, ">", "logs/session.log", "2>&1", "&")
	return cmdArgs
}
