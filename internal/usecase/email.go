package usecase

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"notifier/config"
)

type noopManager struct{}

func (m *noopManager) Send(subject, addressee, body string) error {
	fmt.Printf("--> Mock Email Manager - subject %s, Receipent:%s, content:%s\n", subject, addressee, body)
	return nil
}

type emailManager struct {
	cfg *config.Config
}

func NewEmailManager(cfg *config.Config) EmailManager {
	if config.EnvProduction == cfg.App.Env {
		return &emailManager{
			cfg: cfg,
		}
	}
	return &noopManager{}
}

func (m *emailManager) Send(subject, addressee, body string) error {
	fmt.Printf("--> Real Email Manager - subject %s, Receipent:%s, content:%s\n", subject, addressee, body)
	fmt.Println(m.cfg.SMTP)
	message := gomail.NewMessage()
	message.SetHeader("From", m.cfg.SMTP.User)
	message.SetHeader("To", addressee)
	message.SetHeader("Subject", subject)
	message.SetBody("text", body)

	d := gomail.NewDialer(m.cfg.SMTP.Server, m.cfg.SMTP.Port, m.cfg.SMTP.User, m.cfg.SMTP.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(message); err != nil {
		return err
	}

	return nil
	return nil
}
