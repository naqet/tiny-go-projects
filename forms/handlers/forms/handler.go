package forms

import (
	"encoding/json"
	"naqet/forms/views/fields"
	"net/http"
	"strconv"
)

type formsHandler struct {
	mux     *http.ServeMux
	service *formsService
}

func newFormsHandler(mux *http.ServeMux, service *formsService) *formsHandler {
	return &formsHandler{mux, service}
}

func (h *formsHandler) init() {
	h.mux.HandleFunc("/forms", h.get)
	h.mux.HandleFunc("POST /forms", h.create)
	h.mux.HandleFunc("/forms/add-field", h.addField)
}

func (h *formsHandler) get(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.getAll()

	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(result)

	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func (h *formsHandler) create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")

	if len(title) < 1 {
		http.Error(w, "title fields cannot be empty", http.StatusBadRequest)
		return
	}

    if r.MultipartForm == nil {
		http.Error(w, "Form needs to be submitted as multipart", http.StatusBadRequest)
		return
    }

	questions := r.MultipartForm.Value["question"]

	if len(questions) < 1 {
		http.Error(w, "questions need to be provided", http.StatusBadRequest)
		return
	}

	req, err := h.service.create(title, questions)

	if err != nil {
		http.Error(w, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(strconv.Itoa(int(req))))
}

func(h *formsHandler) addField(w http.ResponseWriter, r *http.Request) {
    fields.Question().Render(r.Context(), w)
}
