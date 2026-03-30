package sendto

import (
	"bytes"
	"fmt"
	"go-ecommerce/global"
	"html/template"
	"net/smtp"
	"strings"

	"go.uber.org/zap"
)

const (
	SMTPHOST     = "smtp.gmail.com"
	SMTPPORT     = 587
	SMTPUsername = "[EMAIL_ADDRESS]"
	SMTPPassword = ""
)

type EmailAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version:1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ", "))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += mail.Body
	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Name:    "test",
			Address: from,
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    fmt.Sprintf("Your OTP is: %s. Please enter it to verify your account.", otp),
	}

	messageMail := BuildMessage(contentEmail)

	// Send SMTP
	authention := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHOST)

	err := smtp.SendMail(SMTPHOST+":587", authention, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed :: ", zap.Error(err))
		return err
	}

	return nil
}

func SendTemplateEmailOtp(
	to []string,
	from string,
	nameTemplate string,
	dataTemplate map[string]interface{},
) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}

	return send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates/email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From: EmailAddress{
			Name:    "test",
			Address: from,
		},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)

	// Send SMTP
	authention := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHOST)

	err := smtp.SendMail(SMTPHOST+":587", authention, from, to, []byte(messageMail))
	if err != nil {
		global.Logger.Error("Email send failed :: ", zap.Error(err))
		return err
	}

	return nil
}
