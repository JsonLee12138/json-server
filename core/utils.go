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

func GetModulePath() (string, error) {
	cmd := exec.Command("go", "list", "-m", "-f", "{{.Dir}}")
	var out bytes.Buffer
	cmd.Stdout = &out

	// 执行命令并捕获错误
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get project directory: %v", err)
	}

	// 去除换行符，返回结果
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
	modulePath, err := GetModulePath()
	if err != nil {
		return "", err
	}
	pathArr := strings.Split(currentDir, modulePath)
	var actualPath string
	if len(pathArr) < 2 {
		actualPath = ""
	} else {
		actualPath = pathArr[1]
	}
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
	if strings.HasSuffix(res, "/") {
		res = strings.TrimSuffix(res, "/")
	}
	return res, nil
}

func UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	c := cases.Title(language.English)
	s = c.String(s)
	return strings.Replace(s, " ", "", -1)
}

func FindPIDByPort(port string) (string, error) {
	cmd := exec.Command("lsof", "-t", "-i", ":"+port)
	output, err := cmd.Output()
	if err != nil {
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
