export function HeroSection() {
  return (
    <section className="mb-16 text-center">
      {/* Pill badge */}
      <div className="mb-6 flex justify-center">
        <span
          className="inline-flex items-center gap-1.5 rounded-full border border-[rgba(255,46,85,0.3)] bg-[rgba(255,46,85,0.08)] px-3 py-1.5 text-sm font-medium"
          style={{ color: '#ff2e55' }}
        >
          <span className="size-1.5 rounded-full bg-[#ff2e55]" />
          Self-hosted · Private · Reliable
        </span>
      </div>

      {/* Headline */}
      <h1
        className="mb-4 text-4xl font-extrabold leading-tight tracking-tight text-[#f4f4f5] sm:text-6xl"
        style={{ fontFamily: 'var(--font-display)' }}
      >
        Download anything with{' '}
        <span
          style={{
            background: 'linear-gradient(135deg, #ff2e55 0%, #be123c 100%)',
            WebkitBackgroundClip: 'text',
            WebkitTextFillColor: 'transparent',
            backgroundClip: 'text',
          }}
        >
          yt-dlp
        </span>
      </h1>
      <p className="mx-auto max-w-xl text-lg text-[#a1a1aa]">
        Paste a URL, choose a preset, and let the backend handle the rest.
        Track progress in real time and manage your downloads from one clean UI.
      </p>
    </section>
  );
}
