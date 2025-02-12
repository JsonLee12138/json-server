package core

import (
	"encoding/json"

	"github.com/JsonLee12138/json-server/pkg/configs"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

func NewI18n(config ...configs.I18nConfig) fiber.Handler {
	if len(config) == 0 {
		return fiberi18n.New()
	}
	cnf := config[0]
	defaultLanguage, err := language.Parse(cnf.DefaultLanguage)
	if err != nil {
		defaultLanguage = language.English
	}
	var acceptLanguages []language.Tag
	for _, v := range cnf.AcceptLanguages {
		acceptLanguage, err := language.Parse(v)
		if err == nil {
			acceptLanguages = append(acceptLanguages, acceptLanguage)
		}
	}
	var unmarshalFunc i18n.UnmarshalFunc
	if cnf.FormatBundleFile == "json" {
		unmarshalFunc = json.Unmarshal
	} else {
		unmarshalFunc = yaml.Unmarshal
	}
	return fiberi18n.New(&fiberi18n.Config{
		RootPath:         cnf.RootPath,
		AcceptLanguages:  acceptLanguages,
		DefaultLanguage:  defaultLanguage,
		FormatBundleFile: cnf.FormatBundleFile,
		UnmarshalFunc:    unmarshalFunc,
	})
}
