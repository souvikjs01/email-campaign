package main

import (
	"bytes"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	recipientChan := make(chan Recipient)
	go func() {
		loadRecipient("./list.csv", recipientChan)
	}()

	var wg sync.WaitGroup

	workerCount := 3

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChan, &wg)
	}

	wg.Wait()

}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, r)

	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
