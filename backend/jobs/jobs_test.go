package jobs

import (
	"database/sql"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"github.com/ilahazs/yt-webui/backend/config"
	"github.com/ilahazs/yt-webui/backend/db"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")

	cfg := &config.Config{
		DBPath: dbPath,
	}

	dbConn, err := db.Init(cfg)
	if err != nil {
		t.Fatalf("failed to initialize test database: %v", err)
	}

	return dbConn
}

func TestCreateAndGetJob(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	// 1. Get non-existent job
	_, err := GetJob(dbConn, "non-existent")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound for non-existent job, got %v", err)
	}

	// 2. Create validation failures
	err = CreateJob(dbConn, &Job{ID: "", URL: "https://example.com"})
	if err == nil {
		t.Error("expected error when creating job with empty ID")
	}

	err = CreateJob(dbConn, &Job{ID: "job-1", URL: ""})
	if err == nil {
		t.Error("expected error when creating job with empty URL")
	}

	// 3. Create a valid job with minimal fields
	title := "Test Video"
	preset := "best"
	now := time.Now().UTC().Truncate(time.Second)

	job1 := &Job{
		ID:        "job-1",
		URL:       "https://example.com/watch?v=1",
		Title:     &title,
		Status:    StatusPending,
		Preset:    &preset,
		CreatedAt: now,
	}

	err = CreateJob(dbConn, job1)
	if err != nil {
		t.Fatalf("failed to create job: %v", err)
	}

	// 4. Retrieve and verify fields
	fetched, err := GetJob(dbConn, "job-1")
	if err != nil {
		t.Fatalf("failed to get job: %v", err)
	}

	if fetched.ID != job1.ID {
		t.Errorf("expected ID %q, got %q", job1.ID, fetched.ID)
	}
	if fetched.URL != job1.URL {
		t.Errorf("expected URL %q, got %q", job1.URL, fetched.URL)
	}
	if fetched.Title == nil || *fetched.Title != *job1.Title {
		t.Errorf("expected Title %q, got %v", *job1.Title, fetched.Title)
	}
	if fetched.Status != job1.Status {
		t.Errorf("expected Status %q, got %q", job1.Status, fetched.Status)
	}
	if fetched.Preset == nil || *fetched.Preset != *job1.Preset {
		t.Errorf("expected Preset %q, got %v", *job1.Preset, fetched.Preset)
	}
	if fetched.Progress != 0 {
		t.Errorf("expected default Progress 0, got %f", fetched.Progress)
	}
	if !fetched.CreatedAt.Equal(job1.CreatedAt) && fetched.CreatedAt.Unix() != job1.CreatedAt.Unix() {
		t.Errorf("expected CreatedAt %v, got %v", job1.CreatedAt, fetched.CreatedAt)
	}

	// Null fields check
	if fetched.Speed != nil {
		t.Errorf("expected Speed to be nil, got %v", fetched.Speed)
	}
	if fetched.ETA != nil {
		t.Errorf("expected ETA to be nil, got %v", fetched.ETA)
	}
	if fetched.OutputPath != nil {
		t.Errorf("expected OutputPath to be nil, got %v", fetched.OutputPath)
	}
	if fetched.ErrorCode != nil {
		t.Errorf("expected ErrorCode to be nil, got %v", fetched.ErrorCode)
	}
	if fetched.ErrorMessage != nil {
		t.Errorf("expected ErrorMessage to be nil, got %v", fetched.ErrorMessage)
	}
	if fetched.StartedAt != nil {
		t.Errorf("expected StartedAt to be nil, got %v", fetched.StartedAt)
	}
	if fetched.FinishedAt != nil {
		t.Errorf("expected FinishedAt to be nil, got %v", fetched.FinishedAt)
	}
}

