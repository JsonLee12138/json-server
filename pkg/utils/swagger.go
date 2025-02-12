package utils

import (
	"os"
	"os/exec"
)

func SwaggerInitCmd(basePath ...string) error {
	var path string
	if len(basePath) == 0 {
		path = "./"
	} else {
		path = basePath[0]
	}
	cmd := exec.Command("swag", "init")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
