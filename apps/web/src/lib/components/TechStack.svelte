<script>
	import cv from '$lib/data/cv.json';

	// Tech stack data
	const technologies = cv.skills;

	// Duplicate for infinite loop
	const displayTechs = [...technologies, ...technologies, ...technologies];
</script>

<div class="w-full py-12 overflow-hidden bg-zinc-950/50 backdrop-blur-sm border-y border-zinc-800">
	<div class="relative w-full flex items-center">
		<!-- Mask for fade effect on edges -->
		<div
			class="absolute inset-y-0 left-0 w-32 bg-gradient-to-r from-zinc-950 to-transparent z-10 pointer-events-none"
		></div>
		<div
			class="absolute inset-y-0 right-0 w-32 bg-gradient-to-l from-zinc-950 to-transparent z-10 pointer-events-none"
		></div>

		<div class="flex animate-marquee hover:pause whitespace-nowrap">
			{#each displayTechs as tech, i (i)}
				<div
					class="mx-6 px-6 py-3 rounded-xl bg-zinc-900/40 border border-zinc-800 backdrop-blur-md text-zinc-400 font-mono text-sm tracking-wider uppercase hover:text-white hover:border-zinc-600 transition-colors cursor-default"
				>
					{tech}
				</div>
			{/each}
		</div>

		<div
			class="flex animate-marquee-reverse hover:pause whitespace-nowrap absolute top-0 left-0 opacity-0 md:opacity-0"
			aria-hidden="true"
		>
			<!-- Hidden reverse layer for complex effects if needed later, currently using simple marquee -->
		</div>
	</div>
</div>

<style>
	.animate-marquee {
		animation: marquee 40s linear infinite;
	}

	.hover\:pause:hover {
		animation-play-state: paused;
	}

	@keyframes marquee {
		0% {
			transform: translateX(0);
		}
		100% {
			transform: translateX(-33.33%);
		} /* Move 1/3 since we tripled the array */
	}
</style>
