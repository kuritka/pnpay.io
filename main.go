package main

import (
	"log"
	"net/http"
	"pnpay.io/internal/cli"

	"github.com/alecthomas/kong"
	"pnpay.io/internal/js2pdf"
)

func main() {
	var config = &cli.Config{}
	_ = kong.Parse(config)

	convertor := js2pdf.NewYamlToJsImpl()
	err := convertor.Convert("./test/data/full.pnpayio.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if config.IsLocalServerAllowed() {
		fs := http.FileServer(http.Dir(config.LocalServerPath))
		http.Handle("/", fs)

		log.Printf("Server runs on http://localhost:%s/\n", config.LocaLServerPort)
		err = http.ListenAndServe(":"+config.LocaLServerPort, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
