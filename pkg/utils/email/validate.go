package email

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/smtp"
	"strings"
	"time"

	"github.com/domgolonka/foretoken/app"
)

var ErrUnresolvableHost = "UNRESOLVED_HOST"

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

func CatchAll(app *app.App, email string) (bool, error) {
	if app.Config.SMTP.Hostname != "" && app.Config.SMTP.MailAddress != "" {
		return true, catchAll(app.Config.SMTP.Hostname, app.Config.SMTP.MailAddress, email)
	}
	return false, nil
}

func catchAll(serverHostName, serverMailAddress, email string) error {
	_, domain := Split(email)
	return validateHostAndEmail(serverHostName, serverMailAddress, randSeq(10)+"@"+domain)
}

func validateHostAndEmail(serverHostName, serverMailAddress, email string) error {
	_, host := Split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return errors.New(ErrUnresolvableHost)
	}
	client, err := DialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), time.Second*1)
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
	err = client.Close()
	if err != nil {
		return err
	}
	return nil
}

// ValidateHost validate mail host.
func validateHost(email string) error {
	_, host := Split(email)
	mx, err := net.LookupMX(host)

	if err != nil {
		return errors.New(ErrUnresolvableHost)
	}

	client, err := DialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), time.Second*1)
	if err != nil {
		return NewSMTPError(err)
	}
	err = client.Close()
	if err != nil {
		return err
	}
	return nil
}

func DialTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	t := time.AfterFunc(timeout, func() {
		err := conn.Close()
		if err != nil {
			return
		}
	})
	defer t.Stop()

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func Split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))] //nolint
	}
	return "ee" + string(b)
}
