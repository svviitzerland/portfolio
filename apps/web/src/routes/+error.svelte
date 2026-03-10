<script>
	import { resolve } from '$app/paths';
	import ParallaxText from '$lib/components/ParallaxText.svelte';
	import MagneticButton from '$lib/components/MagneticButton.svelte';
	import { onMount } from 'svelte';

	let scrollY = $state(0);
	let mouseX = $state(0);
	let mouseY = $state(0);
	let glitchActive = $state(false);

	function handleMouseMove(e) {
		mouseX = (e.clientX / window.innerWidth - 0.5) * 20;
		mouseY = (e.clientY / window.innerHeight - 0.5) * 20;
	}

	onMount(() => {
		// Random glitch effect
		const glitchInterval = setInterval(
			() => {
				glitchActive = true;
				setTimeout(() => {
					glitchActive = false;
				}, 200);
			},
			3000 + Math.random() * 2000
		);

		return () => clearInterval(glitchInterval);
	});
</script>

<svelte:window bind:scrollY on:mousemove={handleMouseMove} />

<main
	class="min-h-screen bg-zinc-950 text-zinc-100 flex flex-col items-center justify-center relative overflow-hidden px-6 py-20"
>
	<!-- Animated background gradient -->
	<div
		class="absolute inset-0 bg-[radial-gradient(circle_at_center,_var(--tw-gradient-stops))] from-zinc-900/40 via-zinc-950 to-zinc-950 pointer-events-none"
		style="transform: translate({mouseX}px, {mouseY}px)"
	></div>

	<!-- Floating grid pattern -->
	<div
		class="absolute inset-0 opacity-[0.02] pointer-events-none"
		style="background-image: linear-gradient(rgba(255,255,255,0.1) 1px, transparent 1px), linear-gradient(90deg, rgba(255,255,255,0.1) 1px, transparent 1px); background-size: 50px 50px;"
	></div>

	<div class="relative z-10 max-w-4xl mx-auto text-center">
		<!-- Large 404 with glitch effect -->
		<ParallaxText speed={0.5}>
			<div class="relative mb-12">
				<h1
					class="text-[clamp(8rem,25vw,20rem)] font-black leading-none tracking-tighter select-none relative"
					class:glitch={glitchActive}
				>
					<span
						class="absolute inset-0 text-transparent bg-clip-text bg-gradient-to-br from-emerald-400 via-cyan-400 to-blue-500 opacity-30 blur-sm"
						>404</span
					>
					<span class="relative text-white drop-shadow-[0_0_30px_rgba(255,255,255,0.3)]">404</span>
				</h1>

				<!-- Animated scan line -->
				<div class="scan-line"></div>
			</div>
		</ParallaxText>

		<!-- Title with typing effect -->
		<ParallaxText speed={0.4}>
			<h2 class="text-2xl md:text-4xl font-light text-zinc-300 mb-6 tracking-wide">
				<span class="font-mono text-emerald-400">&gt;</span> Page Not Found
			</h2>
		</ParallaxText>

		<!-- Description -->
		<ParallaxText speed={0.3}>
			<p
				class="text-base md:text-lg text-zinc-500 max-w-xl mx-auto leading-relaxed mb-12 font-mono"
			>
				// The page you're looking for doesn't exist.<br />
				<span class="text-zinc-600">// Let's navigate you back home.</span>
			</p>
		</ParallaxText>

		<!-- Action buttons -->
		<ParallaxText speed={0.2}>
			<div class="flex flex-col sm:flex-row items-center justify-center gap-4">
				<MagneticButton>
					<a
						href={resolve('/')}
						class="group relative inline-flex items-center gap-3 px-10 py-5 bg-white text-zinc-950 font-mono text-sm uppercase tracking-[0.2em] font-semibold hover:scale-105 transition-all overflow-hidden shadow-[0_0_30px_rgba(255,255,255,0.1)] hover:shadow-[0_0_50px_rgba(255,255,255,0.2)]"
					>
						<span class="relative z-10">Return Home</span>
						<svg
							class="relative z-10 w-4 h-4 group-hover:translate-x-1 transition-transform"
							fill="none"
							stroke="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								stroke-width="2"
								d="M13 7l5 5m0 0l-5 5m5-5H6"
							/>
						</svg>
					</a>
				</MagneticButton>

				<button
					onclick={() => window.history.back()}
					class="group relative inline-flex items-center gap-3 px-10 py-5 bg-transparent font-mono text-sm uppercase tracking-[0.2em] font-medium hover:scale-105 transition-all overflow-hidden border border-zinc-800 hover:border-zinc-600"
				>
					<svg
						class="relative z-10 w-4 h-4 text-zinc-600 group-hover:text-white group-hover:-translate-x-1 transition-all"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M11 17l-5-5m0 0l5-5m-5 5h12"
						/>
					</svg>
					<span class="relative z-10 text-zinc-400 group-hover:text-white transition-colors"
						>Go Back</span
					>
				</button>
			</div>
		</ParallaxText>
	</div>

	<!-- Floating particles -->
	<div class="absolute inset-0 pointer-events-none overflow-hidden">
		<!-- eslint-disable-next-line no-unused-vars -->
		{#each Array(20) as _, idx (idx)}
			<div
				class="absolute w-1 h-1 bg-zinc-700 rounded-full opacity-20"
				style="
					left: {Math.random() * 100}%;
					top: {Math.random() * 100}%;
					animation: float {5 + Math.random() * 10}s ease-in-out infinite;
					animation-delay: {Math.random() * 5}s;
				"
			></div>
		{/each}
	</div>
</main>

<style>
	@keyframes float {
		0%,
		100% {
			transform: translateY(0) translateX(0);
			opacity: 0.2;
		}
		50% {
			transform: translateY(-30px) translateX(20px);
			opacity: 0.4;
		}
	}

	/* Glitch effect */
	.glitch {
		animation: glitch-anim 0.3s cubic-bezier(0.25, 0.46, 0.45, 0.94) both;
	}

	@keyframes glitch-anim {
		0% {
			transform: translate(0);
		}
		20% {
			transform: translate(-3px, 3px);
		}
		40% {
			transform: translate(-3px, -3px);
		}
		60% {
			transform: translate(3px, 3px);
		}
		80% {
			transform: translate(3px, -3px);
		}
		100% {
			transform: translate(0);
		}
	}

	/* Scan line animation */
	.scan-line {
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		height: 2px;
		background: linear-gradient(
			90deg,
			transparent,
			rgba(16, 185, 129, 0.5),
			rgba(6, 182, 212, 0.5),
			transparent
		);
		animation: scan 3s ease-in-out infinite;
		filter: blur(1px);
	}

	@keyframes scan {
		0%,
		100% {
			top: 0;
			opacity: 0;
		}
		10% {
			opacity: 1;
		}
		90% {
			opacity: 1;
		}
		100% {
			top: 100%;
			opacity: 0;
		}
	}
</style>
