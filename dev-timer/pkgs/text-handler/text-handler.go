package textHandler

import (
	"dev-timer/pkgs/client"
)

func Handler(data []client.Languages) string {
	languages := generateLanguageText(data)

	return languages
}
