import { create } from 'zustand';
import { devtools } from 'zustand/middleware';
import type { Job, JobStatus } from '@/lib/types';
import { api } from '@/lib/api';

// ── State shape ─────────────────────────────────────────────────

interface JobStore {
  /** All known jobs, keyed by ID for O(1) lookup. */
  jobs: Record<string, Job>;

  /** Active SSE connections keyed by job ID. */
  eventSources: Record<string, EventSource>;

  /** Whether a list fetch is in flight. */
  isLoading: boolean;

  /** Last fetch error message, if any. */
  error: string | null;
}

// ── Actions ─────────────────────────────────────────────────────

interface JobActions {
  /** Fetch all jobs from the backend and hydrate the store. */
  fetchJobs: () => Promise<void>;

  /** Upsert a single job (used by SSE updates). */
  upsertJob: (job: Job) => void;

  /** Update a specific field on an existing job. */
  patchJob: (jobId: string, patch: Partial<Job>) => void;

  /**
   * Subscribe to SSE events for a job.
   * Idempotent — does nothing if already subscribed.
   */
  subscribe: (jobId: string) => void;

  /** Unsubscribe and close the SSE connection for a job. */
  unsubscribe: (jobId: string) => void;

  /** Unsubscribe from all active SSE connections. */
  unsubscribeAll: () => void;

  /** Clear any error state. */
  clearError: () => void;
}

// ── Store ───────────────────────────────────────────────────────

export const useJobStore = create<JobStore & JobActions>()(
  devtools(
    (set, get) => ({
      // Initial state
      jobs: {},
      eventSources: {},
      isLoading: false,
      error: null,

      // ── Actions ────────────────────────────────────────────

      fetchJobs: async () => {
        set({ isLoading: true, error: null });
        const res = await api.listJobs();

        if (res.error !== null) {
          set({ isLoading: false, error: res.error.message });
          return;
        }

        const jobMap: Record<string, Job> = {};
        for (const job of res.data.items) {
          jobMap[job.id] = job;
        }
        set({ jobs: jobMap, isLoading: false });
      },

      upsertJob: (job) => {
        set((state) => ({
          jobs: { ...state.jobs, [job.id]: job },
        }));
      },

      patchJob: (jobId, patch) => {
        set((state) => {
          const existing = state.jobs[jobId];
          if (!existing) return state;
          return {
            jobs: { ...state.jobs, [jobId]: { ...existing, ...patch } },
          };
        });
      },

      subscribe: (jobId) => {
        const { eventSources, patchJob } = get();

        // Already subscribed — skip
        if (eventSources[jobId]) return;

        const es = api.jobEvents(jobId);

        es.addEventListener('status', (e: MessageEvent) => {
          const { status } = JSON.parse(e.data) as { status: JobStatus };
          patchJob(jobId, { status });

          // Auto-unsubscribe on terminal states
          if (status === 'completed' || status === 'failed' || status === 'cancelled') {
            get().unsubscribe(jobId);
          }
        });

        es.addEventListener('progress', (e: MessageEvent) => {
          const { progress, speed, eta } = JSON.parse(e.data) as {
            progress: number;
            speed: string;
            eta: string;
          };
          patchJob(jobId, { progress, speed, eta });
        });

        es.onerror = () => {
          // Connection dropped — clean up gracefully
          get().unsubscribe(jobId);
        };

        set((state) => ({
          eventSources: { ...state.eventSources, [jobId]: es },
        }));
      },

      unsubscribe: (jobId) => {
        const { eventSources } = get();
        eventSources[jobId]?.close();
        set((state) => {
          const { [jobId]: _, ...rest } = state.eventSources;
          return { eventSources: rest };
        });
      },

      unsubscribeAll: () => {
        const { eventSources } = get();
        for (const es of Object.values(eventSources)) {
          es.close();
        }
        set({ eventSources: {} });
      },

      clearError: () => set({ error: null }),
    }),
    { name: 'JobStore' },
  ),
);
