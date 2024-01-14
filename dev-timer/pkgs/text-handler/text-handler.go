package textHandler

import (
	"dev-timer/pkgs/client"
	"fmt"
)

func Handler(data []client.Languages) string {
	fmt.Println("Hello World")

	fmt.Println(generateHeader("Dev Weekly Status"))

	languages := generateLanguageText(data)

	fmt.Println(languages)

	return languages
}
