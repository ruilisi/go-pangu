package util

import (
	"strings"

	"github.com/gogf/gf/i18n/gi18n"
)

func Contains(s []string, target string) bool {
	for _, element := range s {
		if strings.EqualFold(element, target) {
			return true
		}
	}
	return false
}

func I18N(locale string) *gi18n.Manager {
	t := gi18n.New()
	parsedLocale := strings.TrimSpace(locale)
	if strings.HasPrefix(parsedLocale, "zh") {
		parsedLocale = "zh"
	}
	if strings.HasPrefix(parsedLocale, "en") {
		parsedLocale = "en"
	}
	if parsedLocale != "en" && parsedLocale != "zh" {
		parsedLocale = "zh"
	}
	t.SetLanguage(parsedLocale)
	return t
}
