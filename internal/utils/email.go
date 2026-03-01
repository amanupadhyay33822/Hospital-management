package utils

import (
	"fmt"
	"time"
)

type EmailData struct {
	To      string
	Subject string
	Body    string
}

func SendEmailAsync(email EmailData) {
	go func() {
		time.Sleep(2 * time.Second) // Simulate email sending delay
		fmt.Printf("Email sent to %s: %s\n", email.To, email.Subject)
	}()
}
