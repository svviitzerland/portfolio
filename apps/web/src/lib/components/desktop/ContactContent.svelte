<script>
	import cv from '$lib/data/cv.json';
	import MagneticButton from '$lib/components/MagneticButton.svelte';

	const { basics, socials } = cv;
	const { email, phone, website } = basics;
</script>

<div class="p-8 md:p-12 text-zinc-100">
	<!-- Header -->
	<div class="mb-12 text-center">
		<h1 class="text-4xl md:text-5xl font-black text-white tracking-tight mb-3">Get In Touch</h1>
		<p class="text-lg text-zinc-400 font-light max-w-lg mx-auto">
			Have a project in mind, or just want to say hi? I'd love to hear from you.
		</p>
	</div>

	<!-- Contact Info Cards -->
	<div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-12 max-w-3xl mx-auto">
		<!-- Email -->
		<a
			href="mailto:{email}"
			class="group p-6 bg-zinc-800/50 border border-zinc-700/50 rounded-xl hover:border-zinc-600 transition-all duration-300 text-center"
		>
			<div class="text-3xl mb-3">✉️</div>
			<p class="text-xs text-zinc-500 uppercase tracking-widest mb-1 font-mono">Email</p>
			<p class="text-sm text-zinc-300 group-hover:text-white transition-colors break-all">{email}</p>
		</a>

		<!-- Phone -->
		<a
			href="tel:{phone}"
			class="group p-6 bg-zinc-800/50 border border-zinc-700/50 rounded-xl hover:border-zinc-600 transition-all duration-300 text-center"
		>
			<div class="text-3xl mb-3">📱</div>
			<p class="text-xs text-zinc-500 uppercase tracking-widest mb-1 font-mono">Phone</p>
			<p class="text-sm text-zinc-300 group-hover:text-white transition-colors">{phone}</p>
		</a>

		<!-- Website -->
		<a
			href="https://{website}"
			target="_blank"
			rel="noopener noreferrer"
			class="group p-6 bg-zinc-800/50 border border-zinc-700/50 rounded-xl hover:border-zinc-600 transition-all duration-300 text-center"
		>
			<div class="text-3xl mb-3">🌐</div>
			<p class="text-xs text-zinc-500 uppercase tracking-widest mb-1 font-mono">Website</p>
			<p class="text-sm text-zinc-300 group-hover:text-white transition-colors">{website}</p>
		</a>
	</div>

	<!-- Social Links -->
	<div class="mb-12">
		<h2 class="text-xl font-bold text-white mb-6 text-center">Social Links</h2>
		<div class="grid grid-cols-2 md:grid-cols-3 gap-3 max-w-2xl mx-auto">
			{#each socials as social (social.network)}
				<a
					href={social.url}
					target="_blank"
					rel="noopener noreferrer"
					class="group flex items-center gap-3 p-4 bg-zinc-800/30 border border-zinc-700/50 rounded-xl hover:border-zinc-600 hover:bg-zinc-800/60 transition-all duration-300"
				>
					<span class="text-sm font-medium text-zinc-400 group-hover:text-white transition-colors">{social.network}</span>
					<span class="text-xs text-zinc-600 ml-auto">@{social.username}</span>
				</a>
			{/each}
		</div>
	</div>

	<!-- CTA -->
	<div class="flex justify-center">
		<MagneticButton>
			<a
				href="mailto:{email}"
				class="inline-flex items-center justify-center w-36 h-36 md:w-44 md:h-44 bg-white text-zinc-950 rounded-full text-lg font-bold uppercase tracking-widest hover:scale-105 transition-transform"
			>
				Say Hello
			</a>
		</MagneticButton>
	</div>
</div>
