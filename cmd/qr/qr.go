package main

import (
	"context"
	"fmt"
	"github.com/ArtDark/bgo_client/pkg/qr"
	"log"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {

	var text qr.Data = "https://netology.ru"
	size := qr.Size{Height: 100, Weight: 100}
	timeoutEnv, ok := os.LookupEnv("qr_timeout")
	if !ok {
		log.Println("no timeout specified (qr_timeout env variable)")
	}

	timeout, err := strconv.Atoi(timeoutEnv)
	if err != nil {
		log.Println(err)
	}

	srv := qr.NewService("https://", "api.qrserver.com", "v1", "create-qr-code", text, size, time.Duration(timeout)*time.Millisecond)

	value := make(url.Values)
	value.Set("data", string(text))
	value.Set("size", fmt.Sprintf("%dX%d", srv.Size.Height, srv.Size.Weight))

	urlReq := fmt.Sprintf("%s%s/%s/%s/?%s",
		srv.Protocol,
		srv.Dns,
		srv.Version,
		srv.Method,
		value.Encode())

	ctx, _ := context.WithTimeout(context.Background(), srv.Timeout)

	qrCode, img, err := srv.Encode(ctx, urlReq)
	if err != nil {
		log.Println(err)
	}

	err = srv.QrCreator(qrCode, img)
	if err != nil {
		log.Println(err)
	}

}
