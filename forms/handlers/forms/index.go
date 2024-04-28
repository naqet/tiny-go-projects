package forms

import (
	"net/http"

	"gorm.io/gorm"
)

type Forms struct { }

func Init(mux *http.ServeMux, database *gorm.DB) {
    db := newFormsDb(database)
    db.init()

    service := newFormsService(db)
    handler := newFormsHandler(mux, service)

    handler.init()
}
