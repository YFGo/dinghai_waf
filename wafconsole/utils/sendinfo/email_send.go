package sendinfo

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"wafconsole/utils/const/common"
)

func SendEmail(smtpServer, smtpPort, from, to, verificationCode string, auth smtp.Auth) error {
	serverAddr := fmt.Sprintf("%s:%s", smtpServer, smtpPort)
	tlsConfig := &tls.Config{
		ServerName:         smtpServer,
		InsecureSkipVerify: true,
	}
	// 建立 tls 链接
	conn, err := tls.Dial("tcp", serverAddr, tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()
	// 创建 SMTP 客户端
	client, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		return err
	}
	defer client.Close()
	// 身份认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("authentication failed: %v", err)
	}

	// 设置发件人
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("set sender failed: %v", err)
	}

	// 解析收件人（支持多个，用逗号分隔）
	recipients := strings.Split(to, ",")
	for _, r := range recipients {
		if err = client.Rcpt(r); err != nil {
			return fmt.Errorf("add recipient %s failed: %v", r, err)
		}
	}

	subject := common.AppName
	// 构造包含验证码的邮件内容
	msg := []byte("To: " + to + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		verificationCode)

	// 获取数据写入器
	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to start data transfer: %v", err)
	}
	// 写入邮件内容
	_, err = writer.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to write message: %v", err)
	}
	// 关闭写入器
	err = writer.Close()
	if err != nil {
		return fmt.Errorf("failed to close data transfer: %v", err)
	}

	return nil
}
