package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/Genez-io/pulumi-genezio/provider/domain"
)

func RunScriptsInDirectory(dir string, scripts []string, envVars *[]domain.EnvironmentVariable) error {

	for _, script := range scripts {
		parts := strings.Fields(script)
		cmdName := parts[0]
		cmdArgs := parts[1:]

		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("cmd", append([]string{"/C", cmdName}, cmdArgs...)...)
		} else {
			cmd = exec.Command(cmdName, cmdArgs...)
		}

		cmd.Dir = dir

		if envVars != nil {
			for _, envVar := range *envVars {
				cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", envVar.Name, envVar.Value))
			}
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error running script: %v\nOutput: %s", err, output)
		}
	}

	return nil
}
