<script>
	import cv from '$lib/data/cv.json';

	const { basics, skills, education, experiences_in_organization } = cv;
	const { name, label, summary } = basics;
</script>

<div class="p-8 md:p-12 text-zinc-100">
	<!-- Header -->
	<div class="mb-10">
		<h1 class="text-4xl md:text-5xl font-black text-white tracking-tight mb-3">{name}</h1>
		<p class="text-lg md:text-xl text-zinc-400 font-light">{label}</p>
	</div>

	<!-- Summary -->
	<div class="mb-12">
		<p class="text-lg text-zinc-300 leading-relaxed max-w-3xl">{summary}</p>
	</div>

	<!-- Skills -->
	<div class="mb-12">
		<h2 class="text-xl font-bold text-white mb-6 flex items-center gap-2">
			<span class="text-zinc-500">⚡</span> Skills & Technologies
		</h2>
		<div class="flex flex-wrap gap-2">
			{#each skills as skill (skill)}
				<span
					class="px-3 py-1.5 text-sm bg-zinc-800 border border-zinc-700 rounded-lg text-zinc-300 hover:border-zinc-500 hover:text-white transition-colors duration-200"
				>
					{skill}
				</span>
			{/each}
		</div>
	</div>

	<!-- Education -->
	<div class="mb-12">
		<h2 class="text-xl font-bold text-white mb-6 flex items-center gap-2">
			<span class="text-zinc-500">🎓</span> Education
		</h2>
		<div class="space-y-4">
			{#each education as edu (edu.institution)}
				<div class="p-5 bg-zinc-800/50 border border-zinc-700/50 rounded-xl">
					<div class="flex flex-col md:flex-row md:items-center md:justify-between gap-1 mb-1">
						<h3 class="text-base font-semibold text-white">{edu.institution}</h3>
						<span class="text-xs font-mono text-zinc-500">{edu.startDate} — {edu.endDate}</span>
					</div>
					<p class="text-sm text-zinc-400">
						{edu.studyType} in {edu.area}
					</p>
				</div>
			{/each}
		</div>
	</div>

	<!-- Organizations -->
	<div>
		<h2 class="text-xl font-bold text-white mb-6 flex items-center gap-2">
			<span class="text-zinc-500">🏛️</span> Organizations
		</h2>
		<div class="space-y-4">
			{#each experiences_in_organization as org (org.organization + org.position)}
				<div class="p-5 bg-zinc-800/50 border border-zinc-700/50 rounded-xl">
					<div class="flex flex-col md:flex-row md:items-center md:justify-between gap-1 mb-1">
						<h3 class="text-base font-semibold text-white">{org.organization}</h3>
						<span class="text-xs font-mono text-zinc-500">{org.startDate} — {org.endDate}</span>
					</div>
					<p class="text-sm text-zinc-400 font-medium mb-1">{org.position}</p>
					{#if org.summary}
						<p class="text-sm text-zinc-500 leading-relaxed">{org.summary}</p>
					{/if}
				</div>
			{/each}
		</div>
	</div>
</div>
