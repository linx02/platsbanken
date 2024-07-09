// TODO: Sleep between x amount of requests to avoid rate limiting

package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"platsbanken-api/api"
	"platsbanken-api/db"
	"strings"
	"sync"
	"time"
)

var database *sql.DB

var (
	progressMu sync.Mutex
	progress   int
)

func InitDB(dbInstance *sql.DB) {
	database = dbInstance
}

var retries = 0

func DownloadJobPosting(jobID string) {

	// Fetch job posting from API
	jobData := api.JobPosting(jobID)
	if jobData == nil {
		log.Fatalf("Error fetching job posting from API")
	} else if jobData["id"] == nil {
		time.Sleep(10 * time.Second)
		retries++
		if retries > 5 {
			log.Fatalf("Error fetching job posting from API, max retries reached")
		}
		DownloadJobPosting(jobID)
	}

	existingJob, err := db.GetJob(jobID)
	if err == nil {
		// Job already exists in the database - skip
		fmt.Printf("Job posting with ID %s already exists in the database\n", existingJob.ID)
		return
	}

	// Convert API response to db.JobPosting struct
	job := &db.JobPosting{
		ID:                  getStringValue(jobData, "id"),
		Title:               getStringValue(jobData, "title"),
		Description:         getStringValue(jobData, "description"),
		PublishedDate:       parseTime(getStringValue(jobData, "publishedDate")),
		LastApplicationDate: parseTime(getStringValue(jobData, "lastApplicationDate")),
		Occupation:          getStringValue(jobData, "occupation"),
		Conditions:          getStringValue(jobData, "conditions"),
		SalaryDescription:   getStringValue(jobData, "salaryDescription"),
		SalaryType:          getStringValue(jobData, "salaryType"),
		WorkTimeExtent:      getStringValue(jobData, "workTimeExtent"),
		EmploymentType:      getStringValue(jobData, "employmentType"),
		Duration:            getStringValue(jobData, "duration"),
		Positions:           getIntValue(jobData, "positions"),
		Published:           getBoolValue(jobData, "published"),
		OwnCar:              getBoolValue(jobData, "ownCar"),
		RequiresExperience:  getBoolValue(jobData, "requiresExperience"),
		Logotype:            getStringValue(jobData, "logotype"),
		Company: db.Company{
			Name:               getStringValue(jobData["company"].(map[string]interface{}), "name"),
			StreetAddress:      getStringValue(jobData["company"].(map[string]interface{}), "streetAddress"),
			PostCode:           getStringValue(jobData["company"].(map[string]interface{}), "postCode"),
			City:               getStringValue(jobData["company"].(map[string]interface{}), "city"),
			PhoneNumber:        getStringValue(jobData["company"].(map[string]interface{}), "phoneNumber"),
			WebAddress:         getStringValue(jobData["company"].(map[string]interface{}), "webAddress"),
			Email:              getStringValue(jobData["company"].(map[string]interface{}), "email"),
			OrganisationNumber: getStringValue(jobData["company"].(map[string]interface{}), "organisationNumber"),
		},
		Application: db.Application{
			Mail:        getStringValue(jobData["application"].(map[string]interface{}), "mail"),
			Email:       getStringValue(jobData["application"].(map[string]interface{}), "email"),
			WebAddress:  getStringValue(jobData["application"].(map[string]interface{}), "webAddress"),
			Other:       getStringValue(jobData["application"].(map[string]interface{}), "other"),
			Reference:   getStringValue(jobData["application"].(map[string]interface{}), "reference"),
			Information: getStringValue(jobData["application"].(map[string]interface{}), "information"),
		},
		Workplace: db.Workplace{
			Name:                 getStringValue(jobData["workplace"].(map[string]interface{}), "name"),
			Street:               getStringValue(jobData["workplace"].(map[string]interface{}), "street"),
			PostCode:             getStringValue(jobData["workplace"].(map[string]interface{}), "postCode"),
			City:                 getStringValue(jobData["workplace"].(map[string]interface{}), "city"),
			UnspecifiedWorkplace: getBoolValue(jobData["workplace"].(map[string]interface{}), "unspecifiedWorkplace"),
			Region:               getStringValue(jobData["workplace"].(map[string]interface{}), "region"),
			Country:              getStringValue(jobData["workplace"].(map[string]interface{}), "country"),
			Municipality:         getStringValue(jobData["workplace"].(map[string]interface{}), "municipality"),
			Longitude:            getFloat64Value(jobData["workplace"].(map[string]interface{}), "longitude"),
			Latitude:             getFloat64Value(jobData["workplace"].(map[string]interface{}), "latitude"),
			ShowMap:              getBoolValue(jobData["workplace"].(map[string]interface{}), "showMap"),
		},
	}

	// Add job posting to the database
	err = db.AddJob(job)
	if err != nil {
		log.Fatalf("Error adding job posting to database: %v", err)
	}
}

func DownloadAllJobPostings(query string) {

	// Reset progress
	progressMu.Lock()
	progress = 0
	progressMu.Unlock()

	numberOfPostings, err := api.GetNumberOfPostings(query)
	if err != nil {
		log.Fatalf("Error fetching number of postings from API: %v", err)
	}

	if numberOfPostings > 2000 {
		numberOfPostings = 2000
		fmt.Println("Number of postings limited to 2000")
	}

	for i := 0; i < numberOfPostings; i += 100 {
		result := api.Search(query, i, "", "")
		if result == nil {
			log.Fatalf("Error fetching job postings from API")
		}

		for _, posting := range result["ads"].([]interface{}) {
			jobID := posting.(map[string]interface{})["id"].(string)
			DownloadJobPosting(jobID)
		}

		progressMu.Lock()
		progress = (i + 100) * 100 / numberOfPostings
		progressMu.Unlock()
	}
}

