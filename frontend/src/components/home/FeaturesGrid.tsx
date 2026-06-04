import {
  Search,
  Zap,
  Download,
  ShieldCheck,
  History,
  FileVideo,
  ArrowRight,
} from 'lucide-react';
import { FeatureCard } from '@/components/FeatureCard';

const FEATURES = [
  {
    icon: Search,
    title: 'URL Probing',
    description: 'Inspect any media URL and preview title, thumbnail, and available formats before downloading.',
  },
  {
    icon: Zap,
    title: 'Realtime Progress',
    description: 'Watch download progress live via Server-Sent Events — no page refresh needed.',
  },
  {
    icon: Download,
    title: 'Smart Presets',
    description: 'Choose from curated download presets. No raw yt-dlp flags, just sensible options.',
  },
  {
    icon: ShieldCheck,
    title: 'Safe Execution',
    description: 'Downloads run in a controlled backend worker. No shell injection, no arbitrary commands.',
  },
  {
    icon: History,
    title: 'Job History',
    description: 'All jobs are persisted. Review past downloads, retry failures, or grab completed files.',
  },
  {
    icon: FileVideo,
    title: 'Format Selection',
    description: 'Automatically picks the best format for your preset, with ffmpeg post-processing support.',
  },
] as const;

export function FeaturesGrid() {
  return (
    <section>
      <div className="mb-8 flex items-center justify-between">
        <h2
          className="text-2xl font-bold text-[#f4f4f5]"
          style={{ fontFamily: 'var(--font-display)' }}
        >
          What it does
        </h2>
        <span className="flex items-center gap-1 text-sm text-[#71717a]">
          More features coming <ArrowRight className="size-3" />
        </span>
      </div>
      <div className="grid grid-cols-1 gap-3 sm:grid-cols-2 lg:grid-cols-3">
        {FEATURES.map((feature) => (
          <FeatureCard
            key={feature.title}
            icon={feature.icon}
            title={feature.title}
            description={feature.description}
          />
        ))}
      </div>
    </section>
  );
}
