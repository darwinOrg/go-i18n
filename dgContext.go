package dgi18n

// defaultGetLngHandler ...
func defaultGetLngHandler(lng string, defaultLng string) string {
	if lng != "" {
		return lng
	}

	return defaultLng
}
