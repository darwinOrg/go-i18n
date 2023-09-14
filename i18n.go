package dgi18n

import (
	dgerr "github.com/darwinOrg/go-common/enums/error"
	"github.com/darwinOrg/go-common/result"
)

var atI18n DgI18n

// Localize ...
func Localize(opts ...Option) {
	newI18n(opts...)
}

// newI18n ...
func newI18n(opts ...Option) {
	// init ins
	ins := &dgI18nImpl{}

	// set ins property from opts
	for _, opt := range opts {
		opt(ins)
	}

	// 	if bundle isn't constructed then assign it from default
	if ins.bundle == nil {
		ins.setBundle(defaultBundleConfig)
	}

	// if getLngHandler isn't constructed then assign it from default
	if ins.getLngHandler == nil {
		ins.getLngHandler = defaultGetLngHandler
	}

	atI18n = ins
}

/*
GetMessage get the i18n message

	 param is one of these type: messageID, *i18n.LocalizeConfig
	 Example:
		GetMessage("hello") // messageID is hello
		GetMessage(&i18n.LocalizeConfig{
				MessageID: "welcomeWithName",
				TemplateData: map[string]string{
					"name": context.Param("name"),
				},
		})
*/
func GetMessage(lng string, param interface{}) (string, error) {
	return atI18n.getMessage(lng, param)
}

/*
MustGetMessage get the i18n message without error handling

	  param is one of these type: messageID, *i18n.LocalizeConfig
	  Example:
		MustGetMessage("hello") // messageID is hello
		MustGetMessage(&i18n.LocalizeConfig{
				MessageID: "welcomeWithName",
				TemplateData: map[string]string{
					"name": context.Param("name"),
				},
		})
*/
func MustGetMessage(lng string, param interface{}) string {
	return atI18n.mustGetMessage(lng, param)
}

func GetMessageByDgErrorML(lng string, e *dgerr.DgErrorML) (string, error) {
	return atI18n.getMessage(lng, e.MessageCode)
}

func MustGetMessageByDgErrorML(lng string, e *dgerr.DgErrorML) string {
	return atI18n.mustGetMessage(lng, e.MessageCode)
}

func BuildResultMLByError[T any](lng string, err error) *result.ResultML[T] {
	switch err.(type) {
	case *dgerr.DgErrorML:
		return BuildResultMLByDgErrorML[T](lng, err.(*dgerr.DgErrorML))
	case *dgerr.DgError:
		de := err.(*dgerr.DgError)
		return &result.ResultML[T]{
			Success: false,
			Code:    de.Code,
			Message: de.Message,
		}
	default:
		return BuildResultMLByDgErrorML[T](lng, dgerr.SYSTEM_ERROR_ML)
	}
}

func BuildResultMLByDgErrorML[T any](lng string, e *dgerr.DgErrorML) *result.ResultML[T] {
	return &result.ResultML[T]{
		Success:     false,
		Code:        e.Code,
		MessageCode: e.MessageCode,
		Message:     MustGetMessageByDgErrorML(lng, e),
	}
}

func FillResultMLMessageByDgErrorML[T any](lng string, e *dgerr.DgErrorML, r *result.ResultML[T]) {
	r.Message = MustGetMessageByDgErrorML(lng, e)
}
