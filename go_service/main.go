package main


import(
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)
	
type PredictionRequest struct {
	SepalLength float64 `json:"sepal_length"`
	SepalWidth float64 `json:"sepal_width"`
	PetalLength float64 `json:"petal_length"`
	PetalWidth float64 `json:"petal_width"`
}

type PredictionResponse struct {
	ClassID int `json:"class_id"`
	Probability float64 `json:"probability"`
}

func predictHandler(w http.ResponseWriter, r *http.Request) {
	var request PredictionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Convert request to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "JSON error", http.StatusInternalServerError)
		return
	}
	// Call Python ML service
	resp, err := http.Post(
					"http://fastapi:8000/predict",
					"application/json",
					bytes.NewBuffer(jsonData),
					)
	if err != nil {
		http.Error(w, "ML service unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	// Read Pythonresponse
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w,"Failed to read ML response",http.StatusInternalServerError)
		return
	}
	if resp.StatusCode != http.StatusOK {
		http.Error(w, string(body), resp.StatusCode)
		return
	}
	var prediction PredictionResponse
	err = json.Unmarshal(body, &prediction)
	if err != nil {
		http.Error(w,"Invalid ML response",http.StatusInternalServerError)
		return
	}
	// Return prediction to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prediction)
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/predict", predictHandler)
	fmt.Println("Go backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}