import { useEffect } from 'react';
import { useJobStore } from '@/store/useJobStore';

/**
 * Subscribes to SSE job events for the given job ID and
 * automatically unsubscribes when the component unmounts.
 *
 * Usage:
 * ```tsx
 * useJobEvents(jobId);
 * // Job state updates automatically via the Zustand store
 * const job = useJobStore((s) => s.jobs[jobId]);
 * ```
 */
export function useJobEvents(jobId: string | undefined): void {
  const subscribe = useJobStore((s) => s.subscribe);
  const unsubscribe = useJobStore((s) => s.unsubscribe);

  useEffect(() => {
    if (!jobId) return;

    subscribe(jobId);
    return () => {
      unsubscribe(jobId);
    };
  }, [jobId, subscribe, unsubscribe]);
}
