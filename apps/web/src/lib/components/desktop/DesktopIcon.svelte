<script>
	import { fly } from 'svelte/transition';

	let { label, onclick, children } = $props();
</script>

<button
	class="group flex flex-col items-center gap-1.5 w-[76px] p-2 rounded-lg hover:bg-white/10 transition-all duration-200 cursor-default select-none"
	onclick={onclick}
	in:fly={{ y: 10, duration: 300 }}
>
	<!-- Icon container -->
	<div class="w-14 h-14 flex items-center justify-center rounded-[14px] bg-zinc-800/80 border border-zinc-700/60 shadow-lg group-hover:scale-105 group-hover:border-zinc-600 transition-all duration-200">
		{@render children()}
	</div>

	<span
		class="text-[11px] text-white/90 text-center leading-tight font-medium max-w-full truncate"
		style="text-shadow: 0 1px 4px rgba(0,0,0,0.9), 0 0 10px rgba(0,0,0,0.5);"
	>
		{label}
	</span>
</button>
