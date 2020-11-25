package mailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

func SendEmail(host string, port string, from string, to []string, msg []byte, un string, pwd string) error {
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}
	servername := fmt.Sprintf("%s:%s", host, port)

	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		return err
	}
	auth := smtp.PlainAuth("", un, pwd, host)

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	defer c.Close()
	if err = c.Auth(auth); err != nil {
		return err
	}
	// if err = c.Hello(); err != nil {
	// 	return err
	// }
	if err = c.Mail(from); err != nil {
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
