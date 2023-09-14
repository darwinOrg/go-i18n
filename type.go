package dgi18n

type (
	// GetLngHandler ...
	GetLngHandler = func(lng string, defaultLng string) string

	// Option ...
	Option func(DgI18n)
)
