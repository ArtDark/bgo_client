package qr

import (
	"io"
	"log"
	"net/http"
	"os"
)

type QrApi struct {
	url     string
	version string
	method  string
	data    string
} //TODO: Написать структуру запроса

func QrCreator() error { //TODO: превратить в сервис
	reqURL := "http://api.qrserver.com/v1/create-qr-code/?data=Ты%20пидор!&size=100x100"

	resp, err := http.Get(reqURL)

	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}

	file, err := os.Create("image.png")
	if err != nil {
		log.Println("Cannot open file", err)
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println("Cannot close file", cerr)
		}
	}(file)

	_, err = file.Write(body)
	if err != nil {
		log.Println("Cannot open file", err)
	}

	return nil

}
