package tg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MessagePhoto struct {
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
		Date  int `json:"date"`
		Photo []struct {
			FileID       string `json:"file_id"`
			FileUniqueID string `json:"file_unique_id"`
			FileSize     int    `json:"file_size"`
			Width        int    `json:"width"`
			Height       int    `json:"height"`
		} `json:"photo"`
		Caption string `json:"caption"`
	} `json:"result"`
}

func (c *Client) SendPhoto(chatID int, photoBytes []byte, caption string) (*MessagePhoto, error) {
	cl := http.Client{Timeout: 10 * time.Second}

	url := c.baseUrl + "sendPhoto"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// add chat_id
	fw, err := writer.CreateFormField("chat_id")
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, strings.NewReader(strconv.Itoa(chatID))); err != nil {
		return nil, err
	}

	// add caption
	fw, err = writer.CreateFormField("caption")
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, strings.NewReader(caption)); err != nil {
		return nil, err
	}

	// add photo
	fw, err = writer.CreateFormFile("photo", "problem")
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, bytes.NewBuffer(photoBytes)); err != nil {
		return nil, err
	}
	writer.Close()

	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	res, err := cl.Do(req)
	if err != nil {
		return nil, err
	}

	rawbytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	log.Printf("raw, %v", string(rawbytes))

	var resp MessagePhoto
	if err := json.Unmarshal(rawbytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response %v, err: <%v>", rawbytes, err)
	}

	return &resp, nil
}
