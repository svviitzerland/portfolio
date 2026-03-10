<script>
	import cv from '$lib/data/cv.json';

	const { socials } = cv;

	let dockEl = $state(null);
	let mouseX = $state(-1);
	let isHovering = $state(false);

	// SVG icon paths for each social network
	const iconMap = {
		GitHub: {
			viewBox: '0 0 24 24',
			path: 'M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z'
		},
		LinkedIn: {
			viewBox: '0 0 24 24',
			path: 'M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z'
		},
		Blog: {
			viewBox: '0 0 24 24',
			path: 'M19 20H5a2 2 0 01-2-2V6a2 2 0 012-2h10a2 2 0 012 2v1h2a2 2 0 012 2v9a2 2 0 01-2 2zM5 6v12h14V9h-4V6H5zm2 3h6v2H7V9zm0 4h10v2H7v-2zm0 4h10v2H7v-2z'
		},
		Medium: {
			viewBox: '0 0 24 24',
			path: 'M13.54 12a6.8 6.8 0 01-6.77 6.82A6.8 6.8 0 010 12a6.8 6.8 0 016.77-6.82A6.8 6.8 0 0113.54 12zM20.96 12c0 3.54-1.51 6.42-3.38 6.42-1.87 0-3.39-2.88-3.39-6.42s1.52-6.42 3.39-6.42 3.38 2.88 3.38 6.42M24 12c0 3.17-.53 5.75-1.19 5.75-.66 0-1.19-2.58-1.19-5.75s.53-5.75 1.19-5.75C23.47 6.25 24 8.83 24 12z'
		},
		Instagram: {
			viewBox: '0 0 24 24',
			path: 'M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zM12 0C8.741 0 8.333.014 7.053.072 2.695.272.273 2.69.073 7.052.014 8.333 0 8.741 0 12c0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98C8.333 23.986 8.741 24 12 24c3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98C15.668.014 15.259 0 12 0zm0 5.838a6.162 6.162 0 100 12.324 6.162 6.162 0 000-12.324zM12 16a4 4 0 110-8 4 4 0 010 8zm6.406-11.845a1.44 1.44 0 100 2.881 1.44 1.44 0 000-2.881z'
		}
	};

	function handleMouseMove(e) {
		if (!dockEl) return;
		const rect = dockEl.getBoundingClientRect();
		mouseX = e.clientX - rect.left;
		isHovering = true;
	}

	function handleMouseLeave() {
		isHovering = false;
		mouseX = -1;
	}

	function getScale(index) {
		if (!isHovering || mouseX < 0) return 1;
		// Each icon is ~48px wide with ~12px gap = 60px per item
		const iconCenter = index * 60 + 30;
		const distance = Math.abs(mouseX - iconCenter);
		const maxDistance = 120;

		if (distance > maxDistance) return 1;

		// Cosine-based smooth scaling
		const scale = 1 + 0.8 * Math.cos((distance / maxDistance) * (Math.PI / 2));
		return Math.min(scale, 1.8);
	}

	let tooltipIndex = $state(-1);
</script>

<!-- Dock -->
<div class="fixed bottom-3 left-1/2 -translate-x-1/2 z-40">
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<nav
		bind:this={dockEl}
		onmousemove={handleMouseMove}
		onmouseleave={handleMouseLeave}
		class="flex items-end gap-3 px-4 py-2 bg-zinc-800/50 backdrop-blur-2xl border border-white/20 rounded-2xl"
		aria-label="Social links dock"
	>
		{#each socials as social, i (social.network)}
			{@const scale = getScale(i)}
			<div class="relative flex flex-col items-center">
				<!-- Tooltip -->
				{#if tooltipIndex === i}
					<div
						class="absolute -top-10 left-1/2 -translate-x-1/2 px-3 py-1 bg-zinc-800 border border-zinc-600 rounded-lg text-xs text-zinc-200 whitespace-nowrap shadow-xl pointer-events-none"
					>
						{social.network}
						<div
							class="absolute -bottom-1 left-1/2 -translate-x-1/2 w-2 h-2 bg-zinc-800 border-r border-b border-zinc-600 rotate-45"
						></div>
					</div>
				{/if}

				<a
					href={social.url}
					target="_blank"
					rel="noopener noreferrer"
					class="flex items-center justify-center w-12 h-12 rounded-xl bg-zinc-700/80 text-zinc-300 hover:text-white transition-colors duration-150"
					style="transform: scale({scale}); transform-origin: bottom; transition: transform 0.15s cubic-bezier(0.22, 1, 0.36, 1);"
					onmouseenter={() => (tooltipIndex = i)}
					onmouseleave={() => (tooltipIndex = -1)}
					aria-label={social.network}
				>
					{#if iconMap[social.network]}
						<svg
							class="w-6 h-6"
							viewBox={iconMap[social.network].viewBox}
							fill="currentColor"
						>
							<path d={iconMap[social.network].path} />
						</svg>
					{:else}
						<span class="text-lg font-bold">{social.network[0]}</span>
					{/if}
				</a>
			</div>
		{/each}
	</nav>
</div>
