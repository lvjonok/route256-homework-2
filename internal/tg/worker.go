package tg

import (
	"log"
	"time"
)

type Client struct {
	baseUrl      string
	apiKey       string
	shutdownChan chan bool
	lastUpdateID int
}

func New(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseUrl: "https://api.telegram.org/bot" + apiKey + "/",
	}
}

func (c *Client) Close() {
	c.shutdownChan <- true
}

func (c *Client) GetUpdatesChan() <-chan *Update {
	updatesChan := make(chan *Update)
	go func() {
		for {
			select {
			case <-c.shutdownChan:
				close(updatesChan)
				return
			default:
			}

			updates, err := c.getUpdates()
			if err != nil {
				log.Printf("failed to get updates, %v", err)
				log.Printf("retry in 5 secs...")
				time.Sleep(5 * time.Second)
			}

			for _, update := range updates.Result {
				if update.UpdateID >= c.lastUpdateID {
					c.lastUpdateID = update.UpdateID + 1
				}
				updatesChan <- &update
			}
		}
	}()

	return updatesChan
}
