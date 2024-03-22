package main

import (
	"github.com/EricLin0619/DcardBackend/application"
	"github.com/EricLin0619/DcardBackend/db"
)

func main() {
	db := db.NewDb()
	db.DbTest()
	
    app := application.NewApp("8000")
	app.Run()
}






