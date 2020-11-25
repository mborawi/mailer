package main

import (
	"fmt"
	"github.com/mborawi/mailer"
	"log"
	"net/smtp"
)

func main() {
	host := "smtp.gmail.com"
	port := "587"
	from := "123@gmail.com"
	to := []string{from}
	body := fmt.Sprintf("To: %s\r\nSubject: Why are you not using Mailtrap yet?\r\n\r\nHere’s the space for our great sales pitch\r\n", from)
	msg := []byte(body)
	username := "test"
	password := "1223"

	auth := mailer.UnencryptedAuth{
		smtp.PlainAuth(
			"",
			username,
			password,
			host,
		),
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	// SendEmail(host string, port string, from string, to []string, msg []byte)
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		log.Println(err)
	}
}
