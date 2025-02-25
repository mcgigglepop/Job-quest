package main

import (
	"api/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *application) Health(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "api up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllJobs(w http.ResponseWriter, r *http.Request) {
	var jobs []models.Job
	ad, _ := time.Parse("2006-01-02", "2025-02-25")

	sampleJob := models.Job{
		ID:          1,
		UserID:      1,
		Company:     "Pokemon Company International",
		Position:    "Client Engineer",
		Location:    "Seattle, Washington",
		Remote: 		 "hybrd",
		Status:      "Applied",
		AppliedDate: ad,
		Notes:       "cool company to work for",
		ResumeURL:   "http://s3-url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Now(),
	}

	jobs = append(jobs, sampleJob)

	ad, _ = time.Parse("2006-01-02", "2025-01-24")

	anotherSampleJob := models.Job{
		ID:          2,
		UserID:      1,
		Company:     "Blizzard",
		Position:    "Software Engineer",
		Location:    "Irvine, California",
		Remote: 		 "no",
		Status:      "Round 2 Interview",
		AppliedDate: ad,
		Notes:       "another cool company to work for",
		ResumeURL:   "http://s3-url",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Now(),
	}

	jobs = append(jobs, anotherSampleJob)

	out, err := json.Marshal(jobs)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
