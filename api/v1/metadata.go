package v1

type Metadata struct {
	Draft bool `yaml:"draft" validate:"required"`

	Name string `yaml:"name" validate:"required"`
}
