# ptr
Generic functions to get optional values

## Problems with Go optional values

By convention, optional values in Go are pointers:

```go 
type Input struct {
	RequiredField string
	OptionalField *string
}
```

But such structs **cannot** be initialized in a single expression:

```go
in := Input{
	RequiredField: "works"
	OptionalField: &("does not work")
}
```

And accessing optional fields makes code look ugly:

```go
if in.OptionalField != nil && *in.OptionalField == "value" {
    ...    	
}
```

## This tiny packages simplifies the use of optional values

One-line initialization:

```go
import "github.com/elgopher/ptr"

in := Input{
	RequiredField: "works"
	OptionalField: ptr.To("this also works")
}
```

Getting values without boilerplate code:

```go
if ptr.Value(in.OptionalField) == "value" {
	// if in.OptionalField is nil then zero value is returned ("" for string)
    ...    	
}
```

```go
v := ptr.ValueOrDefault(in.OptionalField, "defaultValue")
```