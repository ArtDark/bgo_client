package qr

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type apiMethod string
type Data string
type Size struct {
	Height int
	Weight int
}
type Service struct {
	Protocol string
	Dns      string
	Version  string
	Method   apiMethod
	Data     Data
	Size     Size
	Timeout  time.Duration
}

func NewService(protocol string, dns string, version string, method apiMethod, data Data, size Size, timeout time.Duration) *Service {
	return &Service{
		Protocol: protocol,
		Dns:      dns,
		Version:  version,
		Method:   method,
		Data:     data,
		Size:     size,
		Timeout:  timeout}
}

func (s *Service) QrCreator(reqURL string, fileName string) (err error) {

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
