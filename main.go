package main

import (
	"log"

	"pnpay.io/internal/js2pdf"
)

func main() {
	convertor := js2pdf.NewYamlToJsImpl()
	err := convertor.Convert("./test/data/full.pnpayio.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
