package email

import (
	"crypto/tls"
	"net/smtp"

	"github.com/jordan-wright/email"
)

// 发送邮件
func SendEmail(address string, subject string, content string) error {
	e := &email.Email{
		ReplyTo: []string{},
		From:    "Miracle <amyunfei@163.com>",
		To:      []string{address},
		Subject: subject,
		HTML:    []byte(content),
	}
	err := e.SendWithTLS(
		"smtp.163.com:465",
		smtp.PlainAuth("", "amyunfei@163.com", "TZJPWJLGJVRXFDHI", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"},
	)
	return err
}
