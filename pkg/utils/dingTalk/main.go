package dingTalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Content struct {
	Content string `json:"content"`
}

type Payload struct {
	MsgType string  `json:"msgtype"`
	Text    Content `json:"text"`
}

func NewDingTalkMsg(content string) *Payload {
	return &Payload{
		MsgType: "text",
		Text: Content{
			Content: content,
		},
	}
}

func SendDingDing(payload *Payload, accessToken, secret string) error {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	signed, err := urlSign(timestamp, secret)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%v&sign=%v&timestamp=%v", accessToken, signed, timestamp), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

func urlSign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
