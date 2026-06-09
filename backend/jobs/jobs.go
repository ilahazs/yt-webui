package jobs

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// JobStatus represents the state of a download job.
type JobStatus string

const (
	StatusPending   JobStatus = "pending"
	StatusRunning   JobStatus = "running"
	StatusCompleted JobStatus = "completed"
	StatusFailed    JobStatus = "failed"
	StatusCancelled JobStatus = "cancelled"
)

// Job represents a single download request and its current state.
type Job struct {
	ID           string     `json:"id"`
	URL          string     `json:"url"`
	Title        *string    `json:"title"`
	Status       JobStatus  `json:"status"`
	Preset       *string    `json:"preset"`
	Progress     float64    `json:"progress"`
	Speed        *string    `json:"speed"`
	ETA          *string    `json:"eta"`
	OutputPath   *string    `json:"output_path"`
	ErrorCode    *string    `json:"error_code"`
	ErrorMessage *string    `json:"error_message"`
	CreatedAt    time.Time  `json:"created_at"`
	StartedAt    *time.Time `json:"started_at"`
	FinishedAt   *time.Time `json:"finished_at"`
}

// ErrNotFound is returned when a job cannot be found in the database.
var ErrNotFound = errors.New("job not found")

// CreateJob inserts a new job record into the database.
func CreateJob(db *sql.DB, job *Job) error {
	if job.ID == "" {
		return fmt.Errorf("job ID is required")
	}
	if job.URL == "" {
		return fmt.Errorf("job URL is required")
	}
	if job.Status == "" {
		job.Status = StatusPending
	}
	if job.CreatedAt.IsZero() {
		job.CreatedAt = time.Now()
	}

	query := `
		INSERT INTO jobs (
			id, url, title, status, preset, progress, speed, eta, 
			output_path, error_code, error_message, created_at, started_at, finished_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(
		query,
		job.ID,
		job.URL,
		job.Title,
		job.Status,
		job.Preset,
		job.Progress,
		job.Speed,
		job.ETA,
		job.OutputPath,
		job.ErrorCode,
		job.ErrorMessage,
		job.CreatedAt,
		job.StartedAt,
		job.FinishedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	return nil
}

// GetJob retrieves a job record by its ID.
func GetJob(db *sql.DB, id string) (*Job, error) {
	query := `
		SELECT 
			id, url, title, status, preset, progress, speed, eta, 
			output_path, error_code, error_message, created_at, started_at, finished_at
		FROM jobs
		WHERE id = ?
	`

	var job Job
	err := db.QueryRow(query, id).Scan(
		&job.ID,
		&job.URL,
		&job.Title,
		&job.Status,
		&job.Preset,
		&job.Progress,
		&job.Speed,
		&job.ETA,
		&job.OutputPath,
		&job.ErrorCode,
		&job.ErrorMessage,
		&job.CreatedAt,
		&job.StartedAt,
		&job.FinishedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed to get job %s: %w", id, err)
	}

	return &job, nil
}

// ListJobs retrieves all jobs ordered deterministically by CreatedAt DESC, then ID DESC.
func ListJobs(db *sql.DB) ([]*Job, error) {
	query := `
		SELECT 
			id, url, title, status, preset, progress, speed, eta, 
			output_path, error_code, error_message, created_at, started_at, finished_at
		FROM jobs
		ORDER BY created_at DESC, id DESC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list jobs: %w", err)
	}
	defer rows.Close()

	var jobs []*Job
	for rows.Next() {
		var job Job
		err := rows.Scan(
			&job.ID,
			&job.URL,
			&job.Title,
			&job.Status,
			&job.Preset,
			&job.Progress,
			&job.Speed,
			&job.ETA,
			&job.OutputPath,
			&job.ErrorCode,
			&job.ErrorMessage,
			&job.CreatedAt,
			&job.StartedAt,
			&job.FinishedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan job: %w", err)
		}
		jobs = append(jobs, &job)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error listing jobs: %w", err)
	}

	return jobs, nil
}

// UpdateJobStatus updates the status and related metadata/timestamps of a job.
func UpdateJobStatus(
	db *sql.DB,
	id string,
	status JobStatus,
	startedAt *time.Time,
	finishedAt *time.Time,
	errCode *string,
	errMsg *string,
) error {
	query := `
		UPDATE jobs
		SET status = ?, started_at = ?, finished_at = ?, error_code = ?, error_message = ?
		WHERE id = ?
	`

	res, err := db.Exec(query, status, startedAt, finishedAt, errCode, errMsg, id)
	if err != nil {
		return fmt.Errorf("failed to update job status: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

// UpdateJobProgress updates the dynamic download progress, speed, and ETA of a job.
func UpdateJobProgress(db *sql.DB, id string, progress float64, speed *string, eta *string) error {
	query := `
		UPDATE jobs
		SET progress = ?, speed = ?, eta = ?
		WHERE id = ?
	`

	res, err := db.Exec(query, progress, speed, eta, id)
	if err != nil {
		return fmt.Errorf("failed to update job progress: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
