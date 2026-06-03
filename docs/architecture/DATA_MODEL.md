# Data Model Draft

This is a general starting point for SQLite or a similar relational database.

## jobs

Stores one download request.

Recommended fields:

```sql
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
```

Recommended statuses:

- `pending`
- `running`
- `completed`
- `failed`
- `cancelled`

## job_events

Stores important state changes, logs, and progress snapshots.

```sql
CREATE TABLE job_events (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  job_id TEXT NOT NULL,
  type TEXT NOT NULL,
  payload TEXT NOT NULL,
  created_at DATETIME NOT NULL
);
```

Recommended event types:

- `status`
- `progress`
- `log`
- `error`

## files

Stores completed output files.

```sql
CREATE TABLE files (
  id TEXT PRIMARY KEY,
  job_id TEXT NOT NULL,
  path TEXT NOT NULL,
  filename TEXT NOT NULL,
  size_bytes INTEGER,
  mime_type TEXT,
  created_at DATETIME NOT NULL
);
```

## settings

Stores simple runtime settings where appropriate.

```sql
CREATE TABLE settings (
  key TEXT PRIMARY KEY,
  value TEXT NOT NULL,
  updated_at DATETIME NOT NULL
);
```

## playlist_items

Not required for MVP. Useful later when playlist support is added.

```sql
CREATE TABLE playlist_items (
  id TEXT PRIMARY KEY,
  parent_job_id TEXT NOT NULL,
  child_job_id TEXT,
  url TEXT NOT NULL,
  title TEXT,
  index_no INTEGER,
  status TEXT NOT NULL
);
```

## Notes

- Prefer immutable event records for debugging.
- Keep frequently read job state on the `jobs` table.
- Store paths carefully and never trust them directly for file serving.
- Use migrations from the beginning, even if the schema is small.
