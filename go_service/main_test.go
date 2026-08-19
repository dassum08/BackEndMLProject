package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

type Response struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func TestHandler(t *testing.T) {
	fakeFastAPI := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"name":"John","age":30}`))
		}),
	)
	defer fakeFastAPI.Close()

	fastAPI := &FastAPIClient{
		BaseURL: fakeFastAPI.URL,
	}

	handler := &Handler{
		FastAPI: fastAPI,
	}

	req := httptest.NewRequest("POST", "/data", nil)
	rec := httptest.NewRecorder()

	handler.dataHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	// Check response body
	var result Response

	err := json.Unmarshal(rec.Body.Bytes(), &result)
	if err != nil {
		t.Fatal(err)
	}

	if result.Name != "John" {
		t.Errorf("expected John, got %s", result.Name)
	}

	if result.Age != 30 {
		t.Errorf("expected 30, got %d", result.Age)
	}
}