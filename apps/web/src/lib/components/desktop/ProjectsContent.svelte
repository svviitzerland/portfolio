<script>
	import cv from '$lib/data/cv.json';

	let { search = '' } = $props();

	const { projects } = cv;

	const filtered = $derived(
		search
			? projects.filter((p) => {
					const q = search.toLowerCase();
					return (
						p.name.toLowerCase().includes(q) ||
						p.summary.toLowerCase().includes(q) ||
						(p.stack && p.stack.toLowerCase().includes(q))
					);
				})
			: projects
	);
</script>

<div class="p-8 md:p-12 text-zinc-100">
	<div class="space-y-2">
		{#each filtered as project (project.name)}
			{@const hasUrl = project.url}
			<a
				href={hasUrl ? (project.url.startsWith('http') ? project.url : `https://${project.url}`) : undefined}
				target={hasUrl ? '_blank' : undefined}
				rel={hasUrl ? 'noopener noreferrer' : undefined}
				class="group flex flex-col md:flex-row md:items-center gap-2 md:gap-6 p-5 rounded-xl hover:bg-zinc-800/50 transition-all duration-200 border border-transparent hover:border-zinc-700/50"
			>
				<!-- Date -->
				<span class="text-xs font-mono text-zinc-500 shrink-0 w-20">
					{project.date}
				</span>

				<!-- Content -->
				<div class="flex-1 min-w-0">
					<h3 class="text-base font-semibold text-white group-hover:text-zinc-200 mb-1 truncate">
						{project.name}
					</h3>
					<p class="text-sm text-zinc-500 line-clamp-1">{project.summary}</p>
				</div>

				<!-- Stack tags -->
				{#if project.stack}
					<div class="flex gap-1.5 shrink-0 flex-wrap md:flex-nowrap">
						{#each project.stack.split(', ').slice(0, 3) as tech (tech)}
							<span
								class="text-[10px] px-2 py-0.5 rounded-full border border-zinc-800 text-zinc-500 uppercase tracking-wider"
							>
								{tech}
							</span>
						{/each}
					</div>
				{/if}

				<!-- Arrow -->
				{#if hasUrl}
					<svg
						class="w-4 h-4 text-zinc-600 group-hover:text-zinc-400 shrink-0 hidden md:block group-hover:translate-x-1 transition-all"
						fill="none"
						stroke="currentColor"
						viewBox="0 0 24 24"
					>
						<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
					</svg>
				{/if}
			</a>
		{/each}

		{#if filtered.length === 0}
			<div class="text-center py-20 text-zinc-500">
				<p class="text-sm">No results for "{search}"</p>
			</div>
		{/if}
	</div>
</div>
