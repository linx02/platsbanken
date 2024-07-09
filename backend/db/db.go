// FIXME
// Dates are not being parsed/saved correctly
// Salarydescription is not

package db

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB(databaseName string) error {
	var err error
	once.Do(func() {
		db, err = sql.Open("sqlite3", databaseName)
		if err != nil {
			err = fmt.Errorf("error opening database: %v", err)
			return
		}

		// Create tables if they don't exist
		err = createTables()
		if err != nil {
			err = fmt.Errorf("error creating tables: %v", err)
			return
		}
	})

	return err
}

func createTables() error {
	// Create tables if they don't exist
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS job_postings (
            id STRING NOT NULL,
            title TEXT NOT NULL,
            description TEXT,
            published_date TEXT,
            last_application_date TEXT,
            occupation TEXT,
            conditions TEXT,
            salary_description TEXT,
            salary_type TEXT,
            work_time_extent TEXT,
            employment_type TEXT,
            duration TEXT,
            positions INTEGER,
            published INTEGER,
            own_car INTEGER,
            requires_experience INTEGER,
            logotype TEXT,
            company_name TEXT,
            company_street_address TEXT,
            company_post_code TEXT,
            company_city TEXT,
            company_phone_number TEXT,
            company_web_address TEXT,
            company_email TEXT,
            company_organisation_number TEXT,
            application_mail TEXT,
            application_email TEXT,
            application_web_address TEXT,
            application_other TEXT,
            application_reference TEXT,
            application_information TEXT,
            workplace_name TEXT,
            workplace_street TEXT,
            workplace_post_code TEXT,
            workplace_city TEXT,
            workplace_unspecified_workplace INTEGER,
            workplace_region TEXT,
            workplace_country TEXT,
            workplace_municipality TEXT,
            workplace_longitude REAL,
            workplace_latitude REAL,
            workplace_show_map INTEGER
        );
    `)
	if err != nil {
		return fmt.Errorf("error creating job_postings table: %v", err)
	}
	return nil
}

func CloseDB() {
	db.Close()
}

func GetJob(id string) (*JobPosting, error) {
	var job JobPosting
	query := `
        SELECT 
            id, title, description, published_date, last_application_date,
            occupation, conditions, salary_description, salary_type,
            work_time_extent, employment_type, duration, positions,
            published, own_car, requires_experience, logotype,
            company_name, company_street_address, company_post_code,
            company_city, company_phone_number, company_web_address,
            company_email, company_organisation_number,
            application_mail, application_email, application_web_address,
            application_other, application_reference, application_information,
            workplace_name, workplace_street, workplace_post_code,
            workplace_city, workplace_unspecified_workplace,
            workplace_region, workplace_country, workplace_municipality,
            workplace_longitude, workplace_latitude, workplace_show_map
        FROM job_postings
        WHERE id = ?
    `
	var publishedDateStr, lastApplicationDateStr string // Declare variables to hold dates as strings
	err := db.QueryRow(query, id).Scan(
		&job.ID, &job.Title, &job.Description, &publishedDateStr, &lastApplicationDateStr,
		&job.Occupation, &job.Conditions, &job.SalaryDescription, &job.SalaryType,
		&job.WorkTimeExtent, &job.EmploymentType, &job.Duration, &job.Positions,
		&job.Published, &job.OwnCar, &job.RequiresExperience, &job.Logotype,
		&job.Company.Name, &job.Company.StreetAddress, &job.Company.PostCode,
		&job.Company.City, &job.Company.PhoneNumber, &job.Company.WebAddress,
		&job.Company.Email, &job.Company.OrganisationNumber,
		&job.Application.Mail, &job.Application.Email, &job.Application.WebAddress,
		&job.Application.Other, &job.Application.Reference, &job.Application.Information,
		&job.Workplace.Name, &job.Workplace.Street, &job.Workplace.PostCode,
		&job.Workplace.City, &job.Workplace.UnspecifiedWorkplace,
		&job.Workplace.Region, &job.Workplace.Country, &job.Workplace.Municipality,
		&job.Workplace.Longitude, &job.Workplace.Latitude, &job.Workplace.ShowMap,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("job posting not found")
		}
		return nil, fmt.Errorf("error fetching job posting: %v", err)
	}

	// Parse published_date and last_application_date strings into time.Time
	job.PublishedDate, err = time.Parse("2006-01-02 15:04:05-07:00", publishedDateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing published_date: %v", err)
	}

	job.LastApplicationDate, err = time.Parse("2006-01-02 15:04:05-07:00", lastApplicationDateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing last_application_date: %v", err)
	}

	return &job, nil
}

func GetAllJobs() ([]*JobPosting, error) {
	var jobs []*JobPosting
	query := `
        SELECT 
            id, title, description, published_date, last_application_date,
            occupation, conditions, salary_description, salary_type,
            work_time_extent, employment_type, duration, positions,
            published, own_car, requires_experience, logotype,
            company_name, company_street_address, company_post_code,
            company_city, company_phone_number, company_web_address,
            company_email, company_organisation_number,
            application_mail, application_email, application_web_address,
            application_other, application_reference, application_information,
            workplace_name, workplace_street, workplace_post_code,
            workplace_city, workplace_unspecified_workplace,
            workplace_region, workplace_country, workplace_municipality,
            workplace_longitude, workplace_latitude, workplace_show_map
        FROM job_postings
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching job postings: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var job JobPosting
		var publishedDateStr, lastApplicationDateStr string // Declare variables to hold dates as strings
		err := rows.Scan(
			&job.ID, &job.Title, &job.Description, &publishedDateStr, &lastApplicationDateStr,
			&job.Occupation, &job.Conditions, &job.SalaryDescription, &job.SalaryType,
			&job.WorkTimeExtent, &job.EmploymentType, &job.Duration, &job.Positions,
			&job.Published, &job.OwnCar, &job.RequiresExperience, &job.Logotype,
			&job.Company.Name, &job.Company.StreetAddress, &job.Company.PostCode,
			&job.Company.City, &job.Company.PhoneNumber, &job.Company.WebAddress,
			&job.Company.Email, &job.Company.OrganisationNumber,
			&job.Application.Mail, &job.Application.Email, &job.Application.WebAddress,
			&job.Application.Other, &job.Application.Reference, &job.Application.Information,
			&job.Workplace.Name, &job.Workplace.Street, &job.Workplace.PostCode,
			&job.Workplace.City, &job.Workplace.UnspecifiedWorkplace,
			&job.Workplace.Region, &job.Workplace.Country, &job.Workplace.Municipality,
			&job.Workplace.Longitude, &job.Workplace.Latitude, &job.Workplace.ShowMap,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning job posting: %v", err)
		}

		// Parse published_date and last_application_date strings into time.Time
		job.PublishedDate, err = time.Parse("2006-01-02 15:04:05-07:00", publishedDateStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing published_date: %v", err)
		}

		job.LastApplicationDate, err = time.Parse("2006-01-02 15:04:05-07:00", lastApplicationDateStr)
		if err != nil {
			return nil, fmt.Errorf("error parsing last_application_date: %v", err)
		}

		jobs = append(jobs, &job)
	}

	return jobs, nil
}

