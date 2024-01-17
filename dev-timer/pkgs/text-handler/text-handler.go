package textHandler

import "dev-timer/pkgs/client"

func Handler(data client.Data) string {
	languages := generateLanguageText(data.Languages, data.HumanReadableTotal)

	return languages
}
