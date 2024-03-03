package main

import (
	"flag"
	"fmt"
	"os"
)

type App struct {
	config *Config
}

type Config struct {
	// Extension of files want to filter
	ext *string

	// Minimum file size in bytes
	min_size *int64

	// Maximum file size in bytes
	max_size *int64

	// Verbose mode
	verbose *bool

	// List mode
	list *bool

	// Root path to start
	root *string
}

func main() {
	app := &App{
		config: &Config{
			ext:      StringPtr(""),
			min_size: Int64Ptr(0),
			max_size: Int64Ptr(0),
			verbose:  BoolPtr(false),
			list:     BoolPtr(false),
			root:     StringPtr("."),
		},
	}
	app.Parse()

	if err := app.run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func (app *App) Parse() {
	// Parse command line flags
	flag.Int64Var(app.config.min_size, "min", 0, "Minimum file size")
	flag.Int64Var(app.config.max_size, "max", int64(100000000000), "Maximum file size")
	flag.BoolVar(app.config.list, "list", false, "List mode")
	flag.BoolVar(app.config.verbose, "v", false, "Verbose mode")
	flag.StringVar(app.config.ext, "ext", "", "Extension of files")
	flag.StringVar(app.config.root, "root", ".", "Root path to start walk")
	flag.Parse()
}

func StringPtr(str string) *string {
	return &str
}

func BoolPtr(b bool) *bool {
	return &b
}

func Int64Ptr(number int64) *int64 {
	return &number
}

func (app *App) run() error {
	return nil
}
