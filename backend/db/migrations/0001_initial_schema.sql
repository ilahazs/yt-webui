-- Migration: 0001_initial_schema
-- Created At: 2026-06-08

CREATE TABLE jobs (
  id TEXT PRIMARY KEY,
  url TEXT NOT NULL,
  title TEXT,
  status TEXT NOT NULL,
  preset TEXT,
  progress REAL DEFAULT 0,
  speed TEXT,
  eta TEXT,
  output_path TEXT,
  error_code TEXT,
  error_message TEXT,
  created_at DATETIME NOT NULL,
  started_at DATETIME,
  finished_at DATETIME
);

CREATE TABLE job_events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  job_id TEXT NOT NULL,
  type TEXT NOT NULL,
  payload TEXT NOT NULL,
  created_at DATETIME NOT NULL,
  FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE
);

CREATE TABLE files (
  id TEXT PRIMARY KEY,
  job_id TEXT NOT NULL,
  path TEXT NOT NULL,
  filename TEXT NOT NULL,
  size_bytes INTEGER,
  mime_type TEXT,
  created_at DATETIME NOT NULL,
  FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE
);

CREATE TABLE settings (
  key TEXT PRIMARY KEY,
  value TEXT NOT NULL,
  updated_at DATETIME NOT NULL
);

-- Indexes for performance and quick lookups
CREATE INDEX idx_job_events_job_id ON job_events(job_id);
CREATE INDEX idx_files_job_id ON files(job_id);
CREATE INDEX idx_jobs_status ON jobs(status);
