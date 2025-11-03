package main

import (
	"encoding/csv"
	"os"
)

func loadRecipient(path string, ch chan Recipient) error {
	f, err := os.Open(path)

	if err != nil {
		return err
	}

	defer f.Close()
	defer close(ch)

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		// fmt.Println(record)
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}

	return nil
}
