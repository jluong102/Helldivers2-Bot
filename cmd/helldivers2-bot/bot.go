package main

import (
	"fmt"
	"time"
	"log"
	"io"
	"encoding/json"
)

func StartBot(discord *Discord, settings *Settings) {
	for {
		showCampaign(discord, settings)
		time.Sleep(time.Hour)
	}
}

func showCampaign(discord *Discord, settings *Settings) {
	campaign := getParsedCampaignInfo()

	// Try again later if HTTP request fails
	if campaign == nil {
		return
	}

	msgList := createCampaignMsg(campaign)

	// Print out campaign info to all listed channels
	for _, channel := range(settings.Channels) {
		for _, msg := range msgList {
			sendMsg(discord, channel, msg)
		}
	}
}

func getParsedCampaignInfo() *WarCampaignResponse {
	info := new(WarCampaignResponse)

	response, err := GetWarCampaign() 

	if err != nil {
		log.Printf("Failed to make HTTP request\nError %s", err)
		return nil
	}

	data, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err := json.Unmarshal(data, info); err != nil {
		log.Printf("Failed to parse JSON\nError %s", err)
		return nil
	}

	return info	
}

func createCampaignMsg(info *WarCampaignResponse) []string {
	var planetMsg []string

	for _, i := range info {
		msg := "```\n"
		msg += fmt.Sprintf("Planet: %s\n", i.Name)
		msg += fmt.Sprintf("Health: %d/%d (%f)\n", i.Health, i.MaxHealth, i.Percentage)
		msg += fmt.Sprintf("Defense: %t", i.Defense)
		msg += fmt.Sprintf("Major Order: %t\n", i.MajorOrder)
		msg += fmt.Sprintf("Biome: %s\n", i.Biome.Slug)
		msg += "```"
		
		planetMsg = append(planetMsg, msg)
	}

	return planetMsg
}

func sendMsg(discord *Discord, channel string, msg string) {
	payload := new(CreateMessagePayload)	
	payload.Content = msg

	response, _ := CreateMessage(channel, payload)
	
	if response.StatusCode != 200 {
		log.Printf("Error: %s", response.Status)
	}
}
