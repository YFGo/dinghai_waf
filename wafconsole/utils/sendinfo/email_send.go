package sendinfo

import (
	"log/slog"
	"net/smtp"
	"wafconsole/utils/const/common"
)

func SendEmail(smtpServer, smtpPort, from, to, content string, auth smtp.Auth) error {
	subject := common.AppName
	msg := []byte("To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		content)
	err := smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		slog.Error("send email err:", err, "user: ", to)
		return err
	}
	return nil
}
