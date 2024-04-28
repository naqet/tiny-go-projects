package forms

import (
	"gorm.io/gorm"
)

type formsDb struct {
	db *gorm.DB
}

func newFormsDb(db *gorm.DB) *formsDb {
	return &formsDb{db}
}

func (f *formsDb) init() {
	f.db.AutoMigrate(&Form{}, &Field{})
}

func (f *formsDb) getAll() ([]Form, error) {
	forms := []Form{}
	result := f.db.Find(&forms)

	if result.Error != nil {
		return []Form{}, result.Error
	}

	return forms, nil
}

func (f *formsDb) create(data *Form) error {
    result := f.db.Create(data)
    return result.Error
}

type Form struct {
	gorm.Model
	Title  string
	Fields []Field
}

type Field struct {
	gorm.Model
	Question string
	Answer   string
	FormID   uint
}
