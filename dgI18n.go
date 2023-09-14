package dgi18n

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var _ DgI18n = (*dgI18nImpl)(nil)

type dgI18nImpl struct {
	bundle          *i18n.Bundle
	localizerByLng  map[string]*i18n.Localizer
	defaultLanguage language.Tag
	getLngHandler   GetLngHandler
}

// getMessage get localize message by lng and messageID
func (i *dgI18nImpl) getMessage(lg string, param interface{}) (string, error) {
	lng := i.getLngHandler(lg, i.defaultLanguage.String())
	localizer := i.getLocalizerByLng(lng)

	localizeConfig, err := i.getLocalizeConfig(param)
	if err != nil {
		return "", err
	}

	message, err := localizer.Localize(localizeConfig)
	if err != nil {
		return "", err
	}

	return message, nil
}

// mustGetMessage ...
func (i *dgI18nImpl) mustGetMessage(lng string, param interface{}) string {
	message, _ := i.getMessage(lng, param)
	return message
}

func (i *dgI18nImpl) setBundle(cfg *BundleCfg) {
	bundle := i18n.NewBundle(cfg.DefaultLanguage)
	bundle.RegisterUnmarshalFunc(cfg.FormatBundleFile, cfg.UnmarshalFunc)

	i.bundle = bundle
	i.defaultLanguage = cfg.DefaultLanguage

	i.loadMessageFiles(cfg)
	i.setLocalizerByLng(cfg.AcceptLanguage)
}

func (i *dgI18nImpl) setGetLngHandler(handler GetLngHandler) {
	i.getLngHandler = handler
}

// loadMessageFiles load all file localize to bundle
func (i *dgI18nImpl) loadMessageFiles(config *BundleCfg) {
	WriteCommonYamlFile(config.RootPath)
	lngFlags, _ := os.ReadDir(config.RootPath)

	for _, lng := range config.AcceptLanguage {
		for _, lngFlag := range lngFlags {
			if lngFlag.IsDir() {
				path := filepath.Join(config.RootPath, lngFlag.Name(), lngFlag.Name()+"."+lng.String()) + "." + config.FormatBundleFile
				if err := i.loadMessageFile(config, lngFlag.Name(), path); err != nil {
					panic(err)
				}
			}
		}
	}
}

func (i *dgI18nImpl) loadMessageFile(config *BundleCfg, parentDir string, path string) error {
	buf, err := config.Loader.LoadMessage(path)
	if err != nil {
		return err
	}

	content := string(buf)
	rows := strings.Split(content, "\n")
	for i, row := range rows {
		if row != "" && strings.Trim(row, " ") != "" {
			rows[i] = parentDir + "." + row
		}
	}
	buf = []byte(strings.Join(rows, "\n"))

	if _, err = i.bundle.ParseMessageFileBytes(buf, path); err != nil {
		return err
	}
	return nil
}

// setLocalizerByLng set localizer by language
func (i *dgI18nImpl) setLocalizerByLng(acceptLanguage []language.Tag) {
	i.localizerByLng = map[string]*i18n.Localizer{}
	for _, lng := range acceptLanguage {
		lngStr := lng.String()
		i.localizerByLng[lngStr] = i.newLocalizer(lngStr)
	}

	// set defaultLanguage if it isn't exist
	defaultLng := i.defaultLanguage.String()
	if _, hasDefaultLng := i.localizerByLng[defaultLng]; !hasDefaultLng {
		i.localizerByLng[defaultLng] = i.newLocalizer(defaultLng)
	}
}

// newLocalizer create a localizer by language
func (i *dgI18nImpl) newLocalizer(lng string) *i18n.Localizer {
	lngDefault := i.defaultLanguage.String()
	lngs := []string{
		lng,
	}

	if lng != lngDefault {
		lngs = append(lngs, lngDefault)
	}

	localizer := i18n.NewLocalizer(
		i.bundle,
		lngs...,
	)
	return localizer
}

// getLocalizerByLng get localizer by language
func (i *dgI18nImpl) getLocalizerByLng(lng string) *i18n.Localizer {
	localizer, hasValue := i.localizerByLng[lng]
	if hasValue {
		return localizer
	}

	return i.localizerByLng[i.defaultLanguage.String()]
}

func (i *dgI18nImpl) getLocalizeConfig(param interface{}) (*i18n.LocalizeConfig, error) {
	switch paramValue := param.(type) {
	case string:
		localizeConfig := &i18n.LocalizeConfig{
			MessageID: paramValue,
		}
		return localizeConfig, nil
	case *i18n.LocalizeConfig:
		return paramValue, nil
	}

	msg := fmt.Sprintf("un supported localize param: %v", param)
	return nil, errors.New(msg)
}
