package tg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Update struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

// IsCommand returns boolean value if command and command with parameters if true
func (u *Update) IsCommand() (bool, string, string) {
	for _, ent := range u.Message.Entities {
		if ent.Type == "bot_command" {
			broken := strings.Split(u.Message.Text, " ")
			if len(broken) < 1 {
				continue
			}
			joined := strings.Join(broken[1:], " ")
			return true, broken[0], strings.TrimSpace(joined)
		}
	}

	return false, "", ""
}

type GetUpdatesResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

func (c *Client) getUpdates() (*GetUpdatesResponse, error) {
	cl := http.Client{Timeout: 10 * time.Second}

	url := c.baseUrl + "getUpdates"

	empty := fmt.Sprintf(`{"offset": %d}`, c.lastUpdateID)

	res, err := cl.Post(url, "application/json", bytes.NewBuffer([]byte(empty)))
	if err != nil {
		return nil, err
	}

	var jsonRes GetUpdatesResponse

	rawbytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.Unmarshal(rawbytes, &jsonRes); err != nil {
		return nil, fmt.Errorf("failed to unmarshall response: %v, err: <%v>", rawbytes, err)
	}

	return &jsonRes, nil
}
