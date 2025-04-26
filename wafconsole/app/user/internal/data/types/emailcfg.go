package types

import "net/smtp"

type EmailCfg struct {
	SmtpUserName string
	SmtpPassword string
	SmtpServer   string
	SmtpProt     string
	Auth         smtp.Auth
}
