package js2pdf

import (
	"bytes"
	"embed"
	"text/template"

	v1 "pnpay.io/api/v1"
)

const (
	cssFile  string = "templates/single_page.template.css"
	htmlFile string = "templates/single_page.template.html"
	jsFile   string = "templates/single_page.template.js"
)

//go:embed templates/*.css
//go:embed templates/*.html
//go:embed templates/*.html
//go:embed templates/*.js
var content embed.FS

// path - output js file
func (y *YamlToJsImpl) convertSinglePageObjectToJsData(sp *v1.SinglePage) (css []byte, html []byte, js []byte, err error) {
	css, err = conversion(sp, sp.Spec.Input.Css.GetFilePath())
	if err != nil {
		return nil, nil, nil, err
	}
	html, err = conversion(sp, sp.Spec.Input.Html.GetFilePath())
	if err != nil {
		return nil, nil, nil, err
	}
	js, err = conversion(sp, sp.Spec.Input.Js.GetFilePath())
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
