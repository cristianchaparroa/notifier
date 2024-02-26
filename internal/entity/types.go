package entity

import (
	"encoding/json"
	"io"
)

const (
	StatusNotificationType    = "status"
	NewsNotificationType      = "news"
	MarketingNotificationType = "marketing"
)

type Notification struct {
	Content   string `json:"content"`
	Type      string `json:"type"`
	Recipient string `json:"recipient"`
}

func NewNotification(content, typ, recipient string) *Notification {
	return &Notification{
		Content:   content,
		Type:      typ,
		Recipient: recipient,
	}
}

func NewNotificationFromRequest(r io.Reader) (*Notification, error) {
	// Create a decoder for the response body.
	decoder := json.NewDecoder(r)

	// Decode the JSON data into a struct.
	var n Notification
	decodeBody := decoder.Decode(&n)
	if decodeBody != nil {
		return nil, decodeBody
	}

	return &n, nil
}
