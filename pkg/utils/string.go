package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	c := cases.Title(language.English)
	s = c.String(s)
	return strings.Replace(s, " ", "", -1)
}
