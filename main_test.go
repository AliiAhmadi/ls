package main

import "testing"

func TestInputFlags(t *testing.T) {

}

func TestPointerGenerators(t *testing.T) {
	t.Run("string pointer", func(t *testing.T) {
		tests := []struct {
			value string
		}{
			{""},
			{"string"},
			{},
		}

		for _, test := range tests {
			ptr := StringPtr(test.value)

			if test.value != *ptr {
				t.Errorf("Test string pointer: invalid value - expected %s - got %s", test.value, *ptr)
			}
		}
	})

	t.Run("bool pointer", func(t *testing.T) {
		tests := []struct {
			value bool
		}{
			{true},
			{false},
		}

		for _, test := range tests {
			ptr := BoolPtr(test.value)

			if test.value != *ptr {
				t.Errorf("Test bool pointer: invalid value - expected %v - got %v", test.value, *ptr)
			}
		}
	})

	t.Run("int64 pointer", func(t *testing.T) {
		tests := []struct {
			number int64
		}{
			{0},
			{10},
			{-100},
		}

		for _, test := range tests {
			ptr := Int64Ptr(test.number)

			if test.number != *ptr {
				t.Errorf("Test int64 pointer: invalid value - expected %d - got %d", test.number, *ptr)
			}
		}
	})
}
