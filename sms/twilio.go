package sms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func Send(to []string, message string) {
	for _, e := range to {
		send(e, message)
	}
}

func send(to, message string) (string, error) {
	// Set account keys & information
	accountSid := "AC06dc812b689948eff7b484b7ad3a3548"
	authToken := "8ade518cea360d0e41547dd8529cf2ef"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	//to fetch number from config.yml
	// Pack up the data for our message
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", "15045798139")
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	// Create HTTP request client
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make HTTP POST request and return message SID
	resp, _ := client.Do(req)
	var err error
	var responseMessage string
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&data)
		if err == nil {
			responseMessage = fmt.Sprintf("%v", data["sid"])
		}
	} else {
		responseMessage = resp.Status
	}
	return responseMessage, err
}
