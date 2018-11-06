package main

import (
	"flag"
	"os"

	"settlement-report/src/application"
)

var app *application.Application

func init() {
	app = createApplication()
}

func createApplication() *application.Application {
	filename := flag.String("filename", "", "The text file that contains the instruction data. (Required)")
	flag.Parse()

	if *filename == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return application.NewApplication(&application.Config{
		DataFilename: *filename,
	})
}

func main() {
	app.Run()
}
