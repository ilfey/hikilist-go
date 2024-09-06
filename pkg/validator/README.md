# Validator

## Example

### Struct

```go
type Example struct {
    Username string
    Password string
    Email    *string
}
```

### Code

```go
st := Example{}

vErr := validator.Validate(st, map[string][]options.Option{
    "Username": {
        options.Required(),
        options.LenGreaterThan(3),
    },
    "Password": {
        options.Required(),
        options.LenGreaterThan(5),
    },
    "Email": {
        options.LenLessThan(64),
    },
})
if vErr != nil {
    // handle error...
}

// Validated data
```