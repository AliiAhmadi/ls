package main

import "flag"

type App struct {
	config *Config
}

type Config struct {
	// extension of files want to filter
	ext *string

	// Minimum file size in bytes
	min_size *int64

	// Maximum file size in bytes
	max_size *int64

	// verbose mode
	verbose *bool

	// list mode
	list *bool
}

func main() {
	app := App{}
	app.Parse()
}

func (app *App) Parse() {
	// Parse command line flags
	flag.Int64Var(app.config.min_size, "min", 0, "Minimum file size")
	flag.Int64Var(app.config.max_size, "max", int64(100000000000), "Maximum file size")
	flag.BoolVar(app.config.list, "list", false, "List mode")
	flag.BoolVar(app.config.verbose, "v", false, "Verbose mode")
	flag.StringVar(app.config.ext, "ext", ".", "Extension of files")
	flag.Parse()
}
