package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/smtp"
	"os"

	"github.com/crud_golang/models"
)

func Send(destinatary string, student models.Student) (bool, error) {

	//from := "mail"
	//pass := "pass"
	from := os.Getenv("MAIL_SEND_FROM")
	pass := os.Getenv("MAIL_SEND_PASS")
	to := destinatary

	out, er := json.Marshal(student)
	if er != nil {
		return false, errors.New("Error convert body")
	}

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Thanks for register in our school " + student.Name + " \n\n" +
		string(out)

	log.Println(msg)
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	log.Println(err)

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false, errors.New("Error smtp")
	}

	return true, nil
}
