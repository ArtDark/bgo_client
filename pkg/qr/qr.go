package qr

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
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

func (s *Service) QrCreator(data []byte, fileName string) (err error) {

	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Cannot open file", err)
		return err
	}
	defer func(c io.Closer) {
		if cerr := c.Close(); cerr != nil {
			log.Println("Cannot close file", cerr)
			return
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		log.Println("Cannot open file", err)
		return err
	}
	return nil

}

func (s *Service) Encode(ctx context.Context, data string) (dataexport []byte, filetype string, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, data, nil)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, "", err

	}
	dataexport, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	imgName := strings.Replace(resp.Header["Content-Type"][0], "/", ".", 1)

	return dataexport, imgName, nil
}
