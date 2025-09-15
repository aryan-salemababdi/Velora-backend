package contact

import (
	dto "backend/src/modules/contact/dto"
	"fmt"
	"net/smtp"

	velora "github.com/aryan-salemababdi/Velora/app"
)

type ContactService struct {
	cfg SMTPConfig
}

func NewContactService() *ContactService {
	config := velora.LoadConfig("config.yaml")

	smtpCfg := SMTPConfig{
		From:     config.Get("SMTP_FROM"),
		Password: config.Get("SMTP_PASSWORD"),
		To:       config.Get("SMTP_TO"),
		Host:     config.Get("SMTP_HOST"),
		Port:     config.Get("SMTP_PORT"),
	}

	return &ContactService{cfg: smtpCfg}
}

func (s *ContactService) SendContactEmail(d dto.CreateContactDTO) error {
	c := s.cfg

	body := fmt.Sprintf(
		"From: %s\nSubject: %s\n\n%s\n\nSender: %s",
		d.FullName, d.Subject, d.Message, d.Email,
	)

	auth := smtp.PlainAuth("", c.From, c.Password, c.Host)

	return smtp.SendMail(c.Host+":"+c.Port, auth, c.From, []string{c.To}, []byte(body))
}
