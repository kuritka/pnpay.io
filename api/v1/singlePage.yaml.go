// https://github.com/go-playground/validator
package v1

import (
	"fmt"
	"strings"
)

type LineType string

const (
	LineTypeSolid  LineType = "solid"
	LineTypeDotted LineType = "dotted"
	LineTypeDashed LineType = "dashed"
)

type FilePath struct {
	Dir      string `yaml:"dir" validate:"required"`
	Filename string `yaml:"filename" validate:"required"`
}

func (f FilePath) GetFilePath() string {
	return fmt.Sprintf("%s/%s", strings.TrimSuffix(f.Dir, "/"), f.Filename)
}

type Output struct {
	Css  FilePath `yaml:"css" validate:"required"`
	Html FilePath `yaml:"html" validate:"required"`
	Pdf  FilePath `yaml:"pdf" validate:"required"`
	Js   FilePath `yaml:"js" validate:"required"`
}

type Input struct {
	Css  FilePath `yaml:"css" validate:"required"`
	Html FilePath `yaml:"html" validate:"required"`
	Js   FilePath `yaml:"js" validate:"required"`
}

type Canvas struct {
	Width string `yaml:"width" validate:"required"`

	Height string `yaml:"height" validate:"required"`

	// ColorOverlap: Spadavka https://www.tonerpartner.cz/clanky/orezove-znacky-proc-jsou-pro-tisk-dulezite-a-jak-je-vytvorit-53251cz39332/
	ColorOverlap string `yaml:"overlap" validate:"required"`
}

type Generic struct {
	Type LineType `yaml:"type"`

	Size int `yaml:"size"`

	Color string `yaml:"color"`
}

type Corner struct {
	Generic Generic

	Enabled bool `yaml:"enabled"`
}

type TopLeftCross struct {
	Corner Corner
}

type TopRightCross struct {
	Corner Corner
}

type BottomLeftCross struct {
	Corner Corner
}

type BottomRightCross struct {
	Corner Corner
}

type Image struct {
	Path string `yaml:"path" validate:"required"`
}

type Background struct {
	Color string `yaml:"color" validate:"required"`
	Image Image  `yaml:"image"`
}

type Marks struct {
	Border string `yaml:"border"`

	Generic Generic `yaml:"generic"`

	TopLeftCross TopLeftCross `yaml:"topLeftCross"`

	TopRightCross TopRightCross `yaml:"topRightCross"`

	BottomLeftCross BottomLeftCross `yaml:"bottomLeftCross"`

	BottomRightCross BottomRightCross `yaml:"bottomRightCross"`
}

type LayerType string

const (
	LayerTypeImage      LayerType = "Image"
	LayerType3dBarCode  LayerType = "3DBarCode"
	LayerTypeText       LayerType = "Text"
	LayerTypeBqckground LayerType = "Background"
)

type Layer struct {
	Name string `yaml:"name" validate:"required" json:"name,omitempty"`

	Kind LayerType `yaml:"kind" validate:"required" json:"kind,omitempty"`

	Value string `yaml:"value" validate:"required" json:"value,omitempty"`

	Class string `yaml:"class" json:"class,omitempty"`
}

type SinglePageSpec struct {
	Output Output `yaml:"output" validate:"required"`

	Input Input `yaml:"input" validate:"required"`

	Canvas Canvas `yaml:"canvas" validate:"required"`

	Marks Marks `yaml:"marks" validate:"required"`

	Background Background `yaml:"background" validate:"required"`

	Layers []Layer `yaml:"layers"`
}

type SinglePage struct {
	Version string `yaml:"version" validate:"required"`

	// Metadata
	Metadata Metadata `yaml:"metadata" validate:"required"`

	Kind string `yaml:"kind" validate:"required"`

	// Spec
	Spec SinglePageSpec `yaml:"spec" validate:"required"`
}
