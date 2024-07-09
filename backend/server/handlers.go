package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"platsbanken-api/service"
	"strconv"

	"github.com/gorilla/mux"
)

func InitialDownloadHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Query string `json:"query"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	go service.DownloadAllJobPostings(requestData.Query)

	response := map[string]string{
		"message": "triggered",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetDownloadProgressHandler(w http.ResponseWriter, r *http.Request) {
	progress := service.GetDownloadProgress()
	response := map[string]int{
		"message": progress,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func QuickUpdateHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"message": "Quick update triggered",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		PositiveSearchTerms []string `json:"positiveSearchTerms"`
		NegativeSearchTerms []string `json:"negativeSearchTerms"`
		AdvancedSearchQuery string   `json:"advancedSearchQuery"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Placeholder for search logic
	// For now, just return the search terms

	data, err := service.Search(requestData.PositiveSearchTerms, requestData.NegativeSearchTerms, requestData.AdvancedSearchQuery)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching job postings: %v", err), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}

// func PositiveSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	var requestData struct {
// 		SearchTerms []string `json:"searchTerms"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	// Placeholder for search logic
// 	// For now, just return the search terms

// 	data, err := service.PositiveSearch(requestData.SearchTerms, nil)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error fetching job postings: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Set Content-Type header
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write JSON data to response
// 	_, err = w.Write(data)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error writing JSON response: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// }

// func NegativeSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	var requestData struct {
// 		SearchTerms []string `json:"searchTerms"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	// Placeholder for search logic
// 	// For now, just return the search terms

// 	data, err := service.NegativeSearch(requestData.SearchTerms, nil)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error fetching job postings: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Set Content-Type header
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write JSON data to response
// 	_, err = w.Write(data)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error writing JSON response: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// }

// func AdvancedSearchHandler(w http.ResponseWriter, r *http.Request) {
// 	var requestData struct {
// 		SearchTerm string `json:"searchTerm"`
// 	}
// 	err := json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Bad request", http.StatusBadRequest)
// 		return
// 	}

// 	results, err := service.AdvancedSearch(requestData.SearchTerm, nil)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error fetching job postings: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	_, err = w.Write(results)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error writing JSON response: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// }

// 	// Placeholder for search logic
// 	// For now, just return the search terms

// 	data, err := service.AdvancedSearch(requestData.SearchTerm, nil)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error fetching job postings: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	// Set Content-Type header
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write JSON data to response
// 	_, err = w.Write(data)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Error writing JSON response: %v", err), http.StatusInternalServerError)
// 		return
// 	}
// }

func GetJobPostingsHandler(w http.ResponseWriter, r *http.Request) {
	amountStr := mux.Vars(r)["amount"]

	// Convert amount to integer
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Call service function to retrieve job posting
	data, err := service.GetJobPostings(amount)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching job posting: %v", err), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}

func GetJobPostingHandler(w http.ResponseWriter, r *http.Request) {
	// Extract job ID from request parameters
	id := mux.Vars(r)["id"]

	// Call service function to retrieve job posting
	data, err := service.GetJobPosting(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching job posting: %v", err), http.StatusInternalServerError)
		return
	}

	// Set Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write JSON data to response
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing JSON response: %v", err), http.StatusInternalServerError)
		return
	}
}

func GetAmountOfJobPostingsHandler(w http.ResponseWriter, r *http.Request) {
	// Placeholder for fetching amount of job postings from DB
	// For now, just return a static number
	response := map[string]int{
		"message": service.GetAmountOfJobPostings(),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
