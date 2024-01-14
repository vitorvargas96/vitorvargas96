package main

import (
	"dev-timer/pkgs/client"
	textHandler "dev-timer/pkgs/text-handler"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := client.GetWeeklyTimer()
	fmt.Println(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	text := textHandler.Handler(data.Data.Languages)

	err = updateReadme(text)

	if err != nil {
		fmt.Println(err)
		return
	}
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
