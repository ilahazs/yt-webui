<script lang="ts">
	import type { Snippet } from 'svelte';

	interface Props {
		title: string;
		description: string;
		icon: Snippet;
	}

	let { title, description, icon }: Props = $props();
</script>

<div class="feature-card">
	<div class="card-icon">
		{@render icon()}
	</div>
	<h3>{title}</h3>
	<p>{description}</p>
</div>

<style>
	.feature-card {
		background: rgba(24, 24, 27, 0.4);
		border: 1px solid var(--border-color);
		border-radius: var(--radius-md);
		padding: 1.5rem;
		transition: var(--transition-all);
	}

	.feature-card:hover {
		background: var(--bg-card-hover);
		transform: translateY(-2px);
		border-color: rgba(255, 46, 85, 0.15);
	}

	.card-icon {
		width: 2.5rem;
		height: 2.5rem;
		border-radius: var(--radius-sm);
		background: rgba(255, 46, 85, 0.05);
		color: var(--primary);
		display: flex;
		align-items: center;
		justify-content: center;
		margin-bottom: 1.25rem;
		border: 1px solid rgba(255, 46, 85, 0.1);
	}

	.card-icon :global(svg) {
		width: 1.25rem;
		height: 1.25rem;
	}

	.feature-card h3 {
		font-family: var(--font-display);
		font-size: 1.1rem;
		font-weight: 600;
		margin-bottom: 0.5rem;
	}

	.feature-card p {
		font-size: 0.875rem;
		color: var(--text-secondary);
		line-height: 1.5;
	}
</style>
