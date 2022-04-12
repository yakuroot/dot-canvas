package locales

import (
	"github.com/diamondburned/arikawa/v3/discord"
)

var (
	Languages = []discord.Language{discord.Korean, discord.EnglishUS}
)

func Text(key, lng string, opt ...map[string]interface{}) string {
	var t string

	if len(opt) > 0 && opt[0] != nil {
		t = locale.Tr(lng, key, opt[0])
	} else {
		t = locale.Tr(lng, key)
	}

	if t == "" {
		return "Locale Error"
	}

	return t
}

func IsSupportLanguage(lng string) bool {
	for _, l := range Languages {
		if string(l) == lng {
			return true
		}
	}
	return false
}
