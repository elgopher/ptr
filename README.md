# ptr
Generic functions to get optional values

[![Go Reference](https://pkg.go.dev/badge/github.com/elgopher/ptr.svg)](https://pkg.go.dev/github.com/elgopher/ptr)

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

Sometimes even unsafe:

```go
value := "v1"
in.OptionalField = &value
value = "v2" // ups... in.OptionalField is changed too! 
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

Safe code:

```go
value := "v1"
in.OptionalField = ptr.To(value)
value = "v2" // in.OptionalField is not changed
```

```go
copyOfPointer := ptr.Copy(in.OptionalField)
```

## Installation

```shell
go get github.com/elgopher/ptr@latest
```