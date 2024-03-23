package main

import (
	"github.com/EricLin0619/DcardBackend/app"
	// "fmt"
)

func main() {
    app := app.NewApp("8000")
	go func () {
		app.ListenToChannel()
	} ()
	app.Run()
}







