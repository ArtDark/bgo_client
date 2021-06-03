package main

import (
	"fmt"
	"github.com/ArtDark/bgo_client/pkg/qr"
	"log"
	"net/url"
)

func main() {

	var text qr.Data = "Netology"
	height := 100
	weight := 100

	createQrApi := &qr.Api{
		Protocol: "http://",
		Dns:      "api.qrserver.com",
		Version:  "v1",
		Method:   "create-qr-code",
		Data:     text,
		Size: qr.Size{
			height,
			weight,
		},
	}

	value := make(url.Values)
	value.Set("data", string(text))
	value.Set("size", fmt.Sprintf("%dX%d", createQrApi.Size.Height, createQrApi.Size.Weight))

	urlReq := fmt.Sprintf("%s%s/%s/%s/%s",
		createQrApi.Protocol,
		createQrApi.Dns,
		createQrApi.Version,
		createQrApi.Method,
		value.Encode())

	err := qr.QrCreator(urlReq)
	if err != nil {
		log.Println(err)
		return
	}

}
