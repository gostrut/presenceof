package presenceof

import "reflect"
import "github.com/gostrut/invalid"

// Validator for presenceof validates the presence of a value on a field
// Zero values are considered invalid and will invalidate
// StructTag value is passed as the 2nd string arg, but we don't need it here
func Validator(name, _ string, vo *reflect.Value) (invalid.Field, error) {
	vi := vo.Interface()
	z := reflect.Zero(vo.Type())
	if reflect.DeepEqual(z.Interface(), vi) {
		return iField{name}, nil
	}

	return nil, nil
}
