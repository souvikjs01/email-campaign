package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "text email body.")
		// msg := []byte(formattedMsg)

		formattedMsg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("worker :%d error parsing template for %s", id, recipient.Email)
			// todo: add to queue
			continue
		}

		fmt.Printf("worker %d: sending email to %s\n", id, recipient.Email)
		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "souvik741156@gmail.com", []string{recipient.Email}, []byte(formattedMsg))
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(50 * time.Millisecond)
		fmt.Printf("worker %d: sent email to %s\n", id, recipient.Email)

	}
}
