package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
	"semantic-search-demo/src/app"
	"io/ioutil"
)

func GetSynonymsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	phrase := r.URL.Query().Get("phrase")
	if phrase == "" {
		http.Error(w, "Phrase is required", http.StatusBadRequest)
		return
	}

	var synonym app.Synonym
	if err := app.DB.First(&synonym, "phrase = ?", phrase).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Call semantic-tools to generate synonyms
			response, err := http.Post("http://semantic-tools:8000/semantic-tools/synonyms/",
				"application/json", 
				strings.NewReader(fmt.Sprintf(`{"phrase": "%s"}`, phrase)))

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer response.Body.Close()

			if response.StatusCode != http.StatusOK {
				http.Error(w, "Failed to get synonyms", response.StatusCode)
				return
			}

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(synonym.Synonyms)
}
