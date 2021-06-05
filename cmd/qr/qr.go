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

	createQrUrl := qr.NewApi("https://", "api.qrserver.com", "v1", "create-qr-code", text, size)

	value := make(url.Values)
	value.Set("data", string(text))
	value.Set("size", fmt.Sprintf("%dX%d", createQrUrl.Size.Height, createQrUrl.Size.Weight))

	urlReq := fmt.Sprintf("%s%s/%s/%s/?%s",
		createQrUrl.Protocol,
		createQrUrl.Dns,
		createQrUrl.Version,
		createQrUrl.Method,
		value.Encode())

	err := qr.QrCreator(urlReq, fileName)
	if err != nil {
		log.Println(err)
		return
	}

}
