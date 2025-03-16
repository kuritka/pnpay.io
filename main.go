package main

import (
	"log"
	"pnpay.io/internal/cli"
	"pnpay.io/internal/logging"

	"github.com/alecthomas/kong"
	"pnpay.io/internal/js2pdf"
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

	if config.IsLocalServerAllowed() {
		startWebServer(config, err)
	}
}
