package generate

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"path/filepath"
)

func GenerateFileExistsHandler(output string, tmpl string, params map[string]string, override bool) error {
	return core.TryCatchVoid(func() {
		//pathArr := strings.Split(output, "/")
		//outputPathSlice := pathArr[:len(pathArr)-1]
		//outputPath := strings.Join(outputPathSlice, "/")
		outputPath := filepath.Dir(output)
		isDir, exists, err := core.Exists(outputPath)
		if err != nil {
			core.Throw(err)
		}
		if !exists || !isDir {
			core.RaiseVoid(core.CreateDir(outputPath))
		}
		isDir, exists, err = core.Exists(output)
		if err != nil {
			core.Throw(err)
		}
		if exists && !isDir && !override {
			core.Throw(fmt.Errorf("❌ '%s' has already existed!", output))
		}
		core.RaiseVoid(core.GenerateFromTemplateFile(tmpl, output, params))
	}, core.DefaultErrorHandler)
}
