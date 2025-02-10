package core

import (
	"os"
	"strings"
	"text/template"
)

func GenerateFromTemplate(templatePath string, outPath string, params map[string]string) error {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	return generate(tmpl, outPath, params)
}

func GenerateFromTemplateFile(templateContent, outPath string, params map[string]string) error {
	tmpl, err := template.New("model").Parse(templateContent)
	if err != nil {
		return err
	}
	return generate(tmpl, outPath, params)
}

func generate(tmpl *template.Template, outPath string, params map[string]string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	outDir := strings.Replace(outPath, currentDir, currentDir, -1)
	f, err := os.Create(outDir)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tmpl.Execute(f, params)
	if err != nil {
		return err
	}
	return nil
}
