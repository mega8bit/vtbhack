package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func L(MessageID string, Loc i18n.Localizer) string {
	return Loc.MustLocalize(&i18n.LocalizeConfig{MessageID: MessageID})
}
func langInit() {
	bundle = i18n.NewBundle(language.Russian)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	fmt.Println(*flagLocale)

	if err := filepath.Walk(*flagLocale, func(path string, file os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".yaml") {
			log.Println("Load translation file", file.Name())
			bundle.MustLoadMessageFile(path)
			locales = append(locales, strings.TrimSuffix(file.Name(), ".yaml"))
		}
		return nil
	}); err != nil {
		log.Fatalln(err.Error())
	}
	getLangs()
}

func getLangs() {
	langs := bundle.LanguageTags()
	supported := "Languages: "
	for _, lang := range langs {
		loc := i18n.NewLocalizer(bundle, lang.String())
		supported += " " + L("language", *loc)
	}
	fmt.Println(supported)

}
