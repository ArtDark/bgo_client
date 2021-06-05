package qr

import (
	"io"
	"log"
	"net/http"
	"os"
)

type apiMethod string
type Data string
type Size struct {
	Height int
	Weight int
}
type Api struct {
	Protocol string
	Dns      string
	Version  string
	Method   apiMethod
	Data     Data
	Size     Size
}

func NewApi(protocol string, dns string, version string, method apiMethod, data Data, size Size) *Api {
	return &Api{Protocol: protocol, Dns: dns, Version: version, Method: method, Data: data, Size: size}
}

func QrCreator(reqURL string, fileName string) (err error) {

	resp, err := http.Get(reqURL)

	if err != nil {
		log.Println(err)

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)

	}

	file, err := os.Create(fileName)
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
