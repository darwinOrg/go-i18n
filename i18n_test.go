package dgi18n

import (
	"fmt"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestI18nEN(t *testing.T) {
	initI18n()
	fmt.Println(MustGetMessage("en", "common.system_error"))
}

func TestI18nZH(t *testing.T) {
	initI18n()
	fmt.Println(MustGetMessage("zh", "common.system_error"))
}

func initI18n() {
	Localize(WithBundle(&BundleCfg{
		DefaultLanguage:  language.Chinese,
		AcceptLanguage:   []language.Tag{language.English, language.Chinese},
		RootPath:         "_example",
		FormatBundleFile: "yaml",
		UnmarshalFunc:    yaml.Unmarshal,
	}))
}
