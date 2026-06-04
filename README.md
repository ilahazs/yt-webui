# Youtube Downloader - Web UI

<img width="1903" height="1024" alt="image" src="https://github.com/user-attachments/assets/07b20150-f8e9-4662-81ec-ebb221e2aa87" />

aka yt-dlp webUI is a web interface for managing downloads through `yt-dlp`, built with a Go backend and a React frontend.

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

## Tech Stack

- Backend: Go
- Frontend: React 19, TanStack Router, Zustand, shadcn/ui, Tailwind CSS v4
- Package Manager: `pnpm`
- Downloader: `yt-dlp` CLI
- Media processing: `ffmpeg`
- Database: SQLite for the early phase
- Realtime updates: Server-Sent Events
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

To run the application locally for development:

### 1. Start the Go Backend Server

```bash
cd backend
go run main.go
```

The backend server will start listening at `http://localhost:8080`.

### 2. Start the React Frontend Server

In a new terminal window:

```bash
cd frontend
pnpm install
pnpm run dev
```

The Vite development server will start listening at `http://localhost:5173`. Open this URL in your browser to verify it is running and communicating with the backend.

### 3. Check Prerequisites

Ensure that `yt-dlp` and `ffmpeg` are installed and available on your system path.


## License

Choose a license before publishing the repository publicly.
