/**
 * Shared TypeScript types that mirror the backend API contract.
 * See: docs/architecture/API_CONTRACT.md
 */

// ── Common ──────────────────────────────────────────────────────

export interface ApiSuccess<T> {
  data: T;
  error: null;
}

export interface ApiError {
  data: null;
  error: {
    code: string;
    message: string;
  };
}

export type ApiResponse<T> = ApiSuccess<T> | ApiError;

// ── Health ──────────────────────────────────────────────────────

export interface HealthData {
  status: 'ok' | string;
}

// ── Probe ───────────────────────────────────────────────────────

export interface ProbeRequest {
  url: string;
}

export interface FormatInfo {
  format_id: string;
  ext: string;
  resolution?: string;
  fps?: number;
  vcodec?: string;
  acodec?: string;
  filesize?: number;
  format_note?: string;
}

export interface ProbeData {
  title: string;
  webpage_url: string;
  thumbnail?: string;
  duration_seconds?: number;
  uploader?: string;
  formats: FormatInfo[];
}

// ── Jobs ────────────────────────────────────────────────────────

export type JobStatus =
  | 'pending'
  | 'running'
  | 'completed'
  | 'failed'
  | 'cancelled';

export interface JobOptions {
  embed_metadata?: boolean;
  write_subtitles?: boolean;
}

export type DownloadPreset = 'best_video' | 'best_audio' | 'best_720p' | 'best_480p';

export interface CreateJobRequest {
  url: string;
  preset: DownloadPreset;
  options?: JobOptions;
}

export interface CreateJobData {
  job_id: string;
}

export interface Job {
  id: string;
  status: JobStatus;
  progress?: number;
  title?: string;
  speed?: string;
  eta?: string;
  url: string;
  preset: DownloadPreset;
  created_at: string;
  started_at?: string;
  finished_at?: string;
  error?: string;
  file_id?: string;
}

export interface JobListData {
  items: Job[];
}

// ── SSE Events ──────────────────────────────────────────────────

export interface JobStatusEvent {
  status: JobStatus;
}

export interface JobProgressEvent {
  progress: number;
  speed: string;
  eta: string;
}

export interface JobLogEvent {
  line: string;
}
