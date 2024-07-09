package service

import (
	"log"
	"platsbanken-api/db"
	"regexp"
	"strings"
	"time"
)

// getStringValue retrieves a string value from a map[string]interface{} or returns an empty string if not found or not a string
func getStringValue(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return ""
}

// getIntValue retrieves an integer value from a map[string]interface{} or returns 0 if not found or not a number
func getIntValue(m map[string]interface{}, key string) int {
	if val, ok := m[key]; ok {
		if numVal, ok := val.(float64); ok {
			return int(numVal)
		}
	}
	return 0
}

// getFloat64Value retrieves a float64 value from a map[string]interface{} or returns 0.0 if not found or not a number
func getFloat64Value(m map[string]interface{}, key string) float64 {
	if val, ok := m[key]; ok {
		if numVal, ok := val.(float64); ok {
			return numVal
		}
	}
	return 0.0
}

// getBoolValue retrieves a boolean value from a map[string]interface{} or returns false if not found or not a boolean
func getBoolValue(m map[string]interface{}, key string) bool {
	if val, ok := m[key]; ok {
		if boolVal, ok := val.(bool); ok {
			return boolVal
		}
	}
	return false
}

// parseTime parses time string into Time object
func parseTime(timeStr string) time.Time {
	if timeStr == "" {
		return time.Time{}
	}
	layout := "2006-01-02T15:04:05Z" // Adjusted layout without milliseconds
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		log.Fatalf("Error parsing time: %v", err)
	}
	return t
}

type Operator int

const (
	AND Operator = iota
	OR
)

type Condition struct {
	Field    string // Field to check (e.g., "title", "description")
	Operator string // Operator (e.g., "in", "not in")
	Value    string // Value to check against (e.g., "developer", "bachelors degree")
}

// JobPosting represents a job posting with title and description
type JobPosting struct {
	Title       string
	Description string
}

// Query represents a logical query (AND or OR of conditions)
type Query struct {
	Operator   Operator    // AND or OR
	Conditions []Condition // List of conditions
}

// ParseQuery parses a query string into a Query struct
func ParseQuery(query string) (Query, error) {
	var q Query
	conditions := make([]Condition, 0)

	// Regular expression to match quoted strings
	re := regexp.MustCompile(`(['"])([^'"]*)['"]`)

	// Split query by "and" or "or" to determine the main operator
	parts := strings.FieldsFunc(query, func(r rune) bool {
		return r == '(' || r == ')' || r == ','
	})

	var mainOperator Operator = AND
	for _, part := range parts {
		switch strings.ToLower(part) {
		case "and":
			mainOperator = AND
		case "or":
			mainOperator = OR
		case "not", "in", "title", "description":
			continue // Skip these keywords
		default:
			// Find all quoted values in the part
			matches := re.FindAllStringSubmatch(part, -1)
			if len(matches) > 0 && len(matches[0]) == 3 {
				value := matches[0][2]
				// Determine the operator
				if strings.Contains(part, " not in ") {
					conditions = append(conditions, Condition{
						Field:    strings.TrimSpace(part[strings.Index(part, " not in ")+len(" not in "):]),
						Operator: "not in",
						Value:    value,
					})
				} else if strings.Contains(part, " in ") {
					conditions = append(conditions, Condition{
						Field:    strings.TrimSpace(part[strings.Index(part, " in ")+len(" in "):]),
						Operator: "in",
						Value:    value,
					})
				}
			}
		}
	}

	q.Operator = mainOperator
	q.Conditions = conditions

	return q, nil
}

// FilterJobPostings filters job postings based on given conditions
func FilterJobPostings(postings []*db.JobPosting, conditions []Condition) []*db.JobPosting {
	var filtered []*db.JobPosting

	// Iterate through each job posting
	for _, posting := range postings {
		// Check if the posting satisfies all conditions
		if matchesConditions(posting, conditions) {
			filtered = append(filtered, posting)
		}
	}

	return filtered
}

// matchesConditions checks if a job posting matches all given conditions
func matchesConditions(posting *db.JobPosting, conditions []Condition) bool {
	for _, condition := range conditions {
		switch condition.Field {
		case "title":
			if condition.Operator == "in" {
				lower := strings.ToLower(posting.Title)
				if !strings.Contains(lower, condition.Value) {
					return false
				}
			} else if condition.Operator == "not in" {
				lower := strings.ToLower(posting.Title)
				if strings.Contains(lower, condition.Value) {
					return false
				}
			}
		case "description":
			if condition.Operator == "in" {
				lower := strings.ToLower(posting.Description)
				if !strings.Contains(lower, condition.Value) {
					return false
				}
			} else if condition.Operator == "not in" {
				lower := strings.ToLower(posting.Description)
				if strings.Contains(lower, condition.Value) {
					return false
				}
			}
		default:
			// Handle unknown fields or operators
			return false
		}
	}
	return true
}
