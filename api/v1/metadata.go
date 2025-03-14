package v1

type Metadata struct {
	Draft bool `yaml:"draft" validate:"required"`

	Name string `yaml:"name" validate:"required"`

	// Name of js file to be generated. IF not set, than generated from name
	JsFile string `yaml:"jsFile" validate:"file"`

	Customer Address `yaml:"customerAddress"`

	DeliveryAddress Address `yaml:"deliveryAddress"`
}

type Address struct {
	Name         string `yaml:"name"`
	Country      string `yaml:"country"`
	CountryCode  string `yaml:"countryCode"`
	City         string `yaml:"city" `
	Street       string `yaml:"street" `
	StreetNo     string `yaml:"streetNo" `
	Zip          string `yaml:"zip"`
	PostalCode   string `yaml:"postalCode"`
	PhoneNumber1 string `yaml:"phoneNumber1"`
	PhoneNumber2 string `yaml:"phoneNumber2"`
	Email        string `yaml:"email"`
}
