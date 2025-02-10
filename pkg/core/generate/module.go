package generate

import (
	"fmt"
	"github.com/JsonLee12138/json-server/embed"
	core2 "github.com/JsonLee12138/json-server/pkg/core"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"path"
	"slices"
	"strings"
)

func GenerateModule(moduleName, outPath string) error {
	return utils.TryCatchVoid(func() {
		dirs := []string{
			"controller",
			"service",
			"repository",
			"entity",
		}
		for _, dir := range dirs {
			dirPath := path.Join(outPath, dir)
			utils.RaiseVoid(utils.CreateDir(dirPath))
		}
		if slices.Contains(dirs, "controller") {
			utils.RaiseVoid(GenerateController(moduleName, outPath, false, moduleName))
		}
		if slices.Contains(dirs, "service") {
			utils.RaiseVoid(GenerateService(moduleName, outPath, false, moduleName))
		}
		if slices.Contains(dirs, "repository") {
			utils.RaiseVoid(GenerateRepository(moduleName, outPath, false))
		}
		if slices.Contains(dirs, "entity") {
			utils.RaiseVoid(GenerateEntity(moduleName, outPath, false))
		}
		utils.RaiseVoid(GenerateEntry(outPath))
	}, func(err error) error {
		return err
	})
}

func GenerateEntry(outPath string) error {
	pathArr := strings.Split(outPath, "/")
	moduleName := pathArr[len(pathArr)-1]
	currentPath := path.Join(outPath, "entry.go")
	isDir, has, err := utils.Exists(currentPath)
	if err != nil {
		return err
	}
	if has && !isDir {
		return fmt.Errorf("❌ '%s' has already existed!\n", moduleName)
	}
	return utils.TryCatchVoid(func() {
		tmplFile := utils.Raise(embed.TemplatesPath.ReadFile("templates/module.tmpl"))
		pkgPath := utils.Raise(utils.GetModuleFullPath(moduleName))
		upperName := utils.UpperCamelCase(moduleName)
		utils.RaiseVoid(core2.GenerateFromTemplateFile(string(tmplFile), currentPath, map[string]string{
			"Name":      moduleName,
			"UpperName": upperName,
			"PkgPath":   pkgPath,
		}))
		fmt.Printf("✅ '%s' module has been successfully generated!\n", moduleName)
	}, utils.DefaultErrorHandler)
}
