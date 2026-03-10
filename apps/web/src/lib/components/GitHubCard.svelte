<script>
	import { onMount } from 'svelte';

	let { repo = 'svviitzerland/LLMAO' } = $props();

	let repoData = $state(null);
	let loading = $state(true);
	let error = $state(null);

	onMount(async () => {
		try {
			const res = await fetch(`https://api.github.com/repos/${repo}`);
			if (!res.ok) throw new Error('Failed to fetch');
			repoData = await res.json();
		} catch (e) {
			error = e.message;
		} finally {
			loading = false;
		}
	});

	function formatStars(num) {
		if (num >= 1000) return (num / 1000).toFixed(1) + 'k';
		return num;
	}
</script>

<a
	href={`https://github.com/${repo}`}
	target="_blank"
	rel="noopener noreferrer"
	class="github-card"
>
	{#if loading}
		<div class="loading">
			<div class="shimmer"></div>
			<div class="shimmer-text"></div>
			<div class="shimmer-stats"></div>
		</div>
	{:else if error}
		<div class="card-content">
			<div class="header">
				<svg viewBox="0 0 16 16" width="24" height="24" fill="currentColor" class="github-icon">
					<path
						d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"
					/>
				</svg>
				<span class="repo-name">{repo}</span>
			</div>
		</div>
	{:else if repoData}
		<div class="card-content">
			<div class="header">
				<div class="icon-glow">
					<svg viewBox="0 0 16 16" width="28" height="28" fill="currentColor" class="github-icon">
						<path
							d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27.68 0 1.36.09 2 .27 1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.013 8.013 0 0016 8c0-4.42-3.58-8-8-8z"
						/>
					</svg>
				</div>
				<div class="repo-info">
					<span class="repo-name">{repoData.full_name}</span>
					<span class="badge">Open Source</span>
				</div>
			</div>

			{#if repoData.description}
				<p class="description">{repoData.description}</p>
			{/if}

			<div class="stats">
				{#if repoData.language}
					<span class="stat language">
						<span class="lang-dot"></span>
						{repoData.language}
					</span>
				{/if}
				<span class="stat stars">
					<svg viewBox="0 0 16 16" width="16" height="16" fill="currentColor">
						<path
							d="M8 .25a.75.75 0 01.673.418l1.882 3.815 4.21.612a.75.75 0 01.416 1.279l-3.046 2.97.719 4.192a.75.75 0 01-1.088.791L8 12.347l-3.766 1.98a.75.75 0 01-1.088-.79l.72-4.194L.818 6.374a.75.75 0 01.416-1.28l4.21-.611L7.327.668A.75.75 0 018 .25z"
						/>
					</svg>
					{formatStars(repoData.stargazers_count)}
				</span>
				<span class="stat forks">
					<svg viewBox="0 0 16 16" width="16" height="16" fill="currentColor">
						<path
							d="M5 5.372v.878c0 .414.336.75.75.75h4.5a.75.75 0 00.75-.75v-.878a2.25 2.25 0 111.5 0v.878a2.25 2.25 0 01-2.25 2.25h-1.5v2.128a2.251 2.251 0 11-1.5 0V8.5h-1.5A2.25 2.25 0 013.5 6.25v-.878a2.25 2.25 0 111.5 0zM5 3.25a.75.75 0 10-1.5 0 .75.75 0 001.5 0zm6.75.75a.75.75 0 100-1.5.75.75 0 000 1.5zm-3 8.75a.75.75 0 10-1.5 0 .75.75 0 001.5 0z"
						/>
					</svg>
					{repoData.forks_count}
				</span>
			</div>

			<div class="cta">
				<span>View on GitHub</span>
				<svg viewBox="0 0 16 16" width="14" height="14" fill="currentColor">
					<path
						d="M6.22 3.22a.75.75 0 011.06 0l4.25 4.25a.75.75 0 010 1.06l-4.25 4.25a.75.75 0 01-1.06-1.06L9.94 8 6.22 4.28a.75.75 0 010-1.06z"
					/>
				</svg>
			</div>
		</div>
	{/if}
</a>

<style>
	.github-card {
		display: block;
		background: linear-gradient(
			145deg,
			rgba(249, 115, 22, 0.08) 0%,
			rgba(234, 88, 12, 0.04) 50%,
			rgba(20, 20, 20, 0.95) 100%
		);
		border: 1px solid rgba(249, 115, 22, 0.2);
		border-radius: 16px;
		padding: 1.75rem;
		text-decoration: none;
		color: #e5e7eb;
		transition: all 0.4s cubic-bezier(0.4, 0, 0.2, 1);
		margin: 2.5rem 0;
		position: relative;
		overflow: hidden;
	}

	.github-card::before {
		content: '';
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		height: 1px;
		background: linear-gradient(90deg, transparent, rgba(249, 115, 22, 0.5), transparent);
	}

	.github-card:hover {
		border-color: rgba(249, 115, 22, 0.5);
		transform: translateY(-4px);
		box-shadow:
			0 20px 40px rgba(0, 0, 0, 0.4),
			0 0 40px rgba(249, 115, 22, 0.15);
	}

	.loading {
		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.shimmer {
		width: 40px;
		height: 40px;
		border-radius: 12px;
		background: linear-gradient(90deg, #27272a 25%, #3f3f46 50%, #27272a 75%);
		background-size: 200% 100%;
		animation: shimmer 1.5s infinite;
	}

	.shimmer-text {
		width: 70%;
		height: 20px;
		border-radius: 6px;
		background: linear-gradient(90deg, #27272a 25%, #3f3f46 50%, #27272a 75%);
		background-size: 200% 100%;
		animation: shimmer 1.5s infinite;
	}

	.shimmer-stats {
		width: 40%;
		height: 16px;
		border-radius: 6px;
		background: linear-gradient(90deg, #27272a 25%, #3f3f46 50%, #27272a 75%);
		background-size: 200% 100%;
		animation: shimmer 1.5s infinite;
	}

	@keyframes shimmer {
		0% {
			background-position: 200% 0;
		}
		100% {
			background-position: -200% 0;
		}
	}

	.header {
		display: flex;
		align-items: center;
		gap: 1rem;
		margin-bottom: 1rem;
	}

	.icon-glow {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 48px;
		height: 48px;
		border-radius: 12px;
		background: linear-gradient(135deg, rgba(249, 115, 22, 0.2) 0%, rgba(234, 88, 12, 0.1) 100%);
		box-shadow: 0 0 20px rgba(249, 115, 22, 0.2);
	}

	.github-icon {
		color: #f97316;
	}

	.repo-info {
		display: flex;
		flex-direction: column;
		gap: 0.25rem;
	}

	.repo-name {
		font-weight: 700;
		color: #fff;
		font-size: 1.2rem;
		letter-spacing: -0.02em;
	}

	.badge {
		font-size: 0.65rem;
		font-weight: 600;
		text-transform: uppercase;
		letter-spacing: 0.1em;
		color: #f97316;
		background: rgba(249, 115, 22, 0.1);
		padding: 0.2rem 0.5rem;
		border-radius: 4px;
		width: fit-content;
	}

	.description {
		color: #a1a1aa;
		font-size: 1rem;
		line-height: 1.6;
		margin: 0 0 1.25rem 0;
	}

	.stats {
		display: flex;
		flex-wrap: wrap;
		gap: 1.25rem;
		margin-bottom: 1.25rem;
		padding-bottom: 1.25rem;
		border-bottom: 1px solid rgba(255, 255, 255, 0.06);
	}

	.stat {
		display: flex;
		align-items: center;
		gap: 0.4rem;
		font-size: 0.9rem;
		color: #a1a1aa;
		font-weight: 500;
	}

	.stat svg {
		color: #71717a;
	}

	.stat.stars svg {
		color: #fbbf24;
	}

	.lang-dot {
		width: 12px;
		height: 12px;
		border-radius: 50%;
		background: linear-gradient(135deg, #dea584 0%, #f97316 100%);
		box-shadow: 0 0 8px rgba(249, 115, 22, 0.4);
	}

	.cta {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		color: #f97316;
		font-weight: 600;
		font-size: 0.9rem;
		transition: all 0.3s ease;
	}

	.github-card:hover .cta {
		gap: 0.75rem;
	}

	.cta svg {
		transition: transform 0.3s ease;
	}

	.github-card:hover .cta svg {
		transform: translateX(4px);
	}
</style>
