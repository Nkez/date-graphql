package repository

import (
	"context"
	"crypto/tls"
	"github.com/go-mail/mail"
)

type SMTPEmailRepository struct {
	from     string
	password string
	host     string
	port     int
	dialer   *mail.Dialer
}

func NewSMTPEmailRepository(email, password, host string, port int) *SMTPEmailRepository {
	d := mail.NewDialer(host, port, email, password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true} // nolint: gosec
	return &SMTPEmailRepository{
		from:     email,
		password: password,
		host:     host,
		port:     port,
		dialer:   d,
	}
}

func (r *SMTPEmailRepository) Send(_ context.Context, email string) error {
	email = "kezmikita@gmail.com"
	m := mail.NewMessage()
	m.SetHeader("From", r.from)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "You create user gratz!!!!!!!!")
	m.SetBody("text/plain", "Тут съёмка видео запрещена, блядь. Камеру вырубай нахуй. Камеру вырубай блядь, мудила блядь")
	if err := r.dialer.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
