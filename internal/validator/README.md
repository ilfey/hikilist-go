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

vErr := validator.Validate(st, map[string][]validator.Option{
    "Username": {
        validator.Required(),
        validator.LenGreaterThat(3),
    },
    "Password": {
        validator.Required(),
        validator.LenGreaterThat(5),
    },
    "Email": {
        validator.LenLessThat(64),
    },
})
if vErr != nil {
    // handle error...
}

// Validated data
```