<script>
	import { onMount } from 'svelte';

	let currentTime = $state('');

	function updateTime() {
		const now = new Date();
		currentTime = new Intl.DateTimeFormat('en-US', {
			weekday: 'short',
			month: 'short',
			day: 'numeric',
			hour: 'numeric',
			minute: '2-digit',
			hour12: true
		}).format(now);
	}

	onMount(() => {
		updateTime();
		const interval = setInterval(updateTime, 30000);
		return () => clearInterval(interval);
	});
</script>

<header
	class="fixed top-0 left-0 right-0 z-50 h-7 bg-zinc-900/80 backdrop-blur-xl border-b border-white/10 flex items-center justify-between px-4 select-none"
>
	<!-- Left: Apple logo + Name -->
	<div class="flex items-center gap-3">
		<span class="text-sm font-bold text-zinc-300"></span>
		<span class="text-[13px] font-semibold text-zinc-300 tracking-wide">Farhan Aulianda</span>
	</div>

	<!-- Right: Date/Time -->
	<div class="text-[12px] text-zinc-400 font-medium tabular-nums">
		{currentTime}
	</div>
</header>
