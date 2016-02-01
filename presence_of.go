package presenceof

import (
	"reflect"

	"github.com/gostrut/strut/invalid"
)

// Validator for presenceof validates the presence of a value on a field
// Zero values are considered invalid and will invalidate
// StructTag value is passed as the 2nd string arg, but we don't need it here
func Validator(n, _ string, v *reflect.Value) (invalid.Field, error) {
	z := reflect.Zero(v.Type())

	if reflect.DeepEqual(z.Interface(), v.Interface()) {
		return field{n}, nil
	}

	return nil, nil
}
