package js2pdf

import (
	"fmt"
	"os"

	"github.com/go-yaml/yaml"
	v1 "pnpay.io/api/v1"
)

type SinglePageConvertor interface {
	// Convert ConvertYamlToJs converts yaml to js
	Convert(yamlPath string) error
}

type YamlToJsImpl struct {
}

func NewYamlToJsImpl() *YamlToJsImpl {
	return &YamlToJsImpl{}
}

func (y *YamlToJsImpl) Convert(yamlPath string) error {
	t, err := y.convertYamlToSinglePageObject(yamlPath)
	if err != nil {
		return err
	}

	data, err := y.convertSinglePageObjectToJsData(t)
	if err == nil {
		return err
	}
	file, err := os.Create(t.Metadata.JsFile)
	if err != nil {
		return fmt.Errorf("error creating .js file: %v", err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing .js file: %v", err)
	}
	return nil
}

// path to yaml file
func (y *YamlToJsImpl) convertYamlToSinglePageObject(path string) (*v1.SinglePage, error) {
	t := &v1.SinglePage{}
	yamlData, err := os.ReadFile(path)
	if err != nil {
		return t, fmt.Errorf("parsing error read file: %v", err)
	}
	err = yaml.Unmarshal(yamlData, &t)
	if err != nil {
		return t, fmt.Errorf("parsing error unmarhall: %v", err)
	}
	return t, nil
}
