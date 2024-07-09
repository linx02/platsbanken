package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var url = "https://platsbanken-api.arbetsformedlingen.se"

func Search(query string, startIndex int, fromDate string, toDate string) map[string]interface{} {
	endpoint := url + "/jobs/v1/search"
	payload := map[string]interface{}{
		"filters": []map[string]interface{}{
			{"type": "freetext", "value": query},
		},
		"fromDate":   fromDate,
		"toDate":     toDate,
		"order":      "relevance",
		"maxRecords": 100,
		"startIndex": startIndex,
		"source":     "pb",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Read and print the response body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return result
}

func GetNumberOfPostings(query string) (int, error) {
	endpoint := url + "/jobs/v1/search"
	payload := map[string]interface{}{
		"filters": []map[string]interface{}{
			{"type": "freetext", "value": query},
		},
		"fromDate":   nil,
		"toDate":     nil,
		"order":      "relevance",
		"maxRecords": 0,
		"startIndex": 0,
		"source":     "pb",
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return 0, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return 0, err
	}
	defer resp.Body.Close()

	// Read and print the response body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return 0, err
	}

	// Extract the value from the result map
	positionsFloat, ok := result["positions"].(float64)
	if !ok {
		return 0, fmt.Errorf("expected positions to be float64, got %T", result["positions"])
	}

	// Convert float64 to int
	positions := int(positionsFloat)

	return positions, nil
}

func JobPosting(id string) map[string]interface{} {
	endpoint := url + "/jobs/v1/job/" + id
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error getting job posting:", err)
		return nil
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return result

}
