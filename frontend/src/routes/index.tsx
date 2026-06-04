import { createFileRoute } from '@tanstack/react-router';
import { useState } from 'react';
import { useProbe } from '@/hooks/useProbe';
import type { DownloadPreset } from '@/lib/types';
import { api } from '@/lib/api';
import { HeroSection } from '@/components/home/HeroSection';
import { ProbeForm } from '@/components/home/ProbeForm';
import { ProbeResult } from '@/components/home/ProbeResult';
import { FeaturesGrid } from '@/components/home/FeaturesGrid';

// ── Route ─────────────────────────────────────────────────────────

export const Route = createFileRoute('/')({
  component: HomePage,
});

// ── Component ─────────────────────────────────────────────────────

function HomePage() {
  const [url, setUrl] = useState('');
  const [selectedPreset, setSelectedPreset] = useState<DownloadPreset>('best_video');
  const [isCreatingJob, setIsCreatingJob] = useState(false);
  const [jobCreated, setJobCreated] = useState<string | null>(null);
  const [jobError, setJobError] = useState<string | null>(null);

  const { data: probeData, isLoading: isProbing, error: probeError, probe, reset } = useProbe();

  const handleProbe = (e: React.FormEvent) => {
    e.preventDefault();
    reset();
    setJobCreated(null);
    setJobError(null);
    probe(url);
  };

  const handleCreateJob = async () => {
    if (!probeData) return;
    setIsCreatingJob(true);
    setJobError(null);

    const res = await api.createJob({
      url: probeData.webpage_url,
      preset: selectedPreset,
      options: { embed_metadata: true },
    });

    setIsCreatingJob(false);

    if (res.error !== null) {
      setJobError(res.error.message);
      return;
    }

    setJobCreated(res.data.job_id);
  };

  return (
    <div className="mx-auto max-w-6xl px-4 py-12 mt-25">
      {/* ── Hero ──────────────────────────────────────────────── */}
      <HeroSection />

      {/* ── Probe card container ──────────────────────────────── */}
      <section className="mb-12">
        <div
          className="rounded-2xl border border-white/[0.08] bg-[#18181b]/80 p-6 backdrop-blur-sm"
          style={{ boxShadow: '0 0 40px rgba(255,46,85,0.04)' }}
        >
          <ProbeForm
            url={url}
            onUrlChange={setUrl}
            onSubmit={handleProbe}
            isProbing={isProbing}
            probeError={probeError}
          />

          {probeData && (
            <ProbeResult
              probeData={probeData}
              selectedPreset={selectedPreset}
              onPresetSelect={setSelectedPreset}
              onCreateJob={handleCreateJob}
              isCreatingJob={isCreatingJob}
              jobCreated={jobCreated}
              jobError={jobError}
            />
          )}
        </div>
      </section>

      {/* ── Features grid ─────────────────────────────────────── */}
      <FeaturesGrid />
    </div>
  );
}
