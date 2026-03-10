<script>
	import { resolve } from '$app/paths';

	let { posts = [] } = $props();

	function formatDate(date) {
		return new Date(date).toLocaleDateString('en-US', {
			year: 'numeric',
			month: 'short',
			day: 'numeric'
		});
	}
</script>

<div class="p-8 md:p-12 text-zinc-100">
	<h1 class="text-3xl font-black text-white tracking-tight mb-2">Blog</h1>
	<p class="text-zinc-400 mb-10 font-light">
		Insights, tutorials, and thoughts on building scalable systems.
	</p>

	<div class="space-y-2">
		{#each posts as post (post.slug)}
			<a
				href={resolve(`/blog/${post.slug}`)}
				class="group flex flex-col md:flex-row md:items-center gap-2 md:gap-6 p-5 rounded-xl hover:bg-zinc-800/50 transition-all duration-200 border border-transparent hover:border-zinc-700/50"
			>
				<!-- Date -->
				<span class="text-xs font-mono text-zinc-500 shrink-0 w-28">
					{formatDate(post.meta.date)}
				</span>

				<!-- Content -->
				<div class="flex-1 min-w-0">
					<h3 class="text-base font-semibold text-white group-hover:text-zinc-200 mb-1 truncate">
						{post.meta.title}
					</h3>
					{#if post.meta.description}
						<p class="text-sm text-zinc-500 line-clamp-1">{post.meta.description}</p>
					{/if}
				</div>

				<!-- Tags -->
				{#if post.meta.tags}
					<div class="flex gap-1.5 shrink-0">
						{#each post.meta.tags.slice(0, 3) as tag (tag)}
							<span
								class="text-[10px] px-2 py-0.5 rounded-full border border-zinc-800 text-zinc-500 uppercase tracking-wider"
							>
								{tag}
							</span>
						{/each}
					</div>
				{/if}

				<!-- Arrow -->
				<svg
					class="w-4 h-4 text-zinc-600 group-hover:text-zinc-400 shrink-0 hidden md:block group-hover:translate-x-1 transition-all"
					fill="none"
					stroke="currentColor"
					viewBox="0 0 24 24"
				>
					<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
				</svg>
			</a>
		{/each}

		{#if posts.length === 0}
			<div class="text-center py-20 text-zinc-500">
				<p class="text-lg">No posts yet.</p>
			</div>
		{/if}
	</div>
</div>
