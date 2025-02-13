package openapi3

import (
	"context"

	"github.com/missmp/kin-openapi/jsoninfo"
)

// Discriminator is specified by OpenAPI/Swagger standard version 3.0.
type Discriminator struct {
	ExtensionProps
	PropertyName string            `json:"propertyName"`
	Mapping      map[string]string `json:"mapping,omitempty"`
}

func (value *Discriminator) MarshalJSON() ([]byte, error) {
	return jsoninfo.MarshalStrictStruct(value)
}

func (value *Discriminator) UnmarshalJSON(data []byte) error {
	return jsoninfo.UnmarshalStrictStruct(data, value)
}

func (value *Discriminator) Validate(c context.Context) error {
	return nil
}
