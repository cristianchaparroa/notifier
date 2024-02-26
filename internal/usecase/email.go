package usecase

import "fmt"

type emailManager struct {
}

func NewEmailManager() EmailManager {
	return &emailManager{}
}

func (m *emailManager) Send(subject, addressee, body string) error {
	fmt.Printf("--> Mock Email Manager - subject %s, Receipent:%s, content:%s\n", subject, addressee, body)
	return nil
}
