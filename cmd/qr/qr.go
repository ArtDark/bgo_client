package main

import (
	"github.com/ArtDark/bgo_client/pkg/qr"
	"log"
)

func main() {

	//TODO: написать тайм-аут

	err := qr.QrCreator()
	if err != nil {
		log.Println(err)
	}

	log.Println("File created...")

}
