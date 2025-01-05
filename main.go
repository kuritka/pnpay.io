package main

import (
	"log"
	"os"

	"github.com/go-yaml/yaml"
	v1 "pnpay.io/api/v1"
)

func main() {
	t := v1.SinglePage{}
	yamlData, err := os.ReadFile("./test/data/full.pnpayio.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(yamlData), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

}
