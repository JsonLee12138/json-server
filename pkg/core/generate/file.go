package generate

import (
	"fmt"
	core2 "github.com/JsonLee12138/json-server/pkg/core"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"path/filepath"
)

func GenerateFileExistsHandler(output string, tmpl string, params map[string]string, override bool) error {
	return utils.TryCatchVoid(func() {
		//pathArr := strings.Split(output, "/")
		//outputPathSlice := pathArr[:len(pathArr)-1]
		//outputPath := strings.Join(outputPathSlice, "/")
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
		utils.RaiseVoid(core2.GenerateFromTemplateFile(tmpl, output, params))
	}, utils.DefaultErrorHandler)
}
