package main

import (
	"net/http"
)

// Ref: https://helldiverstrainingmanual.com/api
const HELLDIVERS_URL string = "https://helldiverstrainingmanual.com/api/v1/war"

type WarStatusResponse struct {
	WarId            int     `json:"warId"`
	Time             int     `json:"time"`
	ImpactMultiplier float32 `json:"impactMultiplier"`
	StoryBeatId32    int32   `json:"storyBeatId32"`
	PlanetStatus     []struct {
		Index          int     `json:"index"`
		Owner          int     `json:"owner"`
		Health         int     `json:"health"`
		RegenPerSecond float32 `json:"regenPerSecond"`
	} `json:"planetStatus"`
}

type WarInfoResponse struct {
	WarId                int    `json:"warId"`
	StartDate            int    `json:"startDate"`
	EndDate              int    `json:"endDate"`
	MinimumClientVersion string `json:"minimumClientVersion"`
	PlanetInfos          []struct {
		Index        int `json:"index"`
		SettingsHash int `json:"settingsHash"`
		Position     struct {
			X int `json:"x"`
			Y int `json:"y"`
		} `json:"position"`
		Waypoints   []int `json:"waypoints"`
		Sector      int   `json:"sector"`
		MaxHealth   int   `json:"maxHealth"`
		Disabled    bool  `json:"disabled"`
		InitalOwner int   `json:"initalOwner"`
	} `json:"planetInfos"`
}

// Response should return "WarStatusResponse"
func GetWarStatus() (*http.Response, error) {
	url := HELLDIVERS_URL + "/status"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return MakeRequest(request)
}

// Response should return "WarInfoResponse"
func GetWarInfo() (*http.Response, error) {
	url := HELLDIVERS_URL + "/info"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	return MakeRequest(request)
}
