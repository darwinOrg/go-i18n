package dgi18n

import (
	"gopkg.in/yaml.v3"
	"os"

	"golang.org/x/text/language"
)

const (
	defaultFormatBundleFile = "yaml"
	defaultRootPath         = "configs/i18n"
)

var (
	defaultLanguage       = language.Chinese
	defaultUnmarshalFunc  = yaml.Unmarshal
	defaultAcceptLanguage = []language.Tag{
		defaultLanguage,
		language.English,
	}

	defaultLoader = LoaderFunc(os.ReadFile)

	defaultBundleConfig = &BundleCfg{
		RootPath:         defaultRootPath,
		AcceptLanguage:   defaultAcceptLanguage,
		FormatBundleFile: defaultFormatBundleFile,
		DefaultLanguage:  defaultLanguage,
		UnmarshalFunc:    defaultUnmarshalFunc,
		Loader:           defaultLoader,
	}
)
