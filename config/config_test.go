package config_test

import (
	"os"
	"testing"

	"github.com/jichenssg/ftbbackup/config"
)

func TestRead(t *testing.T) {
	if err := os.Chdir("../"); err != nil {
		t.Fatalf("Failed to change working directory: %v", err)
	}
	config.GetConfig()
}
