package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
	"semantic-search-demo/src/app"
)

type Synonym struct {
	Synonym string  `json:"synonym" validate:"required"`
	Score   float64 `json:"score" validate:"required"`
}

type SynonymRequest struct {
	Phrase   string    `json:"phrase" validate:"required"`
	Synonyms []Synonym `json:"synonyms" validate:"required,dive"`
}

func AddSynonymsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var req SynonymRequest

	// Decode and validate request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert synonyms to JSONB format
	synonymsJson, err := json.Marshal(req.Synonyms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Insert/Update synonyms in the database
	synonym := &app.Synonym{
		Phrase:   req.Phrase,
		Synonyms: synonymsJson,
	}

	if err := app.DB.Clauses(gorm.Clauses{}).Save(synonym).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Synonyms saved successfully"))
}
