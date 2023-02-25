// (c) 2023 Jacek Olszak
// This code is licensed under MIT license (see LICENSE for details)

package ptr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/elgopher/ptr"
)

func TestTo(t *testing.T) {
	t.Run("should return pointer to value", func(t *testing.T) {
		value := "value"
		pointer := ptr.To(value)
		require.NotNil(t, pointer)
		assert.Equal(t, value, *pointer)
	})

	t.Run("pointer should point to copy of the value", func(t *testing.T) {
		value := "before"
		pointer := ptr.To(value)
		value = "after"
		assert.Equal(t, "before", *pointer)
	})
}

func TestValue(t *testing.T) {
	t.Run("should return pointer value", func(t *testing.T) {
		expected := "value"
		pointer := ptr.To(expected)
		actual := ptr.Value(pointer)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return zero value when pointer is nil", func(t *testing.T) {
		var nilPointer *string
		actual := ptr.Value(nilPointer)
		assert.Zero(t, actual)
	})
}

func TestValueOrDefault(t *testing.T) {
	defaultValue := "default"

	t.Run("should return pointer value", func(t *testing.T) {
		expected := "value"
		pointer := ptr.To(expected)
		actual := ptr.ValueOrDefault(pointer, defaultValue)
		assert.Equal(t, expected, actual)
	})

	t.Run("should return default value when pointer is nil", func(t *testing.T) {
		var nilPointer *string
		actual := ptr.ValueOrDefault(nilPointer, defaultValue)
		assert.Equal(t, defaultValue, actual)
	})
}

func TestCopy(t *testing.T) {
	t.Run("should return nil", func(t *testing.T) {
		var nilString *string
		theCopy := ptr.Copy(nilString)
		assert.Nil(t, theCopy)
	})

	t.Run("should return new pointer", func(t *testing.T) {
		original := ptr.To("str")
		theCopy := ptr.Copy(original)
		assert.Equal(t, *original, *theCopy)
		assert.NotSame(t, original, theCopy)
	})

	t.Run("should return new pointer to struct", func(t *testing.T) {
		original := &struct{ field string }{field: "value"}
		theCopy := ptr.Copy(original)
		assert.Equal(t, *original, *theCopy)
		assert.NotSame(t, original, theCopy)
	})

	t.Run("pointer should point to copy of the value", func(t *testing.T) {
		original := "original"
		theCopy := ptr.Copy(&original)
		original = "altered"
		assert.Equal(t, "original", *theCopy)
	})
}
