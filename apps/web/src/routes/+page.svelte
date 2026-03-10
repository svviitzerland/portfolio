<script>
	import SectionHeader from '$lib/components/SectionHeader.svelte';
	import ParallaxText from '$lib/components/ParallaxText.svelte';
	import MagneticButton from '$lib/components/MagneticButton.svelte';
	import ResumeModal from '$lib/components/ResumeModal.svelte';
	import { resolve } from '$app/paths';

	// Import data directly
	import cv from '$lib/data/cv.json';

	let { data } = $props();

	let scrollY = $state(0);
	let innerHeight = $state(0);
	let isContactExpanded = $state(false);
	let showResume = $state(false);

	// Extract data for easier access
	const { basics, socials } = cv;
	const { name, label, summary, location } = basics;

	// Get geo data
	const geo = $derived(data?.geo || {});

	// Create greeting message
	const greeting = $derived.by(() => {
		if (geo.countryName) {
			return `Hi Wanderer from ${geo.countryName}!`;
		} else if (geo.country) {
			return `Hi Wanderer from ${geo.country}!`;
		}
		return 'Hi there!';
	});

	// Helper to get location string
	const locationString = [location.city, location.countryCode].filter(Boolean).join(', ');

	function handleContactClick() {
		isContactExpanded = true;
		setTimeout(() => {
			window.location.href = `mailto:${basics.email}`;
			setTimeout(() => {
				isContactExpanded = false;
			}, 1000);
		}, 800);
	}
</script>

<svelte:window bind:scrollY bind:innerHeight />

<!-- Resume Modal -->
{#if showResume}
	<ResumeModal close={() => (showResume = false)} />
{/if}

<main
	class="bg-zinc-950 min-h-screen text-zinc-100 selection:bg-white selection:text-zinc-950 overflow-x-hidden"
>
	<!-- Hero / Intro -->
	<section class="min-h-screen flex flex-col justify-center items-center relative px-6 py-20">
		<div
			class="absolute inset-0 bg-[radial-gradient(circle_at_center,_var(--tw-gradient-stops))] from-zinc-900/40 via-zinc-950 to-zinc-950 z-0 pointer-events-none"
		></div>

		<div class="z-10 max-w-5xl mx-auto w-full">
			{#if greeting}
				<ParallaxText speed={0.15}>
					<div class="mb-6 text-zinc-400 font-mono text-sm md:text-base tracking-wide">
						👋 {greeting}
					</div>
				</ParallaxText>
			{/if}

			<ParallaxText speed={0.2}>
				<div
					class="mb-4 flex items-center gap-2 text-zinc-500 font-mono text-xs md:text-sm tracking-widest uppercase"
				>
					<div class="w-12 h-[1px] bg-zinc-700"></div>
					{locationString || 'Earth'}
				</div>
			</ParallaxText>

			<ParallaxText speed={0.4}>
				<h1
					class="text-[clamp(3.5rem,9vw,9rem)] font-black leading-[0.85] tracking-tighter text-white mb-8"
				>
					{#each name.split(' ') as n (n)}
						<span class="block">{n}</span>
					{/each}
				</h1>
			</ParallaxText>

			<ParallaxText speed={0.3}>
				<p class="text-xl md:text-3xl font-light text-zinc-400 max-w-2xl leading-relaxed mb-8">
					{label}
				</p>
			</ParallaxText>

			<ParallaxText speed={0.3}>
				<button
					onclick={() => (showResume = true)}
					class="group relative inline-flex items-center gap-3 px-10 py-4 bg-transparent overflow-hidden rounded-full transition-all duration-500 hover:scale-[1.02]"
				>
					<!-- Elegant Border with gradient -->
					<div
						class="absolute inset-0 rounded-full border border-zinc-700 group-hover:border-zinc-500 transition-all duration-500"
					></div>

					<!-- Subtle animated background -->
					<div
						class="absolute inset-0 rounded-full opacity-0 group-hover:opacity-100 bg-gradient-to-r from-zinc-800/50 via-zinc-700/30 to-zinc-800/50 transition-opacity duration-500"
					></div>

					<!-- Glow effect -->
					<div
						class="absolute inset-0 rounded-full opacity-0 group-hover:opacity-20 bg-white blur-2xl transition-all duration-500"
					></div>

					<span
						class="relative z-10 font-mono text-zinc-400 group-hover:text-white uppercase tracking-[0.2em] text-sm font-medium transition-colors duration-300"
						>View Resume</span
					>

					<!-- Arrow icon instead of dot -->
					<svg
						class="relative z-10 w-4 h-4 text-zinc-600 group-hover:text-white group-hover:translate-x-1 transition-all duration-300"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M9 5l7 7-7 7"
						/>
					</svg>
				</button>
			</ParallaxText>
		</div>

		<!-- Scroll hint -->
		<div class="absolute bottom-12 flex flex-col items-center gap-2 opacity-50 animate-pulse">
			<span class="text-[10px] font-mono uppercase tracking-widest text-zinc-600">The Story</span>
			<div class="h-12 w-[1px] bg-gradient-to-b from-zinc-600 to-transparent"></div>
		</div>
	</section>

	<!-- The Story (Summary) -->
	<section id="prologue" class="py-32 px-6 border-t border-zinc-900">
		<div class="max-w-4xl mx-auto">
			<SectionHeader title="The Prologue" subtitle="WHO I AM" />

			<ParallaxText speed={0.1}>
				<p class="text-2xl md:text-4xl leading-tight font-light text-zinc-300">
					{summary}
				</p>
			</ParallaxText>
		</div>
	</section>

	<!-- Connect -->
	<section
		id="contact"
		class="py-40 px-6 flex flex-col items-center justify-center text-center relative overflow-hidden"
	>
		<SectionHeader title="Epilogue" subtitle="LET'S CONNECT" class="items-center" />

		<h2 class="text-5xl md:text-8xl font-black text-white mb-12 tracking-tighter">
			Start a new<br /><span class="text-zinc-600">chapter.</span>
		</h2>

		<MagneticButton>
			<button
				onclick={handleContactClick}
				class="relative z-20 w-40 h-40 md:w-56 md:h-56 bg-white text-zinc-950 rounded-full text-xl font-bold uppercase tracking-widest hover:scale-105 transition-transform flex items-center justify-center"
			>
				Say Hello
			</button>
		</MagneticButton>

		<div class="mt-20 flex flex-wrap justify-center gap-8">
			{#each socials as social (social.network)}
				<a
					href={social.url.startsWith('http') ? social.url : resolve(social.url)}
					target="_blank"
					class="text-zinc-500 hover:text-white transition-colors text-sm font-mono uppercase tracking-widest border-b border-transparent hover:border-white pb-1"
				>
					{social.network}
				</a>
			{/each}
		</div>
	</section>

	<footer class="py-12 text-center text-zinc-700 text-xs uppercase tracking-widest">
		&copy; {new Date().getFullYear()}
		{name}. Built with SvelteKit.
	</footer>
</main>

<!-- Contact Expansion Overlay -->
<div
	class="fixed inset-0 z-50 pointer-events-none flex items-center justify-center overflow-hidden"
	aria-hidden="true"
>
	<div
		class="w-[10vw] h-[10vw] rounded-full bg-white transition-transform duration-1000 ease-[cubic-bezier(0.87,0,0.13,1)] will-change-transform"
		class:scale-[30]={isContactExpanded}
		class:scale-0={!isContactExpanded}
	></div>
</div>
