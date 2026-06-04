import { useState, useCallback } from 'react';
import { api } from '@/lib/api';
import type { ProbeData } from '@/lib/types';

// ── Types ────────────────────────────────────────────────────────

interface UseProbeState {
  /** Probed metadata, or null if not yet fetched. */
  data: ProbeData | null;
  /** Whether a probe request is in flight. */
  isLoading: boolean;
  /** Error message from the last failed probe, or null. */
  error: string | null;
}

interface UseProbeReturn extends UseProbeState {
  /** Trigger a probe for the given URL. */
  probe: (url: string) => Promise<void>;
  /** Clear the current probe result and error state. */
  reset: () => void;
}

// ── Hook ─────────────────────────────────────────────────────────

/**
 * Manages the URL probe lifecycle — fetching metadata before
 * a download job is created.
 */
export function useProbe(): UseProbeReturn {
  const [state, setState] = useState<UseProbeState>({
    data: null,
    isLoading: false,
    error: null,
  });

  const probe = useCallback(async (url: string) => {
    if (!url.trim()) {
      setState((s) => ({ ...s, error: 'Please enter a valid URL.' }));
      return;
    }

    setState({ data: null, isLoading: true, error: null });

    const res = await api.probe({ url: url.trim() });

    if (res.error !== null) {
      setState({ data: null, isLoading: false, error: res.error.message });
      return;
    }

    setState({ data: res.data, isLoading: false, error: null });
  }, []);

  const reset = useCallback(() => {
    setState({ data: null, isLoading: false, error: null });
  }, []);

  return { ...state, probe, reset };
}
