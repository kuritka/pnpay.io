package js2pdf

import (
	"bytes"
	"embed"
	"text/template"

	v1 "pnpay.io/api/v1"
)

//go:embed templates/*.css
//go:embed templates/*.html
//go:embed templates/*.html
//go:embed templates/*.js
var content embed.FS

// path - output js file
func (y *YamlToJsImpl) convertSinglePageObjectToJsData(doc *v1.SinglePageExtended) (css []byte, html []byte, js []byte, err error) {
	css, err = conversion(doc, doc.SinglePage.Spec.Input.Css.GetFilePath())
	if err != nil {
		return nil, nil, nil, err
	}
	html, err = conversion(doc, doc.SinglePage.Spec.Input.Html.GetFilePath())
	if err != nil {
		return nil, nil, nil, err
	}
	js, err = conversion(doc, doc.SinglePage.Spec.Input.Js.GetFilePath())
	if err != nil {
		return nil, nil, nil, err
	}
	return css, html, js, nil
}

func conversion(doc *v1.SinglePageExtended, fileName string) ([]byte, error) {
	file, _ := content.ReadFile(fileName)
	tmpl, err := template.New(fileName).Parse(string(file))
	if err != nil {
		return nil, err
	}
	var output bytes.Buffer
	err = tmpl.Execute(&output, doc)
	if err != nil {
		return nil, err
	}
	return output.Bytes(), nil
}
