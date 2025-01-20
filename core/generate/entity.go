package generate

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/embed"
)

func GenerateEntity(entityName, outPath string, override bool) error {
	return core.TryCatchVoid(func() {
		tmpl := string(core.Raise(embed.TemplatesPath.ReadFile("templates/entity.tmpl")))
		entityPath := fmt.Sprintf("%s/entity/%s_entity.go", outPath, entityName)
		upperName := core.UpperCamelCase(entityName)
		params := map[string]string{
			"Name": upperName,
		}
		core.RaiseVoid(GenerateFileExistsHandler(entityPath, tmpl, params, override))
		fmt.Printf("✅ '%s' entity has been successfully generated!\n", entityName)
	}, core.DefaultErrorHandler)
}
