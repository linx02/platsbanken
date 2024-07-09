// Salarydescription, salarytype, worktimeextent, employmenttype
// Requiresexperience incorrect, application, workplace empty

package db

import "time"

// To be implemented
type SearchQuery struct {
	Query      string
	LastSearch time.Time
}

type JobPosting struct {
	ID                  string       `db:"id" json:"id"`
	Title               string       `db:"title" json:"title"`
	Description         string       `db:"description" json:"description"`
	PublishedDate       time.Time    `db:"published_date" json:"publishedDate"`
	LastApplicationDate time.Time    `db:"last_application_date" json:"lastApplicationDate"`
	Occupation          string       `db:"occupation" json:"occupation"`
	Conditions          string       `db:"conditions" json:"conditions"`
	SalaryDescription   string       `db:"salary_description" json:"salaryDescription"`
	SalaryType          string       `db:"salary_type" json:"salaryType"`
	WorkTimeExtent      string       `db:"work_time_extent" json:"workTimeExtent"`
	EmploymentType      string       `db:"employment_type" json:"employmentType"`
	Duration            string       `db:"duration" json:"duration"`
	Positions           int          `db:"positions" json:"positions"`
	Published           bool         `db:"published" json:"published"`
	OwnCar              bool         `db:"own_car" json:"ownCar"`
	RequiresExperience  bool         `db:"requires_experience" json:"requiresExperience"`
	Logotype            string       `db:"logotype" json:"logotype"`
	Company             Company      `json:"company"`
	Application         Application  `json:"application"`
	Workplace           Workplace    `json:"workplace"`
	DrivingLicenses     []string     `json:"drivingLicense"`
	Skills              []string     `json:"skills"`
	Languages           []Language   `json:"languages"`
	WorkExperiences     []Experience `json:"workExperiences"`
	Contacts            []Contact    `json:"contacts"`
	Keywords            []string     `json:"keywords"`
}

type Company struct {
	Name               string `db:"name" json:"name"`
	StreetAddress      string `db:"street_address" json:"streetAddress"`
	PostCode           string `db:"post_code" json:"postCode"`
	City               string `db:"city" json:"city"`
	PhoneNumber        string `db:"phone_number" json:"phoneNumber"`
	WebAddress         string `db:"web_address" json:"webAddress"`
	Email              string `db:"email" json:"email"`
	OrganisationNumber string `db:"organisation_number" json:"organisationNumber"`
}

type Application struct {
	Mail        string `db:"mail" json:"mail"`
	Email       string `db:"email" json:"email"`
	WebAddress  string `db:"web_address" json:"webAddress"`
	Other       string `db:"other" json:"other"`
	Reference   string `db:"reference" json:"reference"`
	Information string `db:"information" json:"information"`
}

type Workplace struct {
	Name                 string  `db:"name" json:"name"`
	Street               string  `db:"street" json:"street"`
	PostCode             string  `db:"post_code" json:"postCode"`
	City                 string  `db:"city" json:"city"`
	UnspecifiedWorkplace bool    `db:"unspecified_workplace" json:"unspecifiedWorkplace"`
	Region               string  `db:"region" json:"region"`
	Country              string  `db:"country" json:"country"`
	Municipality         string  `db:"municipality" json:"municipality"`
	Longitude            float64 `db:"longitude" json:"longitude"`
	Latitude             float64 `db:"latitude" json:"latitude"`
	ShowMap              bool    `db:"show_map" json:"showMap"`
}

type Language struct {
	Name     string `db:"name" json:"name"`
	Required bool   `db:"required" json:"required"`
}

type Experience struct {
	Name     string `db:"name" json:"name"`
	Required bool   `db:"required" json:"required"`
}

type Contact struct {
	Name string `db:"name" json:"name"`
}
