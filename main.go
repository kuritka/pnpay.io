package main

import (
	"log"
	"time"

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
	err := convertor.Convert("./test/data/full.pnpayio.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		logger.Info().Msg("generate pdf....")
		err = js2pdf.GeneratePDF("http://localhost:8090", "./test/output/output.pdf", 16, 12)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		logger.Info().Msg("done....")
		ch <- struct{}{}
	}()

	if config.IsLocalServerAllowed() {
		startWebServer(config)
	}
	<-ch
	close(ch)
}
