import { createRootRoute, Link, Outlet } from '@tanstack/react-router';
import { TanStackRouterDevtools } from '@tanstack/router-devtools';
import { Download } from 'lucide-react';
import { useState, useEffect, useCallback } from 'react';
import { api } from '@/lib/api';
import { cn } from '@/lib/utils';

// ── Root layout ───────────────────────────────────────────────────

function RootLayout() {
  const [isOnline, setIsOnline] = useState<'checking' | 'online' | 'offline'>('checking');

  const checkHealth = useCallback(async () => {
    try {
      const res = await api.health();
      if (res && res.error === null && res.data?.status === 'ok') {
        setIsOnline('online');
      } else {
        setIsOnline('offline');
      }
    } catch {
      setIsOnline('offline');
    }
  }, []);

  useEffect(() => {
    checkHealth();
    const interval = setInterval(checkHealth, 15000);
    return () => clearInterval(interval);
  }, [checkHealth]);

  return (
    <div className="flex min-h-screen flex-col">
      {/* ── Navigation bar ──────────────────────────────────────── */}
      <header className="sticky top-0 z-50 border-b border-white/[0.06] bg-[#09090b]/80 backdrop-blur-md">
        <div className="mx-auto flex h-14 max-w-6xl items-center justify-between px-4">
          {/* Logo */}
          <Link to="/" className="flex items-center gap-2.5 no-underline">
            <div
              className="flex size-8 items-center justify-center rounded-md"
              style={{
                background: 'linear-gradient(135deg, #ff2e55 0%, #be123c 100%)',
                boxShadow: '0 0 12px rgba(255,46,85,0.4)',
              }}
            >
              <Download className="size-4 text-white" />
            </div>
            <span
              className="text-base font-bold tracking-none text-[#f4f4f5] uppercase"
              style={{ fontFamily: 'var(--font-mono)' }}
            >
              Youtube Web UI
            </span>
          </Link>

          {/* Right section: Health Status + Nav links */}
          <div className="flex items-center gap-4">
            {/* Status indicator */}
            <div className="flex items-center gap-2 rounded-full border border-white/[0.04] bg-[#18181b]/50 px-3 py-1 text-xs font-medium text-[#a1a1aa] backdrop-blur-sm">
              <span className="relative flex size-1.5">
                {isOnline === 'online' && (
                  <span className="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75" />
                )}
                <span className={cn(
                  "relative inline-flex size-1.5 rounded-full",
                  isOnline === 'online' ? "bg-emerald-400" : isOnline === 'offline' ? "bg-rose-400 animate-pulse" : "bg-amber-400"
                )} />
              </span>
              <span>
                {isOnline === 'online' ? "Backend Online" : isOnline === 'offline' ? "Backend Offline" : "Checking Status…"}
              </span>
            </div>

            {/* Nav links */}
            <nav className="flex items-center gap-1">
              <Link
                to="/"
                className="rounded-md px-3 py-1.5 text-sm font-medium text-[#a1a1aa] no-underline transition-colors hover:bg-white/[0.06] hover:text-[#f4f4f5] [&.active]:text-[#f4f4f5]"
              >
                Home
              </Link>
            </nav>
          </div>
        </div>
      </header>

      {/* ── Main content ─────────────────────────────────────────── */}
      <main className="flex-1">
        <Outlet />
      </main>

      {/* ── Footer ───────────────────────────────────────────────── */}
      <footer className="border-t border-white/[0.06] py-6">
        <div className="mx-auto max-w-6xl px-4 text-center text-sm text-[#71717a]">
          yt-dlp Web UI — self-hosted media downloader
        </div>
      </footer>

      {/* Only shown in development */}
      {import.meta.env.DEV && <TanStackRouterDevtools />}
    </div>
  );
}

export const Route = createRootRoute({
  component: RootLayout,
});
