package core

import (
	"bytes"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"os/exec"
	"strings"
)

func GetModuleName() (string, error) {
	cmd := exec.Command("go", "list", "-m")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get module name: %v", err)
	}
	return strings.TrimSpace(out.String()), nil
}

func GetModuleFullPath(module string) (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	moduleName, err := GetModuleName()
	if err != nil {
		return "", err
	}
	actualPath := strings.Split(currentDir, moduleName)[1]
	res := moduleName
	if actualPath != "" {
		if strings.HasPrefix(actualPath, "/") {
			res += actualPath
		} else {
			res += "/" + actualPath
		}
	}
	if strings.HasPrefix(module, "/") {
		res += module
	} else {
		res += "/" + module
	}
	return res, nil
}

func UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	c := cases.Title(language.English)
	s = c.String(s)
	return strings.Replace(s, " ", "", -1)
}
