<script>
	let { data } = $props();
	import SectionHeader from '$lib/components/SectionHeader.svelte';
	import { resolve } from '$app/paths';

	function formatDate(date) {
		return new Date(date).toLocaleDateString('en-US', {
			year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
	}
</script>

<svelte:head>
	<title>Blog - Farhan Aulianda</title>
	<meta
		name="description"
		content="Thoughts on software engineering, cloud architecture, and cybersecurity."
	/>
</svelte:head>

<section class="relative py-32 px-6 w-full min-h-screen bg-zinc-950">
	<div class="max-w-7xl mx-auto relative z-10">
		<div class="mb-24">
			<SectionHeader title="Writing" subtitle="BLOG" />
			<p class="text-zinc-400 max-w-2xl mt-8 text-lg md:text-xl font-light leading-relaxed">
				Insights, tutorials, and thoughts on building scalable systems and secure infrastructure.
			</p>
		</div>

		<div class="grid grid-cols-1 md:grid-cols-2 gap-8">
			{#each data.posts as post (post.slug)}
				<div
					class="group relative flex flex-col justify-between h-full p-8 border border-zinc-900 bg-zinc-900/30 hover:border-zinc-700 transition-colors duration-500 rounded-2xl backdrop-blur-sm"
				>
					<div>
						<div class="flex flex-col gap-3 mb-6">
							<span class="text-xs font-mono text-zinc-500 uppercase tracking-widest"
								>{formatDate(post.meta.date)}</span
							>
							<div class="flex flex-wrap gap-2">
								{#if post.meta.tags}
									{#each post.meta.tags as tag (tag)}
										<span
											class="text-[10px] px-2 py-1 rounded-full border border-zinc-800 text-zinc-500 uppercase tracking-wider"
											>{tag}</span
										>
									{/each}
								{/if}
							</div>
						</div>

						<h2
							class="text-2xl font-bold text-white mb-4 group-hover:text-zinc-200 transition-colors font-outfit"
						>
							<a href={resolve(`/blog/${post.slug}`)} class="before:absolute before:inset-0">
								{post.meta.title}
							</a>
						</h2>

						<p class="text-zinc-400 mb-8 line-clamp-3 font-light leading-relaxed">
							{post.meta.description}
						</p>
					</div>

					<div
						class="flex items-center text-sm font-bold text-white uppercase tracking-widest group-hover:translate-x-2 transition-transform duration-300"
					>
						Read Article <span class="ml-2">→</span>
					</div>
				</div>
			{/each}
		</div>
	</div>
</section>
