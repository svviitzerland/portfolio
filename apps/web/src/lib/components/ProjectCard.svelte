<script>
	let { name, summary, stack, image, url, index } = $props();

	// Computed
	let tags = $derived(stack ? stack.split(',').map((t) => t.trim()) : []);
	let href = $derived(url ? (url.startsWith('http') ? url : `https://${url}`) : '#');

	import { resolve } from '$app/paths';
</script>

<a
	href={resolve(href)}
	target="_blank"
	rel="noopener noreferrer"
	class="group relative block w-full aspect-[4/3] rounded-2xl overflow-hidden bg-zinc-900 border border-zinc-800 transition-all duration-500 hover:scale-[1.02] hover:border-zinc-600 hover:shadow-2xl hover:shadow-white/5"
>
	<div
		class="absolute inset-0 bg-gradient-to-t from-zinc-950 via-zinc-950/50 to-transparent opacity-80 z-10 transition-opacity duration-300 group-hover:opacity-60"
	></div>

	<!-- Image Placeholder or Actual Image -->
	<div class="absolute inset-0 bg-zinc-800 transition-transform duration-700 group-hover:scale-110">
		{#if image}
			<img
				src={image}
				alt={name}
				class="w-full h-full object-cover grayscale group-hover:grayscale-0 transition-all duration-700"
			/>
		{:else}
			<div
				class="w-full h-full flex items-center justify-center text-zinc-700 text-4xl font-black opacity-20"
			>
				PROJECT {index + 1}
			</div>
		{/if}
	</div>

	<div
		class="absolute inset-0 z-20 flex flex-col justify-end p-8 translate-y-2 group-hover:translate-y-0 transition-transform duration-300"
	>
		<div
			class="flex flex-wrap gap-2 mb-4 opacity-0 transform translate-y-4 group-hover:opacity-100 group-hover:translate-y-0 transition-all duration-300 delay-100"
		>
			{#each tags as tag (tag)}
				<span
					class="px-2 py-1 text-xs font-semibold bg-white/10 backdrop-blur-md rounded-md text-white border border-white/10"
				>
					{tag}
				</span>
			{/each}
		</div>

		<h3 class="text-2xl font-bold text-white mb-2 group-hover:text-white transition-colors">
			{name}
		</h3>
		<p class="text-zinc-400 text-sm line-clamp-2 group-hover:text-zinc-300 transition-colors">
			{summary}
		</p>
	</div>
</a>
