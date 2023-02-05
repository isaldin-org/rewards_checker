package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func SendToTelegram(rewards string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	body := []byte(fmt.Sprintf(`{
		"chat_id": %s,
		"text": "%s"
	}`, allowedUserId, rewards))

	// Create a HTTP post request
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	log.Printf("Status code is %s", res.Status)
}
