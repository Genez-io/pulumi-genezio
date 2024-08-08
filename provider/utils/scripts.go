package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func RunScriptInDirectory(dir string, script string, envVars *[]domain.EnvironmentVariable) error {

	if envVars != nil {
		SetProccessEnvironmentVariables(*envVars)
	}

	// Save the current working directory
	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	// Change the working directory to the specified directory
	if err := os.Chdir(dir); err != nil {
		return fmt.Errorf("error changing to directory %s: %v", dir, err)
	}

	// Determine the appropriate shell command based on the current OS
	var shellCmd, shellArg string
	switch runtime.GOOS {
	case "windows":
		shellCmd = "cmd"
		shellArg = "/C"
	default: // Assume Unix-like
		shellCmd = "/bin/sh"
		shellArg = "-c"
	}

	// Create a new command with the script as the argument
	cmd := exec.Command(shellCmd, shellArg, script)

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Restore the original working directory
		if err := os.Chdir(originalDir); err != nil {
			return fmt.Errorf("error restoring working directory: %v", err)
		}
		return fmt.Errorf("error running script: %v\nOutput: %s", err, output)
	}

	// Restore the original working directory
	if err := os.Chdir(originalDir); err != nil {
		return fmt.Errorf("error restoring working directory: %v", err)
	}

	return nil
}

func SetProccessEnvironmentVariables(envVars []domain.EnvironmentVariable) error {
	for _, envVar := range envVars {
		if err := os.Setenv(envVar.Name, envVar.Value); err != nil {
			return fmt.Errorf("error setting environment variable %s: %v", envVar.Name, err)
		}
	}
	return nil
}
