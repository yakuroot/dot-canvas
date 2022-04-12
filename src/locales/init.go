package locales

import (
	"log"
	"os"

	"github.com/kataras/i18n"
)

var (
	locale *i18n.I18n
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	if locale, err = i18n.New(
		i18n.Glob(path+"/src/locales/*/*.json"),
		"ko-KR",
		"en-US",
		"ja-JP",
	); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
