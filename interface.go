package dgi18n

// DgI18n ...
type DgI18n interface {
	getMessage(lng string, param interface{}) (string, error)
	mustGetMessage(lng string, param interface{}) string
	setBundle(cfg *BundleCfg)
	setGetLngHandler(handler GetLngHandler)
}
