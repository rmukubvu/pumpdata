package sms

import (
	"encoding/json"
	"fmt"
	"github.com/rmukubvu/pumpdata/model"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Account struct {
	accountSid string
	authToken  string
	from       string
	to         string
	url        string
	client     *http.Client
	req        *http.Request
}

type TwilioConfig struct {
	Sid, Token, From string
}

func New(c TwilioConfig) *Account {
	return &Account{
		accountSid: c.Sid,
		authToken:  c.Token,
		from:       c.From,
		client:     &http.Client{},
		url:        "https://api.twilio.com/2010-04-01/Accounts/" + c.Sid + "/Messages.json",
	}
}

func (t *Account) buildMessageReader(to, message string) strings.Reader {
	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", t.from)
	msgData.Set("Body", message)
	return *strings.NewReader(msgData.Encode())
}

func (t *Account) Send(to string, message string) model.PumpLogs {
	return t.send(to, message)
}

func (t *Account) SendToMultiple(to []model.DigitalMessage) []model.PumpLogs {
	res := make([]model.PumpLogs, 0)
	for _, e := range to {
		res = append(res, t.send(e.Receiver, e.Message))
	}
	return res
}

func (t *Account) SendToMultipleNumbers(to []string, message string) []model.PumpLogs {
	res := make([]model.PumpLogs, 0)
	for _, e := range to {
		res = append(res, t.send(e, message))
	}
	return res
}

func (t *Account) send(to, message string) model.PumpLogs {
	reader := t.buildMessageReader(to, message)
	t.to = to
	t.req, _ = http.NewRequest("POST", t.url, &reader)
	t.req.SetBasicAuth(t.accountSid, t.authToken)
	t.req.Header.Add("Accept", "application/json")
	t.req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return t.post()
}

func (t *Account) post() model.PumpLogs {
	resp, _ := t.client.Do(t.req)
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
	return model.PumpLogs{
		Payload: responseMessage,
		Error:   resp.Status,
		Date:    time.Now().String(),
		Sender:  t.to,
	}
}
