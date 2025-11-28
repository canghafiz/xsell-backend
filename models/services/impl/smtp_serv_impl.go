package impl

import (
	"be/models/domains"
	"crypto/tls"
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

type SmtpServImpl struct {
	Smtp    domains.Smtp
	AppName string
}

func NewSmtpServImpl(smtp domains.Smtp, appName string) *SmtpServImpl {
	return &SmtpServImpl{Smtp: smtp, AppName: appName}
}

func (serv *SmtpServImpl) SendEmailOtp(email string, code string) error {
	subject := fmt.Sprintf("Verification code %s - ", serv.AppName)
	htmlBody := fmt.Sprintf(`<html><body>
<p>Hello,</p>
<p>Use this OTP: <strong>%s</strong></p>
<p>Code valid for 5 minutes.</p>
<p><small>Do not share this code with anyone.</small></p>
</body></html>`, code)

	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		serv.Smtp.From, email, subject, htmlBody,
	))

	// SMTP configuration
	addr := fmt.Sprintf("%s:%s", serv.Smtp.Host, serv.Smtp.Port)

	// Connect to SMTP server
	client, err := smtp.Dial(addr)
	if err != nil {
		log.Printf("[SMTP] Dial error: %v", err)
		return serv.handleSmtpError(err, "connection")
	}
	defer client.Close()

	// Send EHLO
	if err = client.Hello(serv.Smtp.Host); err != nil {
		log.Printf("[SMTP] Hello error: %v", err)
		return serv.handleSmtpError(err, "hello")
	}

	// Start TLS (untuk port 587)
	if ok, _ := client.Extension("STARTTLS"); ok {
		tlsConfig := &tls.Config{
			ServerName:         serv.Smtp.Host,
			InsecureSkipVerify: false,
		}
		if err = client.StartTLS(tlsConfig); err != nil {
			log.Printf("[SMTP] StartTLS error: %v", err)
			return serv.handleSmtpError(err, "tls")
		}
	}

	// Authenticate
	if serv.Smtp.Username != "" && serv.Smtp.Password != "" {
		auth := smtp.PlainAuth("", serv.Smtp.Username, serv.Smtp.Password, serv.Smtp.Host)
		if err = client.Auth(auth); err != nil {
			log.Printf("[SMTP] Auth error: %v", err)
			return serv.handleSmtpError(err, "auth")
		}
	}

	// Set sender
	if err = client.Mail(serv.Smtp.From); err != nil {
		log.Printf("[SMTP] Mail from error: %v", err)
		return serv.handleSmtpError(err, "sender")
	}

	// Set recipient
	if err = client.Rcpt(email); err != nil {
		log.Printf("[SMTP] Rcpt to error: %v", err)
		return serv.handleSmtpError(err, "recipient")
	}

	// Send email body
	writer, err := client.Data()
	if err != nil {
		log.Printf("[SMTP] Data error: %v", err)
		return serv.handleSmtpError(err, "data")
	}
	defer writer.Close()

	if _, err = writer.Write(msg); err != nil {
		log.Printf("[SMTP] Write error: %v", err)
		return serv.handleSmtpError(err, "write")
	}

	log.Printf("[SMTP] Email sent successfully to: %s", email)
	return nil
}

func (serv *SmtpServImpl) handleSmtpError(err error, stage string) error {
	errStr := err.Error()

	// Network/Connection errors
	if stage == "connection" {
		if strings.Contains(errStr, "connection refused") {
			return fmt.Errorf("mail server unavailable, please check SMTP_HOST and SMTP_PORT configuration")
		}
		if strings.Contains(errStr, "timeout") || strings.Contains(errStr, "deadline exceeded") {
			return fmt.Errorf("connection timeout, mail server not responding")
		}
		if strings.Contains(errStr, "no such host") {
			return fmt.Errorf("invalid SMTP host: %s", serv.Smtp.Host)
		}
		return fmt.Errorf("failed to connect to mail server: %v", err)
	}

	// TLS errors
	if stage == "tls" {
		if strings.Contains(errStr, "certificate") {
			return fmt.Errorf("TLS certificate error, please verify SMTP server configuration")
		}
		if strings.Contains(errStr, "handshake") {
			return fmt.Errorf("TLS handshake failed, SMTP server may not support STARTTLS on port %s", serv.Smtp.Port)
		}
		return fmt.Errorf("TLS connection failed: %v", err)
	}

	// Authentication errors
	if stage == "auth" {
		if strings.Contains(errStr, "535") {
			return fmt.Errorf("authentication failed: invalid SMTP username or password")
		}
		if strings.Contains(errStr, "530") {
			return fmt.Errorf("authentication required: SMTP server requires valid credentials")
		}
		if strings.Contains(errStr, "534") {
			return fmt.Errorf("authentication mechanism not supported by SMTP server")
		}
		return fmt.Errorf("authentication failed: %v", err)
	}

	// Sender/Recipient errors
	if stage == "sender" {
		if strings.Contains(errStr, "550") || strings.Contains(errStr, "553") {
			return fmt.Errorf("sender address rejected: %s", serv.Smtp.From)
		}
		return fmt.Errorf("failed to set sender: %v", err)
	}

	if stage == "recipient" {
		if strings.Contains(errStr, "550") {
			return fmt.Errorf("recipient address rejected or does not exist")
		}
		if strings.Contains(errStr, "551") {
			return fmt.Errorf("recipient not local, relay denied")
		}
		if strings.Contains(errStr, "552") {
			return fmt.Errorf("mailbox full or quota exceeded")
		}
		if strings.Contains(errStr, "554") {
			return fmt.Errorf("transaction failed, recipient address may be invalid")
		}
		return fmt.Errorf("failed to set recipient: %v", err)
	}

	// Data transfer errors
	if stage == "data" || stage == "write" {
		if strings.Contains(errStr, "552") {
			return fmt.Errorf("message size exceeds limit")
		}
		if strings.Contains(errStr, "554") {
			return fmt.Errorf("message rejected by server")
		}
		return fmt.Errorf("failed to send email data: %v", err)
	}

	// Generic SMTP errors
	if stage == "hello" {
		return fmt.Errorf("SMTP handshake failed: %v", err)
	}

	// Default
	return fmt.Errorf("smtp error at %s stage: %v", stage, err)
}
