package forms

type formsService struct {
    db *formsDb
}

func newFormsService(db *formsDb) *formsService {
    return &formsService{db}
}

func (s *formsService) getAll() ([]Form, error) {
    return s.db.getAll()
}

func (s *formsService) create(title string, questions []string) (uint, error) {
    fields := []Field{}

    for _, q := range questions {
        fields = append(fields, Field{Question: q})
    }

    form := &Form{Title: title, Fields: fields}
    err := s.db.create(form)
    return form.ID, err
}
