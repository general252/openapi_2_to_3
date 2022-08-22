package openapi_2_to_3

import (
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
)

// Convert openapi2 convert to openapi3
func Convert(swagger2Data []byte) (openApi3Data []byte, err error) {
	s2 := openapi2.T{}
	if err = s2.UnmarshalJSON(swagger2Data); err != nil {
		return nil, err
	}

	s3, err := openapi2conv.ToV3(&s2)
	if err != nil {
		return nil, err
	}

	result, err := s3.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return result, nil
}