func AddJob(job *JobPosting) error {
	query := `
        INSERT INTO job_postings (
            id, title, description, published_date, last_application_date,
            occupation, conditions, salary_description, salary_type,
            work_time_extent, employment_type, duration, positions,
            published, own_car, requires_experience, logotype,
            company_name, company_street_address, company_post_code,
            company_city, company_phone_number, company_web_address,
            company_email, company_organisation_number,
            application_mail, application_email, application_web_address,
            application_other, application_reference, application_information,
            workplace_name, workplace_street, workplace_post_code,
            workplace_city, workplace_unspecified_workplace,
            workplace_region, workplace_country, workplace_municipality,
            workplace_longitude, workplace_latitude, workplace_show_map
        )
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	_, err := db.Exec(
		query,
		job.ID, job.Title, job.Description, job.PublishedDate, job.LastApplicationDate,
		job.Occupation, job.Conditions, job.SalaryDescription, job.SalaryType,
		job.WorkTimeExtent, job.EmploymentType, job.Duration, job.Positions,
		job.Published, job.OwnCar, job.RequiresExperience, job.Logotype,
		job.Company.Name, job.Company.StreetAddress, job.Company.PostCode,
		job.Company.City, job.Company.PhoneNumber, job.Company.WebAddress,
		job.Company.Email, job.Company.OrganisationNumber,
		job.Application.Mail, job.Application.Email, job.Application.WebAddress,
		job.Application.Other, job.Application.Reference, job.Application.Information,
		job.Workplace.Name, job.Workplace.Street, job.Workplace.PostCode,
		job.Workplace.City, job.Workplace.UnspecifiedWorkplace,
		job.Workplace.Region, job.Workplace.Country, job.Workplace.Municipality,
		job.Workplace.Longitude, job.Workplace.Latitude, job.Workplace.ShowMap,
	)
	if err != nil {
		return fmt.Errorf("error adding job posting: %v", err)
	}
	return nil
}

func GetTotalJobs() (int, error) {
	var total int
	query := "SELECT COUNT(*) FROM job_postings"
	err := db.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error fetching total jobs: %v", err)
	}
	return total, nil
}

func GetDB() *sql.DB {
	return db
}
