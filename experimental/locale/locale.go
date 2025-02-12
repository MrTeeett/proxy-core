package locale

var (
	localeRegistry = make(map[string]*Locale)
	current        = defaultLocal
)

type Locale struct {
	// deprecated messages for graphical clients
	Locale                  string
	DeprecatedMessage       string
	DeprecatedMessageNoLink string
}

var defaultLocal = &Locale{
	Locale:                  "en_US",
	DeprecatedMessage:       "%s is deprecated in proxy-core %s and will be removed in proxy-core %s please checkout documentation for migration.",
	DeprecatedMessageNoLink: "%s is deprecated in proxy-core %s and will be removed in proxy-core %s.",
}

func Current() *Locale {
	return current
}

func Set(localeId string) bool {
	locale, loaded := localeRegistry[localeId]
	if !loaded {
		return false
	}
	current = locale
	return true
}
