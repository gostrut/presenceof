package presenceof

import "fmt"

type field struct {
	name string
}

func (f field) Name() string {
	return f.name
}

func (f field) Validator() string {
	return "PresenceOf"
}

func (f field) Error() string {
	return fmt.Sprintf("%s is required", f.Name())
}
