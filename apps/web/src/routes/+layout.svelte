<script>
	import '@fontsource/inter/400.css';
	import '@fontsource/inter/600.css';
	import '@fontsource/outfit/700.css';
	import '@fontsource/outfit/900.css';
	import './layout.css';
	import IntroAnimation from '$lib/components/IntroAnimation.svelte';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { spring } from 'svelte/motion';
	import { browser } from '$app/environment';
	import { resolve } from '$app/paths';

	let { children } = $props();

	// Intro animation state
	let showIntro = $state(false);
	let introComplete = $state(false);

	onMount(() => {
		// Only show intro on homepage and if not shown before in this session
		if ($page.url.pathname === '/' && !sessionStorage.getItem('intro_shown')) {
			showIntro = true;
		} else {
			introComplete = true;
		}
	});

	function handleIntroComplete() {
		introComplete = true;
		if (browser) {
			sessionStorage.setItem('intro_shown', 'true');
		}
	}

	// Custom Cursor Logic
	let coords = spring({ x: 0, y: 0 }, { stiffness: 0.05, damping: 0.25 });
	let size = spring(10, { stiffness: 0.1, damping: 0.25 });

	function handleMouseMove(e) {
		coords.set({ x: e.clientX, y: e.clientY });
	}

	function handleMouseDown() {
		size.set(20);
	}

	function handleMouseUp() {
		size.set(10);
	}
</script>

<svelte:window
	on:mousemove={handleMouseMove}
	on:mousedown={handleMouseDown}
	on:mouseup={handleMouseUp}
/>

<!-- Custom Cursor -->
<div
	class="hidden md:block fixed top-0 left-0 pointer-events-none z-[100] mix-blend-difference"
	style="transform: translate({$coords.x}px, {$coords.y}px)"
>
	<div
		style="width: {$size}px; height: {$size}px;"
		class="bg-white rounded-full -translate-x-1/2 -translate-y-1/2"
	></div>
</div>

<!-- Intro Animation -->
{#if showIntro}
	<IntroAnimation onComplete={handleIntroComplete} />
{/if}

<div
	class="font-sans antialiased text-zinc-50 selection:bg-zinc-300 selection:text-zinc-900 bg-zinc-950 min-h-screen flex flex-col items-center overflow-x-hidden selection:bg-white/20 transition-opacity duration-500"
	class:opacity-0={showIntro && !introComplete}
>
	{#if $page.url.pathname !== '/'}
	<nav
		class="fixed top-6 z-50 px-6 py-3 rounded-full bg-zinc-900/50 backdrop-blur-md border border-white/10 shadow-xl flex items-center gap-6 saturate-150 transition-all hover:bg-zinc-900/70 hover:scale-105 duration-300"
	>
		<a
			href={resolve('/')}
			class="text-sm font-medium text-zinc-300 hover:text-white transition-colors tracking-wide"
			>Home</a
		>
		<a
			href={resolve('/#prologue')}
			class="text-sm font-medium text-zinc-300 hover:text-white transition-colors tracking-wide"
			>About</a
		>
		<a
			href={resolve('/blog')}
			class="text-sm font-medium text-zinc-300 hover:text-white transition-colors tracking-wide"
			>Blog</a
		>
		<a
			href={resolve('/#contact')}
			class="text-sm font-medium text-zinc-300 hover:text-white transition-colors tracking-wide"
			>Contact</a
		>
	</nav>
	{/if}

	{@render children()}
</div>

<style>
	:global(html) {
		scroll-behavior: smooth;
		background-color: #09090b; /* zinc-950 */
	}
	:global(body) {
		font-family: 'Inter', sans-serif;
	}
	:global(h1, h2, h3, h4, h5, h6) {
		font-family: 'Outfit', sans-serif;
	}
</style>
