# API Contract Draft

This contract is a starting point. Names and shapes may change during implementation.

## Common Response Shape

Recommended success response:

```json
{
  "data": {},
  "error": null
}
```

Recommended error response:

```json
{
  "data": null,
  "error": {
    "code": "invalid_request",
    "message": "Human readable message"
  }
}
```

## Health

```http
GET /api/health
```

Response:

```json
{
  "data": {
    "status": "ok"
  },
  "error": null
}
```

## Probe URL

```http
POST /api/probe
```

Request:

```json
{
  "url": "https://example.com/watch?v=123"
}
```

Response:

```json
{
  "data": {
    "title": "Example title",
    "webpage_url": "https://example.com/watch?v=123",
    "thumbnail": "https://example.com/thumb.jpg",
    "duration_seconds": 123,
    "uploader": "Example uploader",
    "formats": []
  },
  "error": null
}
```

## Create Job

```http
POST /api/jobs
```

Request:

```json
{
  "url": "https://example.com/watch?v=123",
  "preset": "best_video",
  "options": {
    "embed_metadata": true,
    "write_subtitles": false
  }
}
```

Response:

```json
{
  "data": {
    "job_id": "job_123"
  },
  "error": null
}
```

## List Jobs

```http
GET /api/jobs
```

Query parameters may later include:

- `status`
- `limit`
- `cursor`
- `sort`

Response:

```json
{
  "data": {
    "items": []
  },
  "error": null
}
```

## Get Job

```http
GET /api/jobs/{job_id}
```

Response:

```json
{
  "data": {
    "id": "job_123",
    "status": "running",
    "progress": 42.5,
    "title": "Example title",
    "speed": "1.2MiB/s",
    "eta": "00:01",
    "created_at": "2026-01-01T00:00:00Z",
    "started_at": "2026-01-01T00:00:01Z",
    "finished_at": null,
    "error": null
  },
  "error": null
}
```

## Job Events

```http
GET /api/jobs/{job_id}/events
```

Suggested SSE events:

```txt
event: status
data: {"status":"running"}

event: progress
data: {"progress":42.5,"speed":"1.2MiB/s","eta":"00:01"}

event: log
data: {"line":"download output line"}
```

## Cancel Job

```http
POST /api/jobs/{job_id}/cancel
```

Response:

```json
{
  "data": {
    "status": "cancelled"
  },
  "error": null
}
```

## Download Completed File

```http
GET /api/files/{file_id}/download
```

This endpoint should only serve files known to the application and located inside the configured download directory.
