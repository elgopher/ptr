// (c) 2023 Jacek Olszak
// This code is licensed under MIT license (see LICENSE for details)

// Package ptr provides generic functions to get optional values.
package ptr

// To returns a pointer to value.
//
// This function is very useful for single-line expressions, where "&" cannot be used:
//
//	in := struct {
//		OptionalField *string
//	}{
//		OptionalField: ptr.To("works fine"),
//	}
//
// To always create a pointer to copy of the value:
//
//	value := "v1"
//	in.OptionalField = ptr.To(value)
//	value = "v2" // in.OptionalField is not changed
func To[T any](t T) *T {
	return &t
}

// Value returns a value behind the pointer or zero value when value is nil.
func Value[T any](ptr *T) T {
	if ptr != nil {
		return *ptr
	}

	var zeroValue T

	return zeroValue
}

// ValueOrDefault returns a value behind the pointer or defaultValue when value is nil.
func ValueOrDefault[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}

	return defaultValue
}

// Copy copies value behind the pointer and returns a new pointer to the copy.
func Copy[T any](in *T) *T {
	var out *T

	if in != nil {
		out = To(*in)
	}

	return out
}