func TestUpdateJobStatus(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	job := &Job{
		ID:        "job-update-status",
		URL:       "https://example.com/watch?v=update",
		Status:    StatusPending,
		CreatedAt: time.Now().UTC().Truncate(time.Second),
	}
	if err := CreateJob(dbConn, job); err != nil {
		t.Fatalf("failed to create job: %v", err)
	}

	// 1. Update status to running with StartedAt
	startedAt := time.Now().UTC().Truncate(time.Second)
	err := UpdateJobStatus(dbConn, job.ID, StatusRunning, &startedAt, nil, nil, nil)
	if err != nil {
		t.Fatalf("failed to update status to running: %v", err)
	}

	fetched, err := GetJob(dbConn, job.ID)
	if err != nil {
		t.Fatalf("failed to get job: %v", err)
	}
	if fetched.Status != StatusRunning {
		t.Errorf("expected status to be %q, got %q", StatusRunning, fetched.Status)
	}
	if fetched.StartedAt == nil || fetched.StartedAt.Unix() != startedAt.Unix() {
		t.Errorf("expected StartedAt %v, got %v", startedAt, fetched.StartedAt)
	}
	if fetched.FinishedAt != nil {
		t.Errorf("expected FinishedAt to be nil, got %v", fetched.FinishedAt)
	}

	// 2. Update status to completed with FinishedAt
	finishedAt := time.Now().UTC().Truncate(time.Second)
	err = UpdateJobStatus(dbConn, job.ID, StatusCompleted, &startedAt, &finishedAt, nil, nil)
	if err != nil {
		t.Fatalf("failed to update status to completed: %v", err)
	}

	fetched, err = GetJob(dbConn, job.ID)
	if err != nil {
		t.Fatalf("failed to get job: %v", err)
	}
	if fetched.Status != StatusCompleted {
		t.Errorf("expected status to be %q, got %q", StatusCompleted, fetched.Status)
	}
	if fetched.FinishedAt == nil || fetched.FinishedAt.Unix() != finishedAt.Unix() {
		t.Errorf("expected FinishedAt %v, got %v", finishedAt, fetched.FinishedAt)
	}

	// 3. Update status to failed with errors
	errCode := "download_failed"
	errMsg := "yt-dlp exited with code 1"
	err = UpdateJobStatus(dbConn, job.ID, StatusFailed, &startedAt, &finishedAt, &errCode, &errMsg)
	if err != nil {
		t.Fatalf("failed to update status to failed: %v", err)
	}

	fetched, err = GetJob(dbConn, job.ID)
	if err != nil {
		t.Fatalf("failed to get job: %v", err)
	}
	if fetched.Status != StatusFailed {
		t.Errorf("expected status to be %q, got %q", StatusFailed, fetched.Status)
	}
	if fetched.ErrorCode == nil || *fetched.ErrorCode != errCode {
		t.Errorf("expected error code %q, got %v", errCode, fetched.ErrorCode)
	}
	if fetched.ErrorMessage == nil || *fetched.ErrorMessage != errMsg {
		t.Errorf("expected error message %q, got %v", errMsg, fetched.ErrorMessage)
	}

	// 4. Update non-existent job
	err = UpdateJobStatus(dbConn, "non-existent", StatusRunning, nil, nil, nil, nil)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdateJobProgress(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	job := &Job{
		ID:        "job-update-progress",
		URL:       "https://example.com/watch?v=progress",
		Status:    StatusPending,
		CreatedAt: time.Now().UTC().Truncate(time.Second),
	}
	if err := CreateJob(dbConn, job); err != nil {
		t.Fatalf("failed to create job: %v", err)
	}

	// 1. Update progress
	speed := "1.5MiB/s"
	eta := "00:30"
	err := UpdateJobProgress(dbConn, job.ID, 45.5, &speed, &eta)
	if err != nil {
		t.Fatalf("failed to update progress: %v", err)
	}

	fetched, err := GetJob(dbConn, job.ID)
	if err != nil {
		t.Fatalf("failed to get job: %v", err)
	}
	if fetched.Progress != 45.5 {
		t.Errorf("expected progress 45.5, got %f", fetched.Progress)
	}
	if fetched.Speed == nil || *fetched.Speed != speed {
		t.Errorf("expected speed %q, got %v", speed, fetched.Speed)
	}
	if fetched.ETA == nil || *fetched.ETA != eta {
		t.Errorf("expected eta %q, got %v", eta, fetched.ETA)
	}

	// 2. Update non-existent job progress
	err = UpdateJobProgress(dbConn, "non-existent", 50.0, nil, nil)
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestListJobs(t *testing.T) {
	dbConn := setupTestDB(t)
	defer dbConn.Close()

	baseTime := time.Now().UTC().Truncate(time.Second)

	jobA := &Job{
		ID:        "job-a",
		URL:       "https://example.com/a",
		Status:    StatusPending,
		CreatedAt: baseTime.Add(-10 * time.Minute),
	}
	jobB := &Job{
		ID:        "job-b",
		URL:       "https://example.com/b",
		Status:    StatusPending,
		CreatedAt: baseTime.Add(-5 * time.Minute),
	}
	jobC := &Job{
		ID:        "job-c",
		URL:       "https://example.com/c",
		Status:    StatusPending,
		CreatedAt: baseTime,
	}
	// jobD has the same CreatedAt as jobC to test secondary sorting by ID DESC
	jobD := &Job{
		ID:        "job-d",
		URL:       "https://example.com/d",
		Status:    StatusPending,
		CreatedAt: baseTime,
	}

	if err := CreateJob(dbConn, jobA); err != nil {
		t.Fatalf("failed to create jobA: %v", err)
	}
	if err := CreateJob(dbConn, jobB); err != nil {
		t.Fatalf("failed to create jobB: %v", err)
	}
	if err := CreateJob(dbConn, jobC); err != nil {
		t.Fatalf("failed to create jobC: %v", err)
	}
	if err := CreateJob(dbConn, jobD); err != nil {
		t.Fatalf("failed to create jobD: %v", err)
	}

	// List jobs and check deterministic ordering:
	// 1. CreatedAt DESC: jobC and jobD should come before jobB, and jobB before jobA.
	// 2. ID DESC (secondary): between jobC and jobD (which have same CreatedAt), jobD should come before jobC.
	// So expected order is: jobD, jobC, jobB, jobA.
	list, err := ListJobs(dbConn)
	if err != nil {
		t.Fatalf("failed to list jobs: %v", err)
	}

	if len(list) != 4 {
		t.Fatalf("expected 4 jobs in list, got %d", len(list))
	}

	expectedIDs := []string{"job-d", "job-c", "job-b", "job-a"}
	for i, expectedID := range expectedIDs {
		if list[i].ID != expectedID {
			t.Errorf("at index %d, expected job ID %q, got %q", i, expectedID, list[i].ID)
		}
	}
}
