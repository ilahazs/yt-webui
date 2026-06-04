import { User, Clock, Loader2, Download, AlertCircle } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Badge } from '@/components/ui/badge';
import { Separator } from '@/components/ui/separator';
import { cn } from '@/lib/utils';
import type { DownloadPreset, ProbeData } from '@/lib/types';

interface PresetOption {
  id: DownloadPreset;
  label: string;
  description: string;
}

const PRESETS: PresetOption[] = [
  { id: 'best_video', label: 'Best Video', description: 'Highest quality video + audio' },
  { id: 'best_720p', label: '720p HD', description: 'Balanced quality and size' },
  { id: 'best_480p', label: '480p SD', description: 'Smaller file size' },
  { id: 'best_audio', label: 'Audio Only', description: 'Best quality audio (MP3/AAC)' },
];

function formatDuration(seconds: number): string {
  const h = Math.floor(seconds / 3600);
  const m = Math.floor((seconds % 3600) / 60);
  const s = seconds % 60;
  if (h > 0) return `${h}:${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`;
  return `${m}:${String(s).padStart(2, '0')}`;
}

interface ProbeResultProps {
  probeData: ProbeData;
  selectedPreset: DownloadPreset;
  onPresetSelect: (preset: DownloadPreset) => void;
  onCreateJob: () => void;
  isCreatingJob: boolean;
  jobCreated: string | null;
  jobError: string | null;
}

export function ProbeResult({
  probeData,
  selectedPreset,
  onPresetSelect,
  onCreateJob,
  isCreatingJob,
  jobCreated,
  jobError,
}: ProbeResultProps) {
  return (
    <div className="mt-6">
      <Separator className="mb-6 bg-white/[0.06]" />

      {/* Media preview */}
      <div className="flex gap-4">
        {probeData.thumbnail && (
          <div className="relative h-24 w-40 shrink-0 overflow-hidden rounded-lg">
            <img
              src={probeData.thumbnail}
              alt={probeData.title}
              className="h-full w-full object-cover"
            />
          </div>
        )}
        <div className="flex min-w-0 flex-1 flex-col gap-2">
          <h2 className="truncate text-base font-semibold text-[#f4f4f5]">
            {probeData.title}
          </h2>
          <div className="flex flex-wrap items-center gap-3 text-sm text-[#a1a1aa]">
            {probeData.uploader && (
              <span className="flex items-center gap-1">
                <User className="size-3.5" />
                {probeData.uploader}
              </span>
            )}
            {probeData.duration_seconds !== undefined && (
              <span className="flex items-center gap-1">
                <Clock className="size-3.5" />
                {formatDuration(probeData.duration_seconds)}
              </span>
            )}
            {probeData.formats.length > 0 && (
              <Badge variant="secondary" className="text-xs py-0.5 px-2">
                {probeData.formats.length} formats
              </Badge>
            )}
          </div>
        </div>
      </div>

      {/* Preset selection */}
      <div className="mt-5">
        <p className="mb-3 text-sm font-medium text-[#a1a1aa]">Select preset</p>
        <div className="grid grid-cols-2 gap-2.5 sm:grid-cols-4">
          {PRESETS.map((preset) => (
            <button
              key={preset.id}
              id={`preset-${preset.id}`}
              type="button"
              onClick={() => onPresetSelect(preset.id)}
              className={cn(
                'flex flex-col gap-1 rounded-lg border px-4 py-3 text-left transition-all',
                selectedPreset === preset.id
                  ? 'border-[rgba(255,46,85,0.5)] bg-[rgba(255,46,85,0.08)] text-[#f4f4f5]'
                  : 'border-white/[0.08] bg-[#09090b] text-[#a1a1aa] hover:border-white/[0.15] hover:text-[#f4f4f5]',
              )}
            >
              <span className="text-sm font-semibold">{preset.label}</span>
              <span className="text-xs opacity-70">{preset.description}</span>
            </button>
          ))}
        </div>
      </div>

      {/* Job creation */}
      <div className="mt-5 flex items-center gap-3">
        <Button
          id="create-job-button"
          onClick={onCreateJob}
          disabled={isCreatingJob}
          className="h-11 gap-2 px-5 text-base font-semibold"
          style={{
            background: 'linear-gradient(135deg, #ff2e55 0%, #be123c 100%)',
            boxShadow: '0 0 16px rgba(255,46,85,0.2)',
          }}
        >
          {isCreatingJob ? (
            <Loader2 className="size-4 animate-spin" />
          ) : (
            <Download className="size-4" />
          )}
          {isCreatingJob ? 'Creating…' : 'Start Download'}
        </Button>

        {jobCreated && (
          <p className="flex items-center gap-1.5 text-sm text-emerald-400">
            <span className="size-1.5 rounded-full bg-emerald-400" />
            Job created:{' '}
            <span className="font-mono opacity-80">{jobCreated}</span>
          </p>
        )}

        {jobError && (
          <p className="flex items-center gap-1.5 text-sm text-red-400">
            <AlertCircle className="size-3.5" />
            {jobError}
          </p>
        )}
      </div>
    </div>
  );
}
