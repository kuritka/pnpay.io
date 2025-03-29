package main

import (
	"log"
	"time"

	v1 "pnpay.io/api/v1"

	"github.com/alecthomas/kong"
	"pnpay.io/internal/cli"
	"pnpay.io/internal/js2pdf"
	"pnpay.io/internal/logging"
)

var logger = logging.Logger()

func main() {
	var config = &cli.Config{}
	_ = kong.Parse(config)
	logging.Init(config)
	logger.Info().Msg("pnpay.io transition started")

	convertor := js2pdf.NewYamlToJsImpl()
	doc, err := convertor.Convert("./test/data/full.pnpayio.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	ch := make(chan struct{})
	go func(doc *v1.SinglePageExtended) {
		time.Sleep(1 * time.Second)
		logger.Info().Msg("generate pdf....")
		err = js2pdf.GeneratePDF("http://localhost:8090", "./test/output/output.pdf", doc)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		logger.Info().Msg("done....")
		ch <- struct{}{}
	}(doc)

	if config.IsLocalServerAllowed() {
		startWebServer(config)
	}
	<-ch
	close(ch)
}
