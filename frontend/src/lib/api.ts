import type {
  ApiResponse,
  HealthData,
  ProbeRequest,
  ProbeData,
  CreateJobRequest,
  CreateJobData,
  Job,
  JobListData,
} from './types';

const BASE_URL = '/api';

// ── Internal helpers ────────────────────────────────────────────

async function request<T>(
  input: RequestInfo,
  init?: RequestInit,
): Promise<ApiResponse<T>> {
  try {
    const res = await fetch(input, {
      headers: { 'Content-Type': 'application/json', ...init?.headers },
      ...init,
    });

    let json: any;
    try {
      json = await res.json();
    } catch (parseErr) {
      return {
        data: null,
        error: {
          code: 'parse_error',
          message: `Failed to parse response as JSON. Status: ${res.status} ${res.statusText}`,
        },
      };
    }

    if (json && typeof json === 'object' && ('data' in json || 'error' in json)) {
      return json as ApiResponse<T>;
    }

    return {
      data: null,
      error: {
        code: 'invalid_response_format',
        message: 'Server response does not match the API contract.',
      },
    };
  } catch (netErr: any) {
    return {
      data: null,
      error: {
        code: 'network_error',
        message: netErr instanceof Error ? netErr.message : 'A network error occurred.',
      },
    };
  }
}

function post<T>(path: string, body: unknown): Promise<ApiResponse<T>> {
  return request<T>(`${BASE_URL}${path}`, {
    method: 'POST',
    body: JSON.stringify(body),
  });
}

function get<T>(path: string): Promise<ApiResponse<T>> {
  return request<T>(`${BASE_URL}${path}`, { method: 'GET' });
}

// ── Public API client ───────────────────────────────────────────

export const api = {
  /** Check backend health. */
  health(): Promise<ApiResponse<HealthData>> {
    return get<HealthData>('/health');
  },

  /** Probe a URL for metadata before creating a download job. */
  probe(payload: ProbeRequest): Promise<ApiResponse<ProbeData>> {
    return post<ProbeData>('/probe', payload);
  },

  /** Create a new download job. */
  createJob(payload: CreateJobRequest): Promise<ApiResponse<CreateJobData>> {
    return post<CreateJobData>('/jobs', payload);
  },

  /** List all jobs (with optional status filter). */
  listJobs(): Promise<ApiResponse<JobListData>> {
    return get<JobListData>('/jobs');
  },

  /** Get a single job by ID. */
  getJob(jobId: string): Promise<ApiResponse<Job>> {
    return get<Job>(`/jobs/${jobId}`);
  },

  /** Cancel a running job. */
  cancelJob(jobId: string): Promise<ApiResponse<{ status: string }>> {
    return post<{ status: string }>(`/jobs/${jobId}/cancel`, {});
  },

  /**
   * Open an SSE stream for real-time job events.
   * Returns the raw EventSource — caller is responsible for closing it.
   */
  jobEvents(jobId: string): EventSource {
    return new EventSource(`${BASE_URL}/jobs/${jobId}/events`);
  },
};
