<script>
	let { data } = $props();
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
	<title>{data.meta.title} - Farhan Aulianda</title>
	<meta property="og:type" content="article" />
	<meta property="og:title" content={data.meta.title} />
</svelte:head>

<article class="relative pt-16 pb-32 px-6 w-full min-h-screen bg-zinc-950">
	<div class="max-w-4xl mx-auto relative z-10">
		<!-- Header -->
		<header class="mb-16 text-center">
			<div class="flex flex-col items-center gap-4 mb-6">
				<span class="text-xs font-mono text-zinc-500 uppercase tracking-widest"
					>{formatDate(data.meta.date)}</span
				>
				{#if data.meta.tags}
					<div class="flex flex-wrap justify-center gap-2">
						{#each data.meta.tags as tag (tag)}
							<span
								class="text-[10px] px-2 py-1 rounded-full border border-zinc-800 text-zinc-500 uppercase tracking-wider"
								>{tag}</span
							>
						{/each}
					</div>
				{/if}
			</div>

			<h1 class="text-4xl md:text-6xl font-black text-white leading-tight mb-8 font-outfit">
				{data.meta.title}
			</h1>

			<p class="text-xl text-zinc-400 font-light max-w-2xl mx-auto leading-relaxed">
				{data.meta.description}
			</p>
		</header>

		<!-- Content -->
		<div class="blog-content prose prose-invert prose-zinc max-w-none prose-lg prose-headings:font-outfit prose-headings:font-bold prose-p:font-light prose-p:leading-loose prose-p:text-zinc-300 prose-a:text-blue-400 prose-a:no-underline hover:prose-a:text-blue-300 prose-strong:text-white prose-strong:font-semibold prose-code:text-emerald-400 prose-code:font-mono prose-code:text-sm prose-ul:text-zinc-300 prose-ol:text-zinc-300 prose-li:text-zinc-300 prose-img:rounded-xl prose-img:border prose-img:border-zinc-800">
			<data.content />
		</div>

		<!-- Footer / Navigation -->
		<div class="mt-24 pt-12 border-t border-zinc-900">
			<!-- Previous/Next Navigation -->
			{#if data.prevPost || data.nextPost}
				<div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-12">
					<!-- Previous Post -->
					{#if data.prevPost}
						<a
							href={resolve(`/blog/${data.prevPost.slug}`)}
							class="group flex flex-col p-6 border border-zinc-900 bg-zinc-900/30 hover:border-zinc-700 transition-all duration-300 rounded-xl"
						>
							<span class="text-xs font-mono text-zinc-500 uppercase tracking-widest mb-3"
								>← Previous</span
							>
							<h3
								class="text-lg font-bold text-white group-hover:text-zinc-300 transition-colors font-outfit line-clamp-2"
							>
								{data.prevPost.meta.title}
							</h3>
						</a>
					{:else}
						<div></div>
					{/if}

					<!-- Next Post -->
					{#if data.nextPost}
						<a
							href={resolve(`/blog/${data.nextPost.slug}`)}
							class="group flex flex-col p-6 border border-zinc-900 bg-zinc-900/30 hover:border-zinc-700 transition-all duration-300 rounded-xl text-right"
						>
							<span class="text-xs font-mono text-zinc-500 uppercase tracking-widest mb-3"
								>Next →</span
							>
							<h3
								class="text-lg font-bold text-white group-hover:text-zinc-300 transition-colors font-outfit line-clamp-2"
							>
								{data.nextPost.meta.title}
							</h3>
						</a>
					{/if}
				</div>
			{/if}

			<!-- Back to Blog -->
			<div class="text-center">
				<a
					href={resolve('/?open=blog')}
					class="inline-flex items-center text-sm font-bold text-white uppercase tracking-widest hover:text-zinc-400 transition-colors"
				>
					<span class="mr-2">←</span> Back to Blog
				</a>
			</div>
		</div>
	</div>
</article>

<style>
	/* Code Block Styling */
	.blog-content :global(pre) {
		position: relative;
		background-color: rgb(24 24 27);
		border: 1px solid rgb(39 39 42);
		border-radius: 0.75rem;
		padding: 0;
		margin: 2rem 0;
		overflow: hidden;
		counter-reset: line;
	}

	.blog-content :global(pre code) {
		display: block;
		padding: 1.5rem;
		overflow-x: auto;
		font-size: 0.875rem;
		line-height: 1.75;
		color: rgb(212 212 216);
		font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
	}

	/* Add language label */
	.blog-content :global(pre[class*='language-'])::before {
		content: attr(class);
		position: absolute;
		top: 0;
		right: 0;
		padding: 0.5rem 1rem;
		font-size: 0.75rem;
		font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, monospace;
		color: rgb(161 161 170);
		text-transform: uppercase;
		letter-spacing: 0.05em;
		background-color: rgba(39, 39, 42, 0.5);
		border-bottom-left-radius: 0.5rem;
		border-left: 1px solid rgb(63 63 70);
		border-bottom: 1px solid rgb(63 63 70);
	}

	.blog-content :global(pre.language-bash)::before {
		content: 'bash';
	}

	.blog-content :global(pre.language-javascript)::before {
		content: 'javascript';
	}

	.blog-content :global(pre.language-typescript)::before {
		content: 'typescript';
	}

	.blog-content :global(pre.language-python)::before {
		content: 'python';
	}

	.blog-content :global(pre.language-hcl)::before {
		content: 'terraform';
	}

	.blog-content :global(pre.language-json)::before {
		content: 'json';
	}

	.blog-content :global(pre.language-rust)::before {
		content: 'rust';
	}

	/* Inline code */
	.blog-content :global(:not(pre) > code) {
		padding: 0.25rem 0.5rem;
		border-radius: 0.25rem;
		background-color: rgba(39, 39, 42, 0.5);
		border: 1px solid rgb(63 63 70);
	}
</style>
