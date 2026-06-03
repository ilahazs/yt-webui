# Product Brief

## Product Summary

This project is a self-hosted web interface for `yt-dlp`. It helps users inspect media URLs, choose download options, queue downloads, monitor progress, and retrieve completed files through a browser.

## Target User

Primary target:

- A technical user running a private server or local machine.
- Someone who wants a cleaner interface than repeatedly typing `yt-dlp` commands.
- Someone comfortable with self-hosted tools but expecting a pleasant UI.

Secondary target, later:

- Small private groups with individual users, quotas, and permission rules.

## Core User Journey

1. User opens the web UI.
2. User pastes a media URL.
3. The app probes the URL and shows metadata.
4. User chooses a download preset.
5. The app creates a job.
6. User watches realtime progress.
7. User downloads or manages the completed file.

## MVP Features

- Paste URL and probe metadata.
- Display title, thumbnail, uploader, duration, and available formats where possible.
- Create download job from a preset.
- Process jobs through a backend worker.
- Show job status and progress.
- Store job history in a database.
- Provide access to completed files.
- Support basic private access control.

## Non-Goals for MVP

- Public SaaS usage.
- Multi-user administration.
- Playlist downloads.
- Browser extension.
- Mobile native app.
- Arbitrary command builder for every `yt-dlp` option.

## Product Personality

The UI should feel calm, clear, and utilitarian. Downloads can fail for many reasons, so the product should make failures understandable rather than mysterious.

## Success Criteria

The MVP is successful when a user can:

- Run the app locally or through Docker Compose.
- Paste a URL and see metadata.
- Start a download using a preset.
- Watch progress update without refreshing.
- Download the completed file.
- Inspect error output when a download fails.
