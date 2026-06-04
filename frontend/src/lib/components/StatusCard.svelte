<script lang="ts">
	import { onMount } from 'svelte';

	interface Props {
		backendUrl?: string;
	}

	let { backendUrl = 'http://localhost:8080/api/health' }: Props = $props();

	let backendStatus: 'checking' | 'online' | 'offline' = $state('checking');

	async function checkBackend() {
		backendStatus = 'checking';
		try {
			const res = await fetch(backendUrl);
			if (res.ok) {
				const data = await res.json();
				if (data.status === 'OK') {
					backendStatus = 'online';
					return;
				}
			}
			backendStatus = 'offline';
		} catch (e) {
			backendStatus = 'offline';
		}
	}

	onMount(() => {
		checkBackend();
	});
</script>

<div class="status-card">
	<div class="status-header">
		<h2 class="status-title">System Status</h2>
		<button class="btn-refresh" onclick={checkBackend} aria-label="Refresh Status">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="refresh-icon">
				<path d="M21.5 2v6h-6M21.34 15.57a10 10 0 1 1-.57-8.38l5.67-5.67"></path>
			</svg>
		</button>
	</div>

	<div class="status-indicators">
		<div class="indicator-row">
			<div class="indicator-info">
				<span class="indicator-label">Frontend Server</span>
				<span class="indicator-desc">SvelteKit & Vite development server</span>
			</div>
			<div class="status-pill status-online">
				<span class="dot"></span> Running
			</div>
		</div>

		<div class="indicator-row">
			<div class="indicator-info">
				<span class="indicator-label">Backend API</span>
				<span class="indicator-desc">Go HTTP API server at {backendUrl}</span>
			</div>
			{#if backendStatus === 'checking'}
				<div class="status-pill status-checking">
					<span class="dot"></span> Checking...
				</div>
			{:else if backendStatus === 'online'}
				<div class="status-pill status-online">
					<span class="dot"></span> Connected
				</div>
			{:else}
				<div class="status-pill status-offline">
					<span class="dot"></span> Offline
				</div>
			{/if}
		</div>
	</div>
</div>

<style>
	.status-card {
		background: rgba(24, 24, 27, 0.6);
		backdrop-filter: blur(12px);
		border: 1px solid var(--border-color);
		border-radius: var(--radius-lg);
		padding: 1.5rem;
		margin-bottom: 4rem;
		text-align: left;
		box-shadow: var(--shadow-lg), var(--shadow-glow);
		transition: var(--transition-all);
	}

	.status-card:hover {
		border-color: rgba(255, 46, 85, 0.2);
	}

	.status-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 1.25rem;
		padding-bottom: 0.75rem;
		border-bottom: 1px solid var(--border-color);
	}

	.status-title {
		font-family: var(--font-display);
		font-size: 1.25rem;
		font-weight: 600;
	}

	.btn-refresh {
		background: transparent;
		border: none;
		color: var(--text-secondary);
		cursor: pointer;
		padding: 0.25rem;
		border-radius: var(--radius-sm);
		transition: var(--transition-all);
	}

	.btn-refresh:hover {
		color: var(--primary);
		background: rgba(255, 46, 85, 0.05);
	}

	.refresh-icon {
		width: 1.25rem;
		height: 1.25rem;
	}

	.status-indicators {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.indicator-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 0.5rem 0;
	}

	.indicator-info {
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
	}

	.indicator-label {
		font-weight: 500;
		font-size: 0.95rem;
	}

	.indicator-desc {
		font-size: 0.75rem;
		color: var(--text-muted);
	}

	.status-pill {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		font-size: 0.85rem;
		font-weight: 600;
		padding: 0.25rem 0.75rem;
		border-radius: var(--radius-full);
	}

	.status-online {
		background: rgba(16, 185, 129, 0.1);
		color: #10b981;
		border: 1px solid rgba(16, 185, 129, 0.2);
	}

	.status-online .dot {
		background: #10b981;
		box-shadow: 0 0 8px #10b981;
	}

	.status-checking {
		background: rgba(245, 158, 11, 0.1);
		color: #f59e0b;
		border: 1px solid rgba(245, 158, 11, 0.2);
	}

	.status-checking .dot {
		background: #f59e0b;
		box-shadow: 0 0 8px #f59e0b;
		animation: pulse 1.5s infinite;
	}

	.status-offline {
		background: rgba(239, 68, 68, 0.1);
		color: #ef4444;
		border: 1px solid rgba(239, 68, 68, 0.2);
	}

	.status-offline .dot {
		background: #ef4444;
		box-shadow: 0 0 8px #ef4444;
	}

	.dot {
		width: 6px;
		height: 6px;
		border-radius: 50%;
	}

	@keyframes pulse {
		0%, 100% { opacity: 0.5; }
		50% { opacity: 1; }
	}
</style>
