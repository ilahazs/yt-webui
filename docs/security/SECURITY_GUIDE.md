# Security Guide

## Security Posture

This project should be treated as a private self-hosted application by default. A public internet deployment requires additional hardening, abuse protection, and legal review.

## Main Risks

- SSRF through user-provided URLs.
- Command injection through unsafe process execution.
- Path traversal through output paths or filenames.
- Resource exhaustion through large downloads or many jobs.
- Sensitive cookie leakage.
- Exposing downloaded files without authorization.

## URL Validation

Recommended rules:

- Allow only `http` and `https` URLs.
- Reject `file://`, `ftp://`, `data:`, and other schemes.
- Resolve hostnames and block private or local ranges.
- Block localhost names and loopback IPs.
- Consider blocking link-local and cloud metadata IP ranges.
- Apply request size and timeout limits.

## Process Execution

Recommended rules:

- Use process argument arrays, not shell string execution.
- Build arguments from whitelisted options.
- Do not allow arbitrary raw `yt-dlp` arguments in the default UI.
- Use context cancellation for stopped jobs.
- Set timeouts where appropriate.
- Capture stdout/stderr safely.

## Filesystem Safety

Recommended rules:

- Keep all downloads inside a configured download directory.
- Normalize and verify paths before serving files.
- Reject paths that escape the download directory.
- Generate application-owned file IDs.
- Avoid serving arbitrary filesystem paths from request parameters.

## Authentication

MVP recommendation:

- Single-user password or token-based access.
- Secure cookie settings when using browser sessions.
- Optional reverse-proxy auth compatibility.

Later recommendations:

- Multi-user accounts.
- Role-based permissions.
- Per-user quotas.
- Audit logging.

## Rate Limits and Quotas

Recommended controls:

- Maximum concurrent workers.
- Maximum queued jobs.
- Maximum playlist items when playlist support exists.
- Maximum log lines retained per job.
- Optional max duration or file size policy.

## Cookie Handling

Cookie support is useful but sensitive.

Recommended rules:

- Do not add cookie upload in the earliest MVP unless needed.
- Prefer server-side configured cookie profile paths.
- Do not expose cookie content in logs or API responses.
- Restrict filesystem permissions.

## Public Deployment Warning

Before exposing the app publicly, add stronger authentication, rate limits, job quotas, abuse monitoring, and careful policy controls around supported URLs.
