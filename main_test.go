// Author: Ali Ahmadi <ali.ahmadi9@ut.ac.ir>
// Dedicated to MHM

package main

import (
	"bytes"
	"os"
	"testing"
)

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

func TestFilterOutFunction(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		maxSize  int64
		expected bool
	}{
		{
			name:     "filter no extension",
			file:     "testdata/test.log",
			ext:      "",
			minSize:  0,
			maxSize:  100,
			expected: false,
		},
		{
			name:     "filter extension match",
			file:     "testdata/test.log",
			ext:      ".log",
			minSize:  0,
			maxSize:  100,
			expected: false,
		},
		{
			name:     "filter extension no match",
			file:     "testdata/test.log",
			ext:      ".sh",
			minSize:  0,
			maxSize:  100,
			expected: true,
		},
		{
			name:     "filter extension size match",
			file:     "testdata/test.log",
			ext:      ".log",
			minSize:  10,
			maxSize:  100,
			expected: false,
		},
		{
			name:     "filter extension size no match",
			file:     "testdata/test.log",
			ext:      ".log",
			minSize:  200,
			maxSize:  1000,
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			info, err := os.Stat(test.file)

			if err != nil {
				t.Fatal(err)
			}

			res := filterOut(test.file, &test.ext, &test.minSize, &test.maxSize, info)

			if res != test.expected {
				t.Errorf("TestFilterOutFunction: expected '%t' - got '%t'", test.expected, res)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		app      *App
		expected string
	}{
		{
			name: "no filter",
			app: &App{
				config: &Config{
					ext:      StringPtr(""),
					min_size: Int64Ptr(0),
					max_size: Int64Ptr(100),
					list:     BoolPtr(true),
					root:     StringPtr("testdata"),
				},
			},
			expected: "testdata/test.log\ntestdata/test.sh\n",
		},
		{
			name: "filter extension match",
			app: &App{
				config: &Config{
					ext:      StringPtr(".sh"),
					min_size: Int64Ptr(10),
					max_size: Int64Ptr(1000),
					list:     BoolPtr(false),
					root:     StringPtr("testdata"),
				},
			},
			expected: "testdata/test.sh\n",
		},
		{
			name: "filter extension size match",
			app: &App{
				config: &Config{
					ext:      StringPtr(".log"),
					min_size: Int64Ptr(10),
					max_size: Int64Ptr(1000),
					list:     BoolPtr(true),
					root:     StringPtr("testdata"),
				},
			},
			expected: "testdata/test.log\n",
		},
		{
			name: "filter extension size no match",
			app: &App{
				config: &Config{
					ext:      StringPtr(".sh"),
					min_size: Int64Ptr(100),
					max_size: Int64Ptr(1000),
					list:     BoolPtr(true),
					root:     StringPtr("testdata"),
				},
			},
			expected: "",
		},
		{
			name: "filter extension size no match",
			app: &App{
				config: &Config{
					ext:      StringPtr(".sh"),
					min_size: Int64Ptr(0),
					max_size: Int64Ptr(3),
					list:     BoolPtr(true),
					root:     StringPtr("testdata"),
				},
			},
			expected: "",
		},
		{
			name: "filter extension no match",
			app: &App{
				config: &Config{
					ext:      StringPtr(".exe"),
					min_size: Int64Ptr(0),
					max_size: Int64Ptr(100),
					list:     BoolPtr(false),
					root:     StringPtr("testdata"),
				},
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer

			test.app.out = &buf

			if err := test.app.run(); err != nil {
				t.Fatal(err)
			}

			res := buf.String()

			if test.expected != res {
				t.Errorf("TestRun: expected '%s' - got '%s'", test.expected, res)
			}
		})
	}
}
