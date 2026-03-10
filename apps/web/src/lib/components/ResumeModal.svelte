<script>
	import { fade, fly } from 'svelte/transition';
	import { cubicOut } from 'svelte/easing';
	import { resolve } from '$app/paths';
	import ResumeContent from './ResumeContent.svelte';

	let { close } = $props();

	let canClose = $state(false);

	// Prevent body scroll when modal is open
	// Also add a small delay before backdrop click can close the window,
	// so the same click that opens the window doesn't immediately close it
	$effect(() => {
		document.body.style.overflow = 'hidden';
		canClose = false;
		const timer = setTimeout(() => { canClose = true; }, 100);
		return () => {
			document.body.style.overflow = '';
			clearTimeout(timer);
		};
	});

	function handleBackdropClick() {
		if (canClose) close();
	}
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin="anonymous" />
	<link
		href="https://fonts.googleapis.com/css2?family=Merriweather:ital,wght@0,300;0,400;0,700;0,900;1,300;1,400&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<!-- Backdrop -->
<div
	class="fixed inset-0 z-[60] bg-black/60 backdrop-blur-sm flex items-center justify-center p-0"
	transition:fade={{ duration: 200 }}
	onclick={handleBackdropClick}
	onkeydown={(e) => e.key === 'Escape' && close()}
	role="presentation"
>
	<!-- Modal Content - Resume Viewer -->
	<div
		class="bg-zinc-100 w-full max-w-[210mm] h-[90vh] overflow-hidden rounded-t-3xl md:rounded-2xl shadow-2xl relative mx-auto flex flex-col"
		style="background-color: #f4f4f5;"
		transition:fly={{ y: 1000, duration: 400, easing: cubicOut }}
		onclick={(e) => e.stopPropagation()}
		onkeydown={(e) => e.key === 'Enter' && e.stopPropagation()}
		role="dialog"
		aria-modal="true"
		tabindex="-1"
	>
		<!-- Floating Close Button -->
		<button
			onclick={(e) => {
				e.stopPropagation();
				close();
			}}
			class="absolute top-4 right-4 z-20 p-2 rounded-full transition-all"
			style="color: #a1a1aa; background-color: rgba(255, 255, 255, 0.5);"
			onmouseenter={(e) => {
				e.currentTarget.style.color = '#27272a';
				e.currentTarget.style.backgroundColor = 'rgba(255, 255, 255, 1)';
			}}
			onmouseleave={(e) => {
				e.currentTarget.style.color = '#a1a1aa';
				e.currentTarget.style.backgroundColor = 'rgba(255, 255, 255, 0.5)';
			}}
			aria-label="Close modal"
		>
			<svg class="h-6 w-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
				<path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M6 18L18 6M6 6l12 12"
				/>
			</svg>
		</button>

		<!-- Resume Content - Rendered directly -->
		<div class="w-full h-full overflow-y-auto overflow-x-hidden bg-white resume-scroll">
			<ResumeContent />
		</div>
	</div>

	<!-- Floating Download Button -->
	<a
		href={resolve('/resume.pdf')}
		download="Resume - Farhan Aulianda.pdf"
		onclick={(e) => e.stopPropagation()}
		class="fixed bottom-8 right-8 z-[70] flex items-center gap-2 px-6 py-3 rounded-full shadow-2xl hover:scale-105 transition-transform font-sans font-bold"
		style="background-color: #18181b; color: #ffffff;"
		aria-label="Download Resume"
		transition:fly={{ y: 100, duration: 300, delay: 200 }}
	>
		<svg class="h-5 w-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
			<path
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4"
			/>
		</svg>
		Download PDF
	</a>
</div>

<style>
	.resume-scroll::-webkit-scrollbar {
		width: 6px;
	}

	.resume-scroll::-webkit-scrollbar-track {
		background: transparent;
	}

	.resume-scroll::-webkit-scrollbar-thumb {
		background-color: #5e4b35;
		border-radius: 3px;
	}

	.resume-scroll::-webkit-scrollbar-thumb:hover {
		background-color: #6d5e41;
	}
</style>