func GetDownloadProgress() int {
	progressMu.Lock()
	defer progressMu.Unlock()
	return progress
}

func GetAmountOfJobPostings() int {
	amountOfJobs, err := db.GetTotalJobs()
	if err != nil {
		log.Fatalf("Error fetching total jobs: %v", err)
	}
	return amountOfJobs
}

func GetJobPosting(id string) ([]byte, error) {
	job, err := db.GetJob(id)
	if err != nil {
		return nil, fmt.Errorf("error fetching job posting: %v", err)
	}

	// Encode job posting to JSON
	data, err := json.Marshal(job)
	if err != nil {
		return nil, fmt.Errorf("error encoding job posting to JSON: %v", err)
	}

	return data, nil
}

type JobListing struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Occupation  string `json:"occupation"`
	CompanyName string `json:"companyName"`
	Location    string `json:"location"`
	Published   string `json:"published"`
}

func GetAllJobPostings() ([]byte, error) {
	jobs, err := db.GetAllJobs()
	if err != nil {
		return nil, fmt.Errorf("error fetching job postings: %v", err)
	}

	// Encode job postings to JSON
	data, err := json.Marshal(jobs)
	if err != nil {
		return nil, fmt.Errorf("error encoding job postings to JSON: %v", err)
	}

	return data, nil
}

func GetJobPostings(num int) ([]byte, error) {
	jobs, err := db.GetAllJobs()
	if err != nil {
		return nil, fmt.Errorf("error fetching job postings: %v", err)
	}

	var result []JobListing

	// Limit the number of job postings
	if num < len(jobs) {
		jobs = jobs[:num]
	}

	for _, job := range jobs {
		result = append(result, JobListing{
			ID:          job.ID,
			Title:       job.Title,
			Occupation:  job.Occupation,
			CompanyName: job.Company.Name,
			Location:    job.Workplace.Municipality,
			Published:   job.PublishedDate.String(),
		})
	}

	// Encode job postings to JSON
	data, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("error encoding job postings to JSON: %v", err)
	}

	return data, nil
}

func Search(positiveSearchTerms []string, negativeSearchTerms []string, advancedSearchQuery string) ([]byte, error) {
	jobs, err := db.GetAllJobs()
	if err != nil {
		return nil, fmt.Errorf("error fetching job postings: %v", err)
	}

	var result []*db.JobPosting

	if len(positiveSearchTerms) > 0 {
		result, err = PositiveSearch(positiveSearchTerms, jobs)
		if err != nil {
			return nil, fmt.Errorf("error performing positive search: %v", err)
		}
	}

	if len(negativeSearchTerms) > 0 {
		if len(result) == 0 {
			result = jobs
		}
		result, err = NegativeSearch(negativeSearchTerms, result)
		if err != nil {
			return nil, fmt.Errorf("error performing negative search: %v", err)
		}
	}

	if advancedSearchQuery != "" {
		if len(result) == 0 {
			result = jobs
		}
		result, err = AdvancedSearch(advancedSearchQuery, result)
		if err != nil {
			return nil, fmt.Errorf("error performing advanced search: %v", err)
		}
	}

	var data []JobListing

	for _, job := range result {
		data = append(data, JobListing{
			ID:          job.ID,
			Title:       job.Title,
			Occupation:  job.Occupation,
			CompanyName: job.Company.Name,
			Location:    job.Workplace.Municipality,
			Published:   job.PublishedDate.String(),
		})
	}

	// Encode job postings to JSON
	dataJson, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error encoding job postings to JSON: %v", err)
	}

	return dataJson, nil

}

func PositiveSearch(searchTerms []string, jobs []*db.JobPosting) ([]*db.JobPosting, error) {

	if jobs == nil {
		var err error
		jobs, err = db.GetAllJobs()
		if err != nil {
			return nil, fmt.Errorf("error fetching job postings: %v", err)
		}
	}

	var result []*db.JobPosting

	for _, job := range jobs {
		for _, term := range searchTerms {
			if strings.Contains(job.Description, term) {
				result = append(result, job)
			}
		}
	}

	return result, nil
}

func NegativeSearch(searchTerms []string, jobs []*db.JobPosting) ([]*db.JobPosting, error) {

	if jobs == nil {
		var err error
		jobs, err = db.GetAllJobs()
		if err != nil {
			return nil, fmt.Errorf("error fetching job postings: %v", err)
		}
	}

	var result []*db.JobPosting

jobLoop:
	for _, job := range jobs {
		for _, term := range searchTerms {
			if strings.Contains(job.Description, term) {
				// If job contains any negative search term, skip it
				continue jobLoop
			}
		}
		// Job passed all negative search term checks, include it in the result
		result = append(result, job)
	}

	return result, nil
}

func AdvancedSearch(queryString string, jobs []*db.JobPosting) ([]*db.JobPosting, error) {

	if jobs == nil {
		var err error
		jobs, err = db.GetAllJobs()
		if err != nil {
			return nil, fmt.Errorf("error fetching job postings: %v", err)
		}
	}

	query, err := ParseQuery(queryString)
	if err != nil {
		return nil, fmt.Errorf("error parsing query: %v", err)
	}

	filtered := FilterJobPostings(jobs, query.Conditions)

	return filtered, nil
}
