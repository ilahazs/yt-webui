# Deployment Notes

## Recommended Early Deployment

Use Docker Compose for early self-hosted usage.

Recommended mounted directories:

- `/data` for the database and application state.
- `/downloads` for output files.

## Runtime Dependencies

The backend runtime should include:

- Go backend binary.
- `yt-dlp`.
- `ffmpeg`.
- CA certificates.
- A writable data directory.
- A writable download directory.

## Example Environment Variables

```env
APP_ENV=development
APP_BIND_ADDR=:8080
APP_DB_PATH=/data/app.db
APP_DOWNLOAD_DIR=/downloads
APP_MAX_WORKERS=2
APP_AUTH_TOKEN=change-me
APP_YTDLP_PATH=yt-dlp
APP_FFMPEG_PATH=ffmpeg
```

## Health Check

Recommended endpoint:

```http
GET /api/health
```

Recommended checks over time:

- Backend process is alive.
- Database is reachable.
- Download directory is writable.
- `yt-dlp` is available.
- `ffmpeg` is available.

## Backup

Back up:

- SQLite database.
- Configuration.
- Any important downloaded files.

Downloads can become large, so backup policy should be explicit.

## Upgrade Notes

- Run database migrations before serving traffic.
- Keep a backup before migration.
- Avoid changing output path conventions without a migration strategy.
