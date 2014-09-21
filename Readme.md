# presenceof

Validate presence of

## Example

    type Person struct {
      Name string `presence_of:"true"`
    }

    val := NewValidator()
    val.Checks("presence_of", presenceof.Validator)

    p := Person{}
    fields, err := val.Validates(p)

## License

MIT