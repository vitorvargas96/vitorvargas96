package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseUrl = "https://devtimer.vitorvargas.dev/"

type devTimerClient struct {
	httpClient *http.Client
}

type DevTimer struct {
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

func NewHttpClient() *devTimerClient {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}

	return &devTimerClient{
		httpClient: httpClient,
	}
}

func (c devTimerClient) GetData() (string, error) {
	res, err := c.httpClient.Get(baseUrl)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		return "", fmt.Errorf("expected status OK; got %v", res.Status)
	}

	return string(data), nil
}

func GetWeeklyTimer() (DevTimer, error) {
	var data DevTimer
	content, err := NewHttpClient().GetData()

	if err != nil {
		return data, err
	}

	err = json.Unmarshal([]byte(content), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
