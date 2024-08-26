// apps/server/src/api/synonyms_test.go
package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"semantic-search-demo/src/app"
)

func TestGetSynonymsHandler(t *testing.T) {
	// Set up in-memory test database
	db := setupTestDB()

	// Seed the database with a known value
	synonymsJson := `[{"synonym":"automobile","score":0.9},{"synonym":"vehicle","score":0.8}]`
	synonym := &app.Synonym{
		Phrase:   "car",
		Synonyms: []byte(synonymsJson),
	}
	db.Create(synonym)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/synonyms?phrase=car", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	router := httprouter.New()
	router.GET("/synonyms", GetSynonymsHandler)

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Parse the response body
	var synonyms []Synonym
	if err := json.Unmarshal(rr.Body.Bytes(), &synonyms); err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}

	// Check if the response matches the expected value
	expectedSynonyms := []Synonym{
		{Synonym: "automobile", Score: 0.9},
		{Synonym: "vehicle", Score: 0.8},
	}
	if len(synonyms) != len(expectedSynonyms) {
		t.Errorf("Expected %d synonyms, got %d", len(expectedSynonyms), len(synonyms))
	}
	for i, synonym := range synonyms {
		if synonym.Synonym != expectedSynonyms[i].Synonym || synonym.Score != expectedSynonyms[i].Score {
			t.Errorf("Expected synonym %+v, got %+v", expectedSynonyms[i], synonym)
		}
	}
}
