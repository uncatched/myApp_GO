package validate

import (
	"fmt"
	"reflect"
)

// Validator ...
type Validator struct {
}

const tagName = "validate"

// Validate ...
func (v *Validator) Validate(s interface{}) error {
	var valueOfS reflect.Value

	if isPointer(s) {
		valueOfS = reflect.Indirect(reflect.ValueOf(s))
	} else {
		valueOfS = reflect.ValueOf(s)
	}

	for i := 0; i < valueOfS.NumField(); i++ {
		tag := valueOfS.Type().Field(i).Tag.Get(tagName)
		if tag == "" || tag == "-" {
			continue
		}

		switch tag {
		case "required":
			if valueOfS.Field(i).IsZero() {
				fieldName := valueOfS.Type().Field(i).Name
				return &ValidationError{
					reason: fmt.Sprintf("`%s` is required field", fieldName),
				}
			}

			continue
		default:
			continue
		}
	}

	return nil
}

func isPointer(s interface{}) bool {
	return reflect.ValueOf(s).Kind() == reflect.Ptr
}
