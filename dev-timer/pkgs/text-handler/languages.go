package textHandler

import (
	"dev-timer/pkgs/client"
	"fmt"
)

func generateLanguageText(languages []client.Languages, totalTime string) string {
	var text string

	badges, err := addLanguageBadges(languages)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text += badges + "\n\n"

	text += "```txt\n"

	text += fmt.Sprintf("Total Time: %s\n\n", totalTime)

	for _, language := range languages {
		langName := language.Name
		if langName == "unknown" || language.Percent == 0 || language.Minutes == 0 {
			continue
		}

		progressBar := generateProgressBar(language.Percent)

		if langName == "Nginx configuration file" {
			langName = "Nginx"
		}

		text += fmt.Sprintf("%s%s %s%s %s    %.2f %%\n", langName, generateSpaceAfterLanguage(langName), language.Text, generateSpaceAfterTime(language.Text), progressBar, language.Percent)
	}

	text += "```\n\n"

	return text
}

func generateSpaceAfterLanguage(language string) string {
	var spaceAfterLanguage string
	var languageSize = len(language)
	var spaceAfterLanguageCalc = 19 - languageSize

	for i := 0; i < spaceAfterLanguageCalc; i++ {
		spaceAfterLanguage += " "
	}

	return spaceAfterLanguage
}

func generateSpaceAfterTime(time string) string {
	var spaceAfterTime string
	var timeSize = len(time)
	var spaceAfterTimeCalc = 15 - timeSize

	for i := 0; i < spaceAfterTimeCalc; i++ {
		spaceAfterTime += " "
	}

	return spaceAfterTime
}

func generatePercent(percent float64) string {
	var percentString string
	var percentCalc = int(percent * 25 / 100)

	percentString += "["

	for i := 0; i < percentCalc; i++ {
		percentString += ">"
	}

	for i := 0; i < 25-percentCalc; i++ {
		percentString += "="
	}

	percentString += "]"

	return percentString
}

func addLanguageBadges(languages []client.Languages) (string, error) {
	var badges string

	for _, language := range languages {
		if language.Name == "unknown" || language.Percent == 0 || language.Minutes == 0 {
			continue
		}

		icon, err := getIcon(language.Name)

		if err != nil {
			continue
		}

		badges += "<img align=\"center\" width=\"32px\" src=\"" + icon + "\" alt=\"" + language.Name + "\" />"
		badges += "&nbsp;&nbsp;&nbsp;"
	}

	return badges, nil
}
