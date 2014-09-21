package presenceof

import "fmt"

type iField struct {
	name string
}

func (f iField) Name() string {
	return f.name
}

func (f iField) Validator() string {
	return "PresenceOf"
}

func (f iField) Error() string {
	return fmt.Sprintf("%s is required", f.Name())
}
