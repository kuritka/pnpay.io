package main

import (
	"log"
	"pnpay.io/internal/js2pdf"
)

func main() {
	//t := v1.SinglePage{}
	//yamlData, err := os.ReadFile("./test/data/full.pnpayio.yaml")
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	//err = yaml.Unmarshal([]byte(yamlData), &t)
	//if err != nil {
	//	log.Fatalf("error: %v", err)
	//}
	convertor := js2pdf.NewYamlToJsImpl()
	err := convertor.Convert("./test/data/full.pnpayio.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
