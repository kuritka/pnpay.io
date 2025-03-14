package js2pdf

import (
	"bytes"
	"embed"
	v1 "pnpay.io/api/v1"
	"text/template"
)

const (
	cssFile  string = "templates/file.css"
	htmlFile string = "templates/file.html"
	jsFile   string = "templates/file.js"
)

//go:embed templates/*.css
//go:embed templates/*.html
//go:embed templates/*.html
//go:embed templates/*.js
var content embed.FS

// path - output js file
func (y *YamlToJsImpl) convertSinglePageObjectToJsData(sp *v1.SinglePage) (css []byte, html []byte, js []byte, err error) {
	css, err = conversion(sp, cssFile)
	if err != nil {
		return nil, nil, nil, err
	}
	html, err = conversion(sp, htmlFile)
	if err != nil {
		return nil, nil, nil, err
	}
	js, err = conversion(sp, jsFile)
	if err != nil {
		return nil, nil, nil, err
	}
	return css, html, js, nil
}

func conversion(sp *v1.SinglePage, fileName string) ([]byte, error) {
	file, _ := content.ReadFile(fileName)
	tmpl, err := template.New(fileName).Parse(string(file))
	if err != nil {
		return nil, err
	}
	var output bytes.Buffer
	err = tmpl.Execute(&output, sp)
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}

//func fillFromYaml[T any](a T, file string) error {
//	data, err := content.ReadFile(file)
//	if err != nil {
//		return err
//	}
//	err = file.Unmarshal(data, a)
//	return err
//}
