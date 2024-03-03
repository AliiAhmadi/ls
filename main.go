package main

type App struct {
	config Config
}

type Config struct {
	// extension of files want to filter
	ext string

	// Minimum file size in bytes
	min_size int64

	// Maximum file size in bytes
	max_size int64

	// verbose mode
	verbose bool

	// list mode
	list bool
}

func main() {

}

func (app *App) Parse() {

}
