package entity

type Notification struct {
	Content   string
	Type      string
	Recipient string
}

func NewNotification(content, typ, recipient string) *Notification {
	return &Notification{
		Content:   content,
		Type:      typ,
		Recipient: recipient,
	}
}
