package v1

// SinglePageExtended extends SinglePage with additional with additional fields
type SinglePageExtended struct {
	SinglePage *SinglePage
	Extensions *Extensions
}

func NewSinglePageExtended(sp *SinglePage) *SinglePageExtended {
	return &SinglePageExtended{
		SinglePage: sp,
		Extensions: newExtensions(sp),
	}
}

type ExtendedCanvas struct {
	Width string `yaml:"width" validate:"required"`

	Height string `yaml:"height" validate:"required"`
}

type Extensions struct {
	Canvas ExtendedCanvas
}

func newExtensions(sp *SinglePage) *Extensions {
	return &Extensions{}
}
