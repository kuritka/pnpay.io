package js2pdf

import (
	"bytes"
	"embed"
	v1 "pnpay.io/api/v1"
	"text/template"
)

const (
	cssFile string = "templates/file.css"
)

//go:embed templates/*.css
var content embed.FS

// path - output js file
func (y *YamlToJsImpl) convertSinglePageObjectToJsData(_ *v1.SinglePage) ([]byte, error) {
	file, err := css()
	return []byte(file), err
}

func css() (string, error) {
	const (
		sizeX = "15cm"
		sizeY = "11cm"
	)
	file, _ := content.ReadFile(cssFile)
	tmpl, err := template.New("css").Parse(string(file))
	if err != nil {
		return "", err
	}
	data := struct {
		SizeX string
		SizeY string
	}{
		SizeX: sizeX,
		SizeY: sizeY,
	}
	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

//func fillFromYaml[T any](a T, file string) error {
//	data, err := content.ReadFile(file)
//	if err != nil {
//		return err
//	}
//	err = file.Unmarshal(data, a)
//	return err
//}
