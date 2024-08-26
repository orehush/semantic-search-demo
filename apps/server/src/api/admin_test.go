package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"semantic-search-demo/src/app"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&app.Synonym{})
	app.DB = db
	return db
}

func TestAddSynonymsHandler(t *testing.T) {
	// Set up in-memory test database
	db := setupTestDB()

	// Prepare request payload
	payload := SynonymRequest{
		Phrase: "car",
		Synonyms: []Synonym{
			{Synonym: "automobile", Score: 0.9},
			{Synonym: "vehicle", Score: 0.8},
		},
	}
	body, _ := json.Marshal(payload)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/admin/synonyms", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.POST("/admin/synonyms", AddSynonymsHandler)

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check if the data was saved correctly
	var synonym app.Synonym
	db.First(&synonym, "phrase = ?", "car")

	if synonym.Phrase != "car" {
		t.Errorf("Expected phrase 'car', got %v", synonym.Phrase)
	}
	if string(synonym.Synonyms) != `[{"synonym":"automobile","score":0.9},{"synonym":"vehicle","score":0.8}]` {
		t.Errorf("Unexpected synonyms: %v", string(synonym.Synonyms))
	}
}
