# yt-dlp Web UI

A long-term web interface for managing downloads through `yt-dlp`, built with a Go backend and a SvelteKit frontend.

The project aims to provide a clean self-hosted experience for probing media URLs, choosing download presets, tracking download progress, and managing completed files from a browser.

## Goals

- Provide a simple web interface for `yt-dlp` workflows.
- Keep the backend reliable, observable, and safe to run on a private server.
- Keep the frontend fast, understandable, and comfortable for daily use.
- Start small with a strong foundation that can grow over time.

## Intended Scope

Initial scope:

- URL probing and metadata preview.
- Download job creation.
- Basic queue and worker flow.
- Realtime progress updates.
- Completed file access.
- Basic configuration.
- Private self-hosted deployment.

Later scope:

- Playlist support.
- Batch download support.
- Cookie profile support.
- Download history and archive.
- Multi-user support.
- Advanced storage backends.
- More detailed metrics and administration tools.

## Suggested Tech Stack

- Backend: Go
- Frontend: SvelteKit
- Downloader: `yt-dlp` CLI
- Media processing: `ffmpeg`
- Database: SQLite for the early phase
- Realtime updates: Server-Sent Events or WebSocket
- Deployment: Docker Compose for local/self-hosted usage

## Development Principles

- Prefer simple components before distributed systems.
- Avoid exposing arbitrary `yt-dlp` arguments directly to users.
- Keep download output inside a controlled directory.
- Treat URLs, filenames, cookies, and shell arguments as untrusted input.
- Make progress visible and failures easy to inspect.
- Keep public documentation user-focused and implementation-friendly.

## Repository Layout

The exact layout may evolve. A reasonable starting point:

```txt
.
├── backend/
├── frontend/
├── docs/
├── docker-compose.yml
└── README.md
```

## Local Development

Recommended early workflow:

1. Start the backend development server.
2. Start the SvelteKit development server.
3. Configure the frontend API base URL.
4. Confirm `yt-dlp` and `ffmpeg` are available in the backend runtime.
5. Test with a safe public media URL.

## License

Choose a license before publishing the repository publicly.
