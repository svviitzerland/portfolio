<script>
	import cv from '$lib/data/cv.json';

	const { projects } = cv;
</script>

<div class="p-8 md:p-12 text-zinc-100">
	<h1 class="text-3xl font-black text-white tracking-tight mb-2">Projects</h1>
	<p class="text-zinc-400 mb-10 font-light">Things I've built from zero to one.</p>

	<div class="grid grid-cols-1 md:grid-cols-2 gap-5">
		{#each projects as project (project.name)}
			<div
				class="group p-6 bg-zinc-800/50 border border-zinc-700/50 rounded-xl hover:border-zinc-600 transition-all duration-300"
			>
				<div class="flex items-start justify-between gap-3 mb-3">
					<h3 class="text-lg font-bold text-white group-hover:text-zinc-100">{project.name}</h3>
					<span class="text-xs font-mono text-zinc-500 whitespace-nowrap shrink-0">{project.date}</span>
				</div>

				<p class="text-sm text-zinc-400 leading-relaxed mb-4 line-clamp-3">{project.summary}</p>

				{#if project.stack}
					<div class="flex flex-wrap gap-1.5 mb-4">
						{#each project.stack.split(', ') as tech (tech)}
							<span class="text-[10px] px-2 py-0.5 rounded-full bg-zinc-700/50 text-zinc-400 border border-zinc-700">
								{tech}
							</span>
						{/each}
					</div>
				{/if}

				{#if project.url}
					<a
						href={project.url.startsWith('http') ? project.url : `https://${project.url}`}
						target="_blank"
						rel="noopener noreferrer"
						class="inline-flex items-center gap-1 text-xs font-medium text-zinc-400 hover:text-white transition-colors"
					>
						<svg class="w-3.5 h-3.5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
							<path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14" />
						</svg>
						{project.url}
					</a>
				{/if}
			</div>
		{/each}
	</div>
</div>
