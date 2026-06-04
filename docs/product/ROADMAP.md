# Roadmap

## Phase 0: Foundation

Focus: repository setup, starter backend, starter frontend, shared conventions.

Deliverables:

- Go backend starter.
- React frontend starter (React 19, TanStack Router, Zustand, shadcn/ui, Tailwind CSS v4).
- Basic health check endpoint.
- Basic frontend page that can call the backend.
- Initial documentation.

## Phase 1: MVP Download Flow

Focus: single URL download flow.

Deliverables:

- URL probe endpoint.
- Job creation endpoint.
- Basic queue and worker.
- SQLite job persistence.
- `yt-dlp` process runner.
- Progress parsing.
- Job list and job detail UI.
- Completed file download endpoint.

## Phase 2: Reliability and Safety

Focus: making the app safer and more predictable.

Deliverables:

- Authentication.
- URL validation and SSRF protection.
- Download directory boundary checks.
- Worker cancellation.
- Retry failed jobs.
- Config validation.
- Structured logging.
- Error classification.

## Phase 3: Better UX

Focus: daily usability.

Deliverables:

- Better format preset UI.
- Settings page.
- Delete completed files.
- Download history filters.
- Improved progress display.
- Toasts and empty states.
- Mobile responsive polish.

## Phase 4: Advanced Downloads

Focus: useful `yt-dlp` features without overwhelming the product.

Deliverables:

- Subtitle options.
- Audio extraction presets.
- Thumbnail and metadata embedding.
- Custom output template with validation.
- Batch URLs.
- Playlist support with child jobs.

## Phase 5: Deployment and Operations

Focus: easier self-hosting.

Deliverables:

- Dockerfile.
- Docker Compose example.
- Health checks.
- Backup guide.
- Upgrade guide.
- Basic metrics endpoint.

## Phase 6: Long-Term Expansion

Possible future directions:

- Multi-user support.
- Role-based permissions.
- Storage provider abstraction.
- Remote workers.
- Public share links with expiry.
- Browser extension integration.
