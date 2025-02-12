package core

import (
	"fmt"
	"path/filepath"

	"github.com/JsonLee12138/json-server/pkg/utils"
)

func GenerateFileExistsHandler(output string, tmpl string, params map[string]string, override bool) error {
	return utils.TryCatchVoid(func() {
		outputPath := filepath.Dir(output)
		isDir, exists, err := utils.Exists(outputPath)
		if err != nil {
			utils.Throw(err)
		}
		if !exists || !isDir {
			utils.RaiseVoid(utils.CreateDir(outputPath))
		}
		isDir, exists, err = utils.Exists(output)
		if err != nil {
			utils.Throw(err)
		}
		if exists && !isDir && !override {
			utils.Throw(fmt.Errorf("❌ '%s' has already existed!", output))
		}
		utils.RaiseVoid(GenerateFromTemplateFile(tmpl, output, params))
	}, utils.DefaultErrorHandler)
}
