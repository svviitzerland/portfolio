<script>
	import { fade, fly } from 'svelte/transition';
	import { cubicOut } from 'svelte/easing';

	let { title, close, children } = $props();

	let search = $state('');
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

	function handleKeydown(e) {
		if (e.key === 'Escape') close();
	}
</script>

<svelte:window onkeydown={handleKeydown} />

<!-- Backdrop -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	class="fixed inset-0 z-[60] bg-black/60 backdrop-blur-sm flex items-center justify-center p-4"
	transition:fade={{ duration: 200 }}
	onclick={handleBackdropClick}
	onkeydown={(e) => e.key === 'Escape' && close()}
	role="presentation"
>
	<!-- Window -->
	<div
		class="bg-zinc-900 w-full max-w-5xl h-[85vh] rounded-xl shadow-2xl overflow-hidden flex flex-col"
		transition:fly={{ y: 1000, duration: 400, easing: cubicOut }}
		onclick={(e) => e.stopPropagation()}
		onkeydown={(e) => e.stopPropagation()}
		role="dialog"
		aria-modal="true"
		aria-label={title}
		tabindex="-1"
	>
		<!-- Title Bar -->
		<div
			class="h-12 bg-zinc-800 flex items-center px-4 shrink-0 border-b border-zinc-700/50"
		>
			<!-- Traffic Lights -->
			<div class="flex items-center gap-2">
				<button
					onclick={(e) => {
						e.stopPropagation();
						close();
					}}
					class="w-3 h-3 rounded-full bg-[#ff5f57] hover:brightness-110 transition-all group relative"
					aria-label="Close window"
				>
					<svg
						class="w-3 h-3 opacity-0 group-hover:opacity-100 text-black/60 absolute inset-0"
						viewBox="0 0 12 12"
						fill="currentColor"
					>
						<path d="M3.5 3.5l5 5M8.5 3.5l-5 5" stroke="currentColor" stroke-width="1.5" fill="none" />
					</svg>
				</button>
				<div class="w-3 h-3 rounded-full bg-[#febc2e]"></div>
				<div class="w-3 h-3 rounded-full bg-[#28c840]"></div>
			</div>

			<!-- Centered Title -->
			<div class="flex-1 text-center">
				<span class="text-[13px] font-medium text-zinc-400">{title}</span>
			</div>

			<!-- Search -->
			<div class="relative w-44">
				<svg class="absolute left-2 top-1/2 -translate-y-1/2 w-3 h-3 text-zinc-500" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
					<path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
				</svg>
				<input
					type="text"
					placeholder="Search"
					bind:value={search}
					class="w-full h-6 pl-7 pr-2 text-[12px] bg-zinc-700/50 border border-zinc-600/50 rounded-md text-zinc-200 placeholder:text-zinc-500 outline-none focus:border-zinc-500 focus:bg-zinc-700/80 transition-colors"
				/>
			</div>
		</div>

		<!-- Content Area -->
		<div class="flex-1 overflow-y-auto overflow-x-hidden window-scroll">
			{@render children(search)}
		</div>
	</div>
</div>

<style>
	.window-scroll::-webkit-scrollbar {
		width: 6px;
	}

	.window-scroll::-webkit-scrollbar-track {
		background: transparent;
	}

	.window-scroll::-webkit-scrollbar-thumb {
		background-color: #52525b;
		border-radius: 3px;
	}

	.window-scroll::-webkit-scrollbar-thumb:hover {
		background-color: #71717a;
	}
</style>
