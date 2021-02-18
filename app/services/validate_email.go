package services

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/domgolonka/threatdefender/app"
)

type SMTPError struct {
	Err error
}

func (e SMTPError) Error() string {
	return e.Err.Error()
}

func (e SMTPError) Code() string {
	return e.Err.Error()[0:3]
}

func NewSMTPError(err error) SMTPError {
	return SMTPError{
		Err: err,
	}
}

func ValidateEmail(app *app.App, email string) error {
	if app.Config.SMTP.Hostname != "" && app.Config.SMTP.MailAddress != "" {
		return validateHostAndEmail(app.Config.SMTP.Hostname, app.Config.SMTP.MailAddress, email)
	}
	return validateHost(email)
}

func validateHostAndEmail(serverHostName, serverMailAddress, email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return errors.New(ErrUnresolvableHost)
	}
	client, err := DialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), time.Second*5)
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Hello(serverHostName)
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Mail(serverMailAddress)
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Rcpt(email)
	if err != nil {
		return NewSMTPError(err)
	}
	defer client.Close()
	return nil
}

// ValidateHost validate mail host.
func validateHost(email string) error {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return errors.New(ErrUnresolvableHost)
	}
	client, err := DialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), time.Second*5)
	if err != nil {
		return NewSMTPError(err)
	}
	client.Close()
	return nil
}

func DialTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}
