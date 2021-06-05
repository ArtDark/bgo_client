package main

import (
	"fmt"
	"github.com/ArtDark/bgo_client/pkg/qr"
	"log"
	"net/url"
	"os"
)

func main() {

	var text qr.Data = "https://netology.ru"
	size := qr.Size{Height: 100, Weight: 100}
	fileName := "qr.png"
	timeOut, ok := os.LookupEnv("qr_timeout")
	if !ok {
		log.Println("no timeout specified (qr_timeout env variable)")
	}

	log.Println(timeOut)

	srv := qr.NewService("https://", "Service.qrserver.com", "v1", "create-qr-code", text, size)

	value := make(url.Values)
	value.Set("data", string(text))
	value.Set("size", fmt.Sprintf("%dX%d", srv.Size.Height, srv.Size.Weight))

	urlReq := fmt.Sprintf("%s%s/%s/%s/?%s",
		srv.Protocol,
		srv.Dns,
		srv.Version,
		srv.Method,
		value.Encode())

	err := srv.QrCreator(urlReq, fileName)
	if err != nil {
		log.Println(err)
		return
	}

}
