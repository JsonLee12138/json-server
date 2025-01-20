package generate

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/embed"
	"path"
	"slices"
	"strings"
)

func GenerateModule(moduleName, outPath string) error {
	return core.TryCatchVoid(func() {
		dirs := []string{
			"controller",
			"service",
			"repository",
			"entity",
		}
		for _, dir := range dirs {
			dirPath := path.Join(outPath, dir)
			core.RaiseVoid(core.CreateDir(dirPath))
		}
		if slices.Contains(dirs, "controller") {
			core.RaiseVoid(GenerateController(moduleName, outPath, false, moduleName))
		}
		if slices.Contains(dirs, "service") {
			core.RaiseVoid(GenerateService(moduleName, outPath, false, moduleName))
		}
		if slices.Contains(dirs, "repository") {
			core.RaiseVoid(GenerateRepository(moduleName, outPath, false))
		}
		if slices.Contains(dirs, "entity") {
			core.RaiseVoid(GenerateEntity(moduleName, outPath, false))
		}
		core.RaiseVoid(GenerateEntry(outPath))
	}, func(err error) error {
		return err
	})
}

func GenerateEntry(outPath string) error {
	pathArr := strings.Split(outPath, "/")
	moduleName := pathArr[len(pathArr)-1]
	currentPath := path.Join(outPath, "entry.go")
	isDir, has, err := core.Exists(currentPath)
	if err != nil {
		return err
	}
	if has && !isDir {
		return fmt.Errorf("❌ '%s' has already existed!\n", moduleName)
	}
	return core.TryCatchVoid(func() {
		tmplFile := core.Raise(embed.TemplatesPath.ReadFile("templates/module.tmpl"))
		pkgPath := core.Raise(core.GetModuleFullPath(moduleName))
		upperName := core.UpperCamelCase(moduleName)
		core.RaiseVoid(core.GenerateFromTemplateFile(string(tmplFile), currentPath, map[string]string{
			"Name":      moduleName,
			"UpperName": upperName,
			"PkgPath":   pkgPath,
		}))
		fmt.Printf("✅ '%s' module has been successfully generated!\n", moduleName)
	}, core.DefaultErrorHandler)
}
