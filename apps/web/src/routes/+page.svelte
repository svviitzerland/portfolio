<script>
	import MenuBar from '$lib/components/desktop/MenuBar.svelte';
	import DesktopIcon from '$lib/components/desktop/DesktopIcon.svelte';
	import DesktopWindow from '$lib/components/desktop/DesktopWindow.svelte';
	import Dock from '$lib/components/desktop/Dock.svelte';
	import AboutMeContent from '$lib/components/desktop/AboutMeContent.svelte';
	import ProjectsContent from '$lib/components/desktop/ProjectsContent.svelte';
	import BlogContent from '$lib/components/desktop/BlogContent.svelte';
	import ContactContent from '$lib/components/desktop/ContactContent.svelte';
	import ResumeModal from '$lib/components/ResumeModal.svelte';

	let { data } = $props();

	/** @type {null | 'about' | 'projects' | 'blog' | 'resume' | 'contact'} */
	let activeWindow = $state(null);

	const posts = $derived(data?.posts || []);

	function openWindow(name) {
		activeWindow = name;
	}

	function closeWindow() {
		activeWindow = null;
	}
</script>

<svelte:head>
	<title>Farhan Aulianda</title>
	<meta
		name="description"
		content="Full Stack & Cloud Engineer | Security & AI/ML Enthusiast"
	/>
</svelte:head>

<!-- Desktop Shell -->
<div class="fixed inset-0 bg-zinc-950 overflow-hidden select-none">
	<!-- Wallpaper - Dark minimal gradient -->
	<div
		class="absolute inset-0 bg-[radial-gradient(ellipse_at_top,_var(--tw-gradient-stops))] from-zinc-800/30 via-zinc-950 to-zinc-950"
	></div>
	<div
		class="absolute inset-0 bg-[radial-gradient(ellipse_at_bottom_right,_var(--tw-gradient-stops))] from-zinc-800/15 via-transparent to-transparent"
	></div>

	<!-- Menu Bar -->
	<MenuBar />

	<!-- Desktop Icons — LEFT column, macOS style -->
	<div class="absolute top-12 left-5 flex flex-col gap-1 z-10 pt-2">
		<!-- About Me -->
		<DesktopIcon label="About Me" onclick={() => openWindow('about')}>
			<svg class="w-7 h-7 text-blue-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="M15.75 6a3.75 3.75 0 1 1-7.5 0 3.75 3.75 0 0 1 7.5 0ZM4.501 20.118a7.5 7.5 0 0 1 14.998 0A17.933 17.933 0 0 1 12 21.75c-2.676 0-5.216-.584-7.499-1.632Z" />
			</svg>
		</DesktopIcon>

		<!-- Projects -->
		<DesktopIcon label="Projects" onclick={() => openWindow('projects')}>
			<svg class="w-7 h-7 text-amber-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12.75V12A2.25 2.25 0 0 1 4.5 9.75h15A2.25 2.25 0 0 1 21.75 12v.75m-8.69-6.44-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44Z" />
			</svg>
		</DesktopIcon>

		<!-- Blog -->
		<DesktopIcon label="Blog" onclick={() => openWindow('blog')}>
			<svg class="w-7 h-7 text-emerald-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10" />
			</svg>
		</DesktopIcon>

		<!-- Resume.pdf -->
		<DesktopIcon label="Resume.pdf" onclick={() => openWindow('resume')}>
			<svg class="w-7 h-7 text-red-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Z" />
			</svg>
		</DesktopIcon>

		<!-- Contact -->
		<DesktopIcon label="Contact" onclick={() => openWindow('contact')}>
			<svg class="w-7 h-7 text-violet-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
			</svg>
		</DesktopIcon>
	</div>

	<!-- Dock -->
	<Dock />
</div>

<!-- Window Modals -->
{#if activeWindow === 'about'}
	<DesktopWindow title="About Me" close={closeWindow}>
		<AboutMeContent />
	</DesktopWindow>
{/if}

{#if activeWindow === 'projects'}
	<DesktopWindow title="Projects" close={closeWindow}>
		<ProjectsContent />
	</DesktopWindow>
{/if}

{#if activeWindow === 'blog'}
	<DesktopWindow title="Blog" close={closeWindow}>
		<BlogContent {posts} />
	</DesktopWindow>
{/if}

{#if activeWindow === 'contact'}
	<DesktopWindow title="Contact" close={closeWindow}>
		<ContactContent />
	</DesktopWindow>
{/if}

{#if activeWindow === 'resume'}
	<ResumeModal close={closeWindow} />
{/if}
