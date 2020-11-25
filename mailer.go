package mailer

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
)

type UnencryptedAuth struct {
	smtp.Auth
}

func (a UnencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}
