import { Search, Loader2, AlertCircle } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

interface ProbeFormProps {
  url: string;
  onUrlChange: (value: string) => void;
  onSubmit: (e: React.FormEvent) => void;
  isProbing: boolean;
  probeError: string | null;
}

export function ProbeForm({
  url,
  onUrlChange,
  onSubmit,
  isProbing,
  probeError,
}: ProbeFormProps) {
  return (
    <>
      <form onSubmit={onSubmit} className="flex gap-3">
        <div className="relative flex-1">
          <Search
            className="absolute left-4 top-1/2 size-5 -translate-y-1/2 text-[#71717a]"
          />
          <Input
            id="url-input"
            type="url"
            placeholder="Paste a video URL…"
            value={url}
            onChange={(e) => onUrlChange(e.target.value)}
            className="h-12 border-white/[0.08] bg-[#09090b] pl-11 text-base text-[#f4f4f5] placeholder:text-[#71717a] focus-visible:ring-[rgba(255,46,85,0.5)]"
          />
        </div>
        <Button
          id="probe-button"
          type="submit"
          disabled={isProbing || !url.trim()}
          className="h-12 gap-2 px-6 text-base font-semibold"
          style={{
            background: 'linear-gradient(135deg, #ff2e55 0%, #be123c 100%)',
            boxShadow: '0 0 20px rgba(255,46,85,0.25)',
          }}
        >
          {isProbing ? (
            <Loader2 className="size-4 animate-spin" />
          ) : (
            <Search className="size-4" />
          )}
          {isProbing ? 'Probing…' : 'Probe'}
        </Button>
      </form>

      {/* Error banner */}
      {probeError && (
        <div className="mt-4 flex items-center gap-2 rounded-lg border border-red-900/40 bg-red-950/30 px-4 py-3 text-sm text-red-400">
          <AlertCircle className="size-4 shrink-0" />
          {probeError}
        </div>
      )}
    </>
  );
}
