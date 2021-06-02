package main

import (
	"fmt"
	"github.com/ArtDark/bgo_client/pkg/qr"
	"net/url"
)

func main() {

	var text qr.Data = "Netology"

	api := &qr.Api{
		Protocol: "http://",
		Dns:      "api.qrserver.com",
		Version:  "v1",
		Method:   "create-qr-code",
		Data:     text,
		Size: qr.Size{
			100,
			100,
		},
	} //TODO: написать тайм-аут

	value := make(url.Values)
	value.Set("data", string(text))

	urlReq := fmt.Sprintf("%s%s/%s/%s",
		api.Protocol,
		api.Dns,
		api.Version,
		value.Encode())
	println(urlReq)

}
