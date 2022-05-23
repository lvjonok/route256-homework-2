package tg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Message struct {
	Ok     bool `json:"ok"`
	Result struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID        int64  `json:"id"`
			IsBot     bool   `json:"is_bot"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"result"`
}

func (c *Client) SendMessage(chatID int, text string) (*Message, error) {
	cl := http.Client{Timeout: 10 * time.Second}

	url := c.baseUrl + "sendMessage"

	load := fmt.Sprintf(`{"chat_id": %d, "text": "%s"}`, chatID, text)

	res, err := cl.Post(url, "application/json", bytes.NewBuffer([]byte(load)))
	if err != nil {
		return nil, err
	}

	rawbytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp Message
	if err := json.Unmarshal(rawbytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response %v, err: <%v>", rawbytes, err)
	}

	return &resp, nil
}
