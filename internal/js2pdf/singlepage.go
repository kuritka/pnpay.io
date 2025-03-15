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

	dataCSS, dataHTML, dataJS, err := y.convertSinglePageObjectToJsData(t)
	if err != nil {
		return err
	}
	err = y.write(t.Spec.Output.Css.GetFilePath(), dataCSS)
	if err != nil {
		return err
	}
	err = y.write(t.Spec.Output.Html.GetFilePath(), dataHTML)
	if err != nil {
		return err
	}
	err = y.write(t.Spec.Output.Js.GetFilePath(), dataJS)
	if err != nil {
		return err
	}
	return nil
}

func (y *YamlToJsImpl) write(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer func() { _ = file.Close() }()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing file: %v", err)
	}
	// ensure file is written
	return file.Sync()
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
