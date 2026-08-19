package main


import(
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"net/http"
)
	
type PredictionRequest struct {
	SmokeDetector int `json:"smoke_detector"`
	NewBatteries int `json:"new_batteries"`
	ABCExtinguisher int `json:"abc_extinguisher"`
	ClearExitRoutes int `json:"clear_exit_routes"`
}

type PredictionResponse struct {
	ClassID int `json:"class_id"`
	Probability float64 `json:"probability"`
}

type Handler struct {
	FastAPI *FastAPIClient
}

type FastAPIClient struct {
	BaseURL string
}

func (h *Handler)predictHandler(w http.ResponseWriter, r *http.Request) {
	var request PredictionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// Convert request to JSON
	_, err = io.Copy(os.Stdout, r.Body)
	jsonData, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "JSON error", http.StatusInternalServerError)
		return
	}
	//fmt.Println(string(jsonData))
	// Call Python ML service
	resp, err := http.Post(
					h.FastAPI.BaseURL + "/predict",
					"application/json",
					bytes.NewBuffer(jsonData),
					)
	if err != nil {
		http.Error(w, "ML service unavailable", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	// Read Pythonresponse
	//_, err = io.Copy(os.Stdout, resp.Body)
	
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
	var status string = "pending"
	
	if prediction.ClassID == 1 {
       // 3. Assign a new string value using the = operator
       status = "Yes" 
    } else {
       status = "No"
    }
	fmt.Printf("Received payload: prediction.ClassID=%d, prediction.Probability=%f\n", prediction.ClassID, prediction.Probability)
	result := map[string]interface{}{
		"status":   status,
		"probability":  prediction.Probability,
	}
	// Return prediction to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (c *FastAPIClient) GetData() (*http.Response, error) {
	return http.Post(
		c.BaseURL+"/getdata",
		"application/json",
		nil,
	)
}




func (h *Handler)dataHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := h.FastAPI.GetData()
	if err != nil {
		http.Error(w, "FastAPI request failed", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		http.Error(w, "FastAPI error", resp.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Copy FastAPI response directly to HTTP response
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return
	}
}


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		
		http.ServeFile(w, r, "index.html")
	})
	
	fastAPI := &FastAPIClient{
		BaseURL: os.Getenv("FASTAPI_URL"),
	}

	handler := &Handler{
		FastAPI: fastAPI,
	}
	http.HandleFunc("/predict", handler.predictHandler)
	http.HandleFunc("/getdata", handler.dataHandler)
	fmt.Println("Go backend running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}