package proxy

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Event string
}

type Params struct {
	Account  string
	Messages []Message
}

func Request(event []byte) error {
	url := os.Getenv("PROXY_HOST_URL")
	if len(url) < 1 {
		return errors.New("BadRequest: PROXY_HOST_URL not set")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(event))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("Error From Upstream proxy: " + resp.Status) // TODO: Improve string concat
	}
	return nil
}

func Process(event json.RawMessage, messages []Message) error {
	for _, message := range messages {
		log.Println(message)
		if message.Event == "unsubscribe" {
			err := Request(event)
			return err
		}
	}
	return nil
}
