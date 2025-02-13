package openapi3

import (
	"github.com/missmp/kin-openapi/jsoninfo"
)

// Example is specified by OpenAPI/Swagger 3.0 standard.
type Example struct {
	ExtensionProps

	Summary       string      `json:"summary,omitempty"`
	Description   string      `json:"description,omitempty"`
	Value         interface{} `json:"value,omitempty"`
	ExternalValue string      `json:"externalValue,omitempty"`
}

func NewExample(value interface{}) *Example {
	return &Example{
		Value: value,
	}
}

func (example *Example) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(example)
}

func (example *Example) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, example)
}
