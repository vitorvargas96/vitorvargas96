package main

import (
	"dev-timer/pkgs/client"
	textHandler "dev-timer/pkgs/text-handler"
	"fmt"
	"os"
	"strings"
)

const ALL_TIME = ""
const TODAY = "today"
const WEEK = "week"

const README_PATH = "README.md"

func main() {
	updateDevTimer(TODAY)
	updateDevTimer(WEEK)
	updateDevTimer(ALL_TIME)
}

func updateDevTimer(rangeTime string) {
	data, err := client.GetTimer(rangeTime)
	fmt.Println(data)

	if err != nil {
		fmt.Println(err)
		return
	}

	text := textHandler.Handler(data.Data.Languages)

	err = updateReadme(text, rangeTime)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getReadme() ([]byte, error) {
	file, err := os.ReadFile(README_PATH)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func updateReadme(langs string, rangeTime string) error {
	readme, err := getReadme()

	if err != nil {
		fmt.Println(err)
		return err
	}

	rangeTime = strings.ToUpper(rangeTime)

	part1 := strings.Split(string(readme), "<!--DEVTIMER:"+rangeTime+":START-->")[0]
	part2 := strings.Split(string(readme), "<!--DEVTIMER:"+rangeTime+":END-->")[1]

	result := part1 + "<!--DEVTIMER:" + rangeTime + ":START-->\n" + langs + "<!--DEVTIMER:" + rangeTime + ":END-->" + part2

	err = os.WriteFile(README_PATH, []byte(result), 0644)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
