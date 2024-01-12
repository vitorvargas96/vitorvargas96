package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type File struct {
	Data Data `json:"data"`
}

type Data struct {
	Username                  string      `json:"username"`
	UserID                    string      `json:"user_id"`
	Start                     string      `json:"start"`
	End                       string      `json:"end"`
	Status                    string      `json:"status"`
	TotalSeconds              int         `json:"total_seconds"`
	DailyAverage              float64     `json:"daily_average"`
	DaysIncludingHolidays     int         `json:"days_including_holidays"`
	Range                     string      `json:"range"`
	HumanReadableRange        string      `json:"human_readable_range"`
	HumanReadableTotal        string      `json:"human_readable_total"`
	HumanReadableDailyAverage string      `json:"human_readable_daily_average"`
	IsCodingActivityVisible   bool        `json:"is_coding_activity_visible"`
	IsOtherUsageVisible       bool        `json:"is_other_usage_visible"`
	Editors                   []string    `json:"editors"`
	Languages                 []Languages `json:"languages"`
	Machines                  []string    `json:"machines"`
	Projects                  []string    `json:"projects"`
	OperatingSystems          []string    `json:"operating_systems"`
}

type Languages struct {
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Seconds      int     `json:"seconds"`
	Text         string  `json:"text"`
	TotalSeconds int     `json:"total_seconds"`
}

func main() {
	data, err := loadJSON("data.json")

	if err != nil {
		fmt.Println(err)
		return
	}

	langs := getLanguages(data)

	err = updateReadme(langs)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func loadJSON(filename string) (File, error) {
	var data File
	content, err := os.ReadFile(filename)

	if err != nil {
		return data, err
	}

	err = json.Unmarshal(content, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getLanguages(data File) string {
	var result string

	result += "```txt\n"

	result += "Dev Weekly Status\n\n"

	result += fmt.Sprintf("Total time: %s\n\n", data.Data.HumanReadableTotal)

	blocks := []string{"░", "▒", "▓", "█"}

	for _, language := range data.Data.Languages {
		if language.Name == "unknown" || language.Percent == 0 || language.Minutes == 0 {
			continue
		}

		languageSize := len(language.Name)
		spaceAfterLanguageCalc := 19 - languageSize
		var spaceAfterLanguage string

		for i := 0; i < spaceAfterLanguageCalc; i++ {
			spaceAfterLanguage += " "
		}

		percent := int(language.Percent * 25 / 100)

		var percentString string

		for i := 0; i < percent; i++ {
			percentString += string(blocks[3])
		}

		for i := 0; i < 25-percent; i++ {
			percentString += string(blocks[0])
		}

		time := language.Text
		spaceAfterTimeCalc := 15 - len(time)
		var spaceAfterTime string

		for i := 0; i < spaceAfterTimeCalc; i++ {
			spaceAfterTime += " "
		}

		result += fmt.Sprintf("%s%s %s%s %s    %.2f %%\n", language.Name, spaceAfterLanguage, time, spaceAfterTime, percentString, language.Percent)

		// result += fmt.Sprintf("%s %d hrs %d mins", language.Name, language.Hours, language.Minutes)
	}

	result += "```\n"

	return result
}

func getReadme() ([]byte, error) {
	file, err := os.ReadFile("../README.md")

	if err != nil {
		return nil, err
	}

	return file, nil
}

func updateReadme(langs string) error {
	readme, err := getReadme()

	if err != nil {
		fmt.Println(err)
		return err
	}

	part1 := strings.Split(string(readme), "<!--DEVTIMER:START-->")[0]
	part2 := strings.Split(string(readme), "<!--DEVTIMER:END-->")[1]

	result := part1 + "<!--DEVTIMER:START-->\n" + langs + "<!--DEVTIMER:END-->" + part2

	err = os.WriteFile("../README.md", []byte(result), 0644)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
