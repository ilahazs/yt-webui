<script lang="ts">
	import StatusCard from '$lib/components/StatusCard.svelte';
	import FeatureCard from '$lib/components/FeatureCard.svelte';
	import Footer from '$lib/components/Footer.svelte';
</script>

<main class="container">
	<div class="glow-container">
		<div class="glow-spot"></div>
	</div>

	<header class="header">
		<div class="logo">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" stroke-linecap="round" stroke-linejoin="round" class="logo-icon">
				<path d="M22.54 6.42a2.78 2.78 0 0 0-1.94-2C18.88 4 12 4 12 4s-6.88 0-8.6.46a2.78 2.78 0 0 0-1.94 2A29 29 0 0 0 1 11.75a29 29 0 0 0 .46 5.33A2.78 2.78 0 0 0 3.4 19c1.72.46 8.6.46 8.6.46s6.88 0 8.6-.46a2.78 2.78 0 0 0 1.94-2 29 29 0 0 0 .46-5.25 29 29 0 0 0-.46-5.33z"></path>
				<polygon points="9.75 15.02 15.5 11.75 9.75 8.48 9.75 15.02"></polygon>
			</svg>
			<span class="logo-text">yt-webui</span>
		</div>
		<div class="badge">Fully Free</div>
	</header>

	<section class="hero">
		<h1 class="hero-title">Media Downloader <span>Reimagined</span></h1>
		<p class="hero-subtitle">A premium self-hosted web interface powered by Go and SvelteKit for managing your downloads using yt-dlp.</p>

		<StatusCard />

		<div class="card-grid">
			<FeatureCard title="Metadata Probe" description="Analyze target URLs and preview titles, durations, thumbnails, and format options before downloading.">
				{#snippet icon()}
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<circle cx="12" cy="12" r="10"></circle>
						<line x1="12" y1="8" x2="12" y2="12"></line>
						<line x1="12" y1="16" x2="12.01" y2="16"></line>
					</svg>
				{/snippet}
			</FeatureCard>

			<FeatureCard title="Smart Presets" description="Define global profiles for audio conversion, best video quality, or custom resolution targets.">
				{#snippet icon()}
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<line x1="6" y1="3" x2="6" y2="15"></line>
						<circle cx="18" cy="6" r="3"></circle>
						<circle cx="6" cy="18" r="3"></circle>
						<path d="M18 9a9 9 0 0 1-9 9"></path>
					</svg>
				{/snippet}
			</FeatureCard>

			<FeatureCard title="Queue Manager" description="Track active downloads with real-time speed, progress bars, and historical archive logs.">
				{#snippet icon()}
					<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
						<path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"></path>
						<polyline points="7 10 12 15 17 10"></polyline>
						<line x1="12" y1="15" x2="12" y2="3"></line>
					</svg>
				{/snippet}
			</FeatureCard>
		</div>
	</section>

	<Footer />
</main>

<style>
	.container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 2rem 1.5rem;
		position: relative;
	}

	.glow-container {
		position: absolute;
		top: -100px;
		left: 50%;
		transform: translateX(-50%);
		width: 100%;
		max-width: 800px;
		height: 400px;
		overflow: hidden;
		z-index: -1;
		pointer-events: none;
	}

	.glow-spot {
		width: 100%;
		height: 100%;
		background: radial-gradient(circle, rgba(255, 46, 85, 0.08) 0%, rgba(255, 46, 85, 0) 70%);
		filter: blur(40px);
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding-bottom: 2rem;
		border-bottom: 1px solid var(--border-color);
		margin-bottom: 3.5rem;
	}

	.logo {
		display: flex;
		align-items: center;
		gap: 0.75rem;
	}

	.logo-icon {
		width: 2rem;
		height: 2rem;
		color: var(--primary);
		filter: drop-shadow(0 0 8px rgba(255, 46, 85, 0.4));
	}

	.logo-text {
		font-family: var(--font-display);
		font-size: 1.5rem;
		font-weight: 700;
		letter-spacing: -0.025em;
		background: linear-gradient(to right, #ffffff, #d4d4d8);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}

	.badge {
		font-size: 0.75rem;
		font-weight: 600;
		padding: 0.25rem 0.75rem;
		border-radius: var(--radius-full);
		background: rgba(255, 46, 85, 0.1);
		color: var(--primary);
		border: 1px solid rgba(255, 46, 85, 0.2);
	}

	.hero {
		text-align: center;
		max-width: 800px;
		margin: 0 auto;
	}

	.hero-title {
		font-family: var(--font-display);
		font-size: 3.5rem;
		font-weight: 800;
		letter-spacing: -0.03em;
		line-height: 1.15;
		margin-bottom: 1.25rem;
	}

	.hero-title span {
		background: var(--primary-gradient);
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
		position: relative;
	}

	.hero-subtitle {
		font-size: 1.125rem;
		color: var(--text-secondary);
		line-height: 1.6;
		margin-bottom: 3rem;
		max-width: 600px;
		margin-left: auto;
		margin-right: auto;
	}

	.card-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
		gap: 1.5rem;
		text-align: left;
	}

	@media (max-width: 640px) {
		.hero-title {
			font-size: 2.5rem;
		}
		.header {
			margin-bottom: 2rem;
		}
	}
</style>
