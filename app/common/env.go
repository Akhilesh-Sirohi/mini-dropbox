package common

import (
	"fmt"
	"os"
)

var availableEnvs map[string]bool

func init() {
	availableEnvs = map[string]bool{
		"dev_docker": true,
		"dev":        true,
		"stage":      true,
		"devstack":   true,
		"prod":       true,
		"drone":      true,
		"func":       true,
		"perf":       true,
		"automation": true,
		"default":    true,
		"gamma":      true,
		"beta":       true,
		"uat":        true,
		"dark":       true,
	}
}

// GetEnv returns the environment string
func GetEnv() (string, error) {
	// Fetch env for bootstrapping
	environ := os.Getenv("APP_MODE")
	if environ == "" {
		environ = "dev"
	}
	if _, ok := availableEnvs[environ]; !ok {
		return "", fmt.Errorf("not a valid environment value : %s", environ)
	}
	return environ, nil
}
