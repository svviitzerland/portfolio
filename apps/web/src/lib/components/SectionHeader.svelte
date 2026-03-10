<script>
	import { cn } from '$lib/utils';

	let { title, subtitle, class: className = '' } = $props();
</script>

<div class={cn('flex flex-col gap-2 mb-12', className)}>
	<span class="text-sm font-medium tracking-widest text-zinc-500 uppercase">{subtitle}</span>
	<h2 class="text-3xl md:text-5xl font-bold tracking-tight text-white">{title}</h2>
	<div class="h-1 w-20 bg-gradient-to-r from-zinc-500 to-transparent mt-4 rounded-full"></div>
</div>
