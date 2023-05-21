package existing

import (
	"log"
	"os"
	"testing"

	"github.com/mcsymiv/go-gecko/service"
)

func TestMain(m *testing.M) {

	// BeforeTest
	// Starts new gecko proxy process
	cmd, err := service.StartExisting()
	if err != nil {
		log.Fatal("start gecko", err)
	}

	// Creates new driver session
	// Starts TestDriver test
	t := m.Run()

	// AfterTest
	// Kills gecko proxy process
	cmd.Process.Kill()

	os.Exit(t)
}
