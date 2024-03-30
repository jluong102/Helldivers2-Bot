package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const DISCORD_URL string = "https://discord.com/api/v10"
const MAX_MSG_LEN = 2000

type Discord struct {
	token string
}

type CreateMessagePayload struct {
	Content          string        `json:"content"`
	Nonce            int           `json:"nonce"`
	Tts              bool          `json:"tts"`
	Embeds           []interface{} `json:"embeds"`
	AllowedMentions  interface{}   `json:"allowed_mentions"`
	MessageReference interface{}   `json:"message_reference"`
	Components       []interface{} `json:"components"`
	StickerIds       interface{}   `json:"sticker_ids"`
	Files            string        `json:"files"`
	PayloadJson      string        `json:"payload_json"`
	Attachments      interface{}   `json:"attachments"`
	Flags            int           `json:"flags"`
	EnforceNonce     bool          `json:"enforce_nonce"`
}

// Ref: https://discord.com/developers/docs/resources/channel#create-message
func (this Discord) CreateMessage(channel string, payload *CreateMessagePayload) (*http.Response, error) {
	url := fmt.Sprintf("%s/channels/%s/messages", DISCORD_URL, channel)

	if len(payload.Content) > MAX_MSG_LEN {
		err := fmt.Errorf("Unable to send message. Max length %d, message length %d", MAX_MSG_LEN, len(payload.Content))
		return nil, err
	} else if len(payload.Content) < 1 {
		err := fmt.Errorf("Unable to send message. \"Content\" must be set")
		return nil, err
	}

	jsonPayload, err := json.Marshal(payload)

	// Confirm json encoded
	if err != nil {
		return nil, err
	}

	// Create HTTP request
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	SetJson(request)

	if err != nil {
		return nil, err
	}

	return this.MakeRequest(request)
}

func (this Discord) MakeRequest(request *http.Request) (*http.Response, error) {
	request.Header.Add("Authorization", fmt.Sprintf("Bot: %s", this.token))
	request.Header.Add("User-agent", "DiscordBot (Helldivers2-Bot; 1)")

	client := &http.Client{}
	return client.Do(request)
}

// Constructor
func CreateDiscord(token string) *Discord {
	discord := new(Discord)
	discord.token = token

	return discord
}
