package main

import (
	"naqet/forms/handlers"
	"naqet/forms/handlers/forms"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

    if err != nil {
        panic("failed connection to db")
    }

	mux := http.NewServeMux()

    // Static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    // Pages
    pages := handlers.NewPagesHandler()
	mux.HandleFunc("/", pages.Home)

    // Endpoints
    forms.Init(mux, db)

	http.ListenAndServe("127.0.0.1:8000", mux)
}
