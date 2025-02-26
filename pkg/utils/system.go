package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func FindPIDByPort(port string) (string, error) {
	cmd := exec.Command("lsof", "-t", "-i", ":"+port)
	output, err := cmd.Output()
	if err != nil {
		if err.Error() == "exit status 1" {
			return "", fmt.Errorf("No process found on port: %s", port)
		}
		return "", fmt.Errorf("error getting process info: %v", err)
	}
	if len(output) == 0 {
		return "", nil
	}
	return strings.TrimSpace(string(output)), nil
}

func KillProcess(pid string) error {
	killCmd := exec.Command("kill", pid)
	if err := killCmd.Run(); err != nil {
		return err
	}
	return nil
}
