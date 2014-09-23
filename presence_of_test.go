package presenceof

import "fmt"
import "testing"
import "github.com/nowk/assert"
import "github.com/gostrut/strut"

func TestPresenceOf(t *testing.T) {
	val := strut.NewValidator()
	val.Checks("presence_of", Validator)

	type Address struct {
		Street string `presence_of:"true"`
	}

	type Person struct {
		Address Address `presence_of:"true"`
	}

	a := Address{}
	b := Address{Street: ""}
	c := Person{}
	d := Person{Address: Address{}} // NOTE "technically" should be valid

	for _, v := range []struct {
		f string
		i interface{}
	}{
		{"Street", a},
		{"Street", b},
		{"Address", c},
		{"Address", d},
	} {
		fields, err := val.Validates(v.i)
		f := fields.Get(v.f)[0]
		assert.Nil(t, err)
		assert.False(t, fields.Valid())
		assert.Equal(t, fmt.Sprintf("%s is required", v.f), f.Error())
	}

	e := Address{
		Street: "Sesame",
	}

	fields, err := val.Validates(e)
	assert.Nil(t, err)
	assert.True(t, fields.Valid())
}
