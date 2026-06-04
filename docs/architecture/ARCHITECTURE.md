# Architecture

## Overview

The application is split into a Go backend and a SvelteKit frontend. The backend owns download execution, job state, validation, persistence, and file access. The frontend, styled with Tailwind CSS and utilizing the shadcn-svelte component library (Vega style), owns user interaction, job visualization, and settings screens. Package management for the frontend is managed by `pnpm`.

```txt
Browser
  |
  | HTTP / SSE
  v
Go Backend
  |
  | Job orchestration
  v
yt-dlp + ffmpeg
  |
  v
Local storage + database
```

## Backend Responsibilities

- Expose HTTP API.
- Validate URLs and request payloads.
- Probe metadata using `yt-dlp`.
- Create and persist jobs.
- Run workers with controlled concurrency.
- Execute `yt-dlp` without shell interpolation.
- Parse progress output.
- Broadcast job events.
- Serve completed files safely.
- Enforce authentication and access control.

## Frontend Responsibilities

- Provide URL input and metadata preview.
- Present safe download presets.
- Submit jobs to the backend.
- Display queue and history.
- Subscribe to realtime progress updates.
- Show errors in a readable form.
- Provide settings and administration screens over time.

## Suggested Backend Modules

These are recommendations, not strict requirements.

```txt
api         HTTP handlers and request/response mapping
auth        authentication and session/token handling
config      environment and config loading
db          persistence and migrations
downloader  yt-dlp process execution and progress parsing
events      SSE/WebSocket event broker
jobs        queue, worker, state transitions
storage     file path, output, and download serving
security    URL validation and filesystem safety helpers
```

## Suggested Frontend Areas

These are recommendations, not strict requirements.

```txt
routes      SvelteKit pages
lib/api     API client helpers
lib/types   shared frontend types
components  UI components
stores      job state, settings, and realtime subscriptions
```

## Data Flow: Probe

```txt
User submits URL
  -> Frontend calls POST /api/probe
  -> Backend validates URL
  -> Backend runs yt-dlp metadata command
  -> Backend normalizes useful fields
  -> Frontend displays preview and presets
```

## Data Flow: Download

```txt
User creates job
  -> Backend validates payload
  -> Backend stores pending job
  -> Worker picks job
  -> Worker starts yt-dlp
  -> Progress is parsed and persisted
  -> Events are streamed to frontend
  -> Final file path is stored
```

## Realtime Strategy

Use Server-Sent Events for early versions because progress updates are one-way from server to browser. WebSocket can be introduced later if the app needs bidirectional interactions.

## Persistence Strategy

Use SQLite early. It is enough for a self-hosted single-node app and keeps operations simple.

A future migration to PostgreSQL may be useful for multi-user or distributed worker scenarios.

## Process Execution Strategy

Use `exec.CommandContext` or an equivalent safe process API. Avoid shell string execution. Build argument arrays from validated presets and whitelisted options.

## Storage Strategy

All output files should live under a configured download directory. The app should never serve files outside that directory.
