# presenceof

[![Build Status](https://travis-ci.org/gostrut/presenceof.svg?branch=master)](https://travis-ci.org/gostrut/presenceof)

Validate presence of

## Install

    go get gopkg.in/gostrut/presenceof.v1

## Example

    type Person struct {
      Name string `presence_of:"true"`
    }

    val := NewValidator()
    val.Checks("presence_of", presenceof.Validator)

    p := Person{}
    fields, err := val.Validates(p)
    if err != nil {
      // handle validation error
    }
    if !fields.Valid() {
      // handle invalid fields
    }

## License

MIT
