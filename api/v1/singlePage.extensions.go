package v1

import (
	"strconv"
	"strings"
)

// SinglePageExtended extends SinglePage with additional with additional fields
type SinglePageExtended struct {
	SinglePage *SinglePage
	Extensions *Extensions
}

func NewSinglePageExtended(sp *SinglePage) (*SinglePageExtended, error) {
	extensions, err := newExtensions(sp)
	if err != nil {
		return nil, err
	}

	return &SinglePageExtended{
		SinglePage: sp,
		Extensions: extensions,
	}, nil
}

type ExtendedCanvas struct {
	Width float64 `yaml:"width" validate:"required"`

	Height float64 `yaml:"height" validate:"required"`
}

type Extensions struct {
	PDFCanvas ExtendedCanvas
}

func newExtensions(sp *SinglePage) (*Extensions, error) {
	extensions := &Extensions{
		PDFCanvas: ExtendedCanvas{},
	}
	// PDF Canvas
	var err error
	printWidth, err := strconv.ParseFloat(strings.Trim(sp.Spec.Canvas.Width, "cm"), 64)
	if err != nil {
		return nil, err
	}
	printHeight, err := strconv.ParseFloat(strings.Trim(sp.Spec.Canvas.Height, "cm"), 64)
	if err != nil {
		return nil, err
	}
	shadowOverlap, err := strconv.ParseFloat(strings.Trim(sp.Spec.Canvas.ShadowOverlap, "cm"), 64)
	if err != nil {
		return nil, err
	}
	extPadding, err := strconv.ParseFloat(strings.Trim(sp.Spec.Canvas.ExtPadding, "cm"), 64)
	if err != nil {
		return nil, err
	}
	// PDF document has one cm padding around printed document
	extensions.PDFCanvas.Width = printWidth + (shadowOverlap * 2) + (extPadding * 4)
	extensions.PDFCanvas.Height = printHeight + (shadowOverlap * 2) + (extPadding * 4)
	return extensions, err
}
