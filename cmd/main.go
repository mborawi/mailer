package main

import (
	"fmt"
	"github.com/mborawi/mailer"
	"log"
)

func main() {
	host := "smtp.gmail.com"
	port := "587"
	from := "mborawi@gmail.com"
	to := []string{from}
	body := fmt.Sprintf("To: %s\r\nSubject: Why are you not using Mailtrap yet?\r\n\r\nHere’s the space for our great sales pitch\r\n", from)
	msg := []byte(body)
	username := "test"
	password := "1223"

	// SendEmail(host string, port string, from string, to []string, msg []byte)
	err := mailer.SendEmail(host, port, from, to, msg, username, password)
	if err != nil {
		log.Println(err)
	}
}
