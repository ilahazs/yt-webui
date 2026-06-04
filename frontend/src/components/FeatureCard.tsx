import type { LucideIcon } from 'lucide-react';
import { cn } from '@/lib/utils';

// ── Types ─────────────────────────────────────────────────────────

export interface FeatureCardProps {
  /** Lucide icon component to display. */
  icon: LucideIcon;
  /** Short title for the feature. */
  title: string;
  /** One or two sentence description. */
  description: string;
  /** Optional extra class names. */
  className?: string;
}

// ── Component ─────────────────────────────────────────────────────

/**
 * Displays a single feature/capability highlight on the landing page.
 * Uses the app's dark glassmorphism style.
 */
export function FeatureCard({
  icon: Icon,
  title,
  description,
  className,
}: FeatureCardProps) {
  return (
    <div
      className={cn(
        'group relative flex flex-col gap-3 rounded-xl border p-6',
        'border-white/[0.08] bg-[#18181b]/80 backdrop-blur-sm',
        'transition-all duration-300',
        'hover:border-[rgba(255,46,85,0.3)] hover:bg-[#27272a]/80',
        'hover:shadow-[0_0_20px_0_rgba(255,46,85,0.08)]',
        className,
      )}
    >
      {/* Glow accent */}
      <div
        className="pointer-events-none absolute inset-0 rounded-xl opacity-0 transition-opacity duration-300 group-hover:opacity-100"
        style={{
          background:
            'linear-gradient(135deg, rgba(255,46,85,0.04) 0%, rgba(190,18,60,0.02) 100%)',
        }}
      />

      {/* Icon */}
      <div
        className="flex size-10 items-center justify-center rounded-lg"
        style={{ background: 'linear-gradient(135deg, rgba(255,46,85,0.15) 0%, rgba(190,18,60,0.08) 100%)' }}
      >
        <Icon
          className="size-5 transition-transform duration-300 group-hover:scale-110"
          style={{ color: '#ff2e55' }}
        />
      </div>

      {/* Text */}
      <div className="flex flex-col gap-1.5">
        <h3 className="text-base font-semibold text-[#f4f4f5]">{title}</h3>
        <p className="text-sm leading-relaxed text-[#a1a1aa]">{description}</p>
      </div>
    </div>
  );
}
