<script>
	import MenuBar from '$lib/components/desktop/MenuBar.svelte';
	import DesktopIcon from '$lib/components/desktop/DesktopIcon.svelte';
	import DesktopWindow from '$lib/components/desktop/DesktopWindow.svelte';
	import Dock from '$lib/components/desktop/Dock.svelte';
	import ProjectsContent from '$lib/components/desktop/ProjectsContent.svelte';
	import BlogContent from '$lib/components/desktop/BlogContent.svelte';
	import SpotifyContent from '$lib/components/desktop/SpotifyContent.svelte';
	import ResumeModal from '$lib/components/ResumeModal.svelte';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	let { data } = $props();

	/** @type {null | 'projects' | 'blog' | 'dossier' | 'spotify'} */
	let activeWindow = $state(null);
	let brightness = $state(100);

	const posts = $derived(data?.posts || []);

	// Desktop icon positions with localStorage persistence
	const POSITIONS_KEY = 'desktop-icon-positions';
	const defaultPositions = {
		dossier: { x: 20, y: 60 },
		projects: { x: 20, y: 160 },
		blog: { x: 20, y: 260 },
		spotify: { x: 20, y: 360 }
	};

	let iconPositions = $state(structuredClone(defaultPositions));

	function savePositions() {
		try {
			localStorage.setItem(POSITIONS_KEY, JSON.stringify(iconPositions));
		} catch (e) {
			// localStorage may be unavailable (private browsing, etc.)
		}
	}

	function moveIcon(name, x, y) {
		iconPositions[name] = { x, y };
		savePositions();
	}

	// Auto-open folder from ?open= query param + load saved icon positions
	onMount(() => {
		const openParam = $page.url.searchParams.get('open');
		if (openParam && ['projects', 'blog', 'dossier', 'spotify'].includes(openParam)) {
			activeWindow = openParam;
			// Clean URL without reloading
			goto('/', { replaceState: true, noScroll: true });
		}

		// Load saved icon positions from localStorage
		try {
			const saved = localStorage.getItem(POSITIONS_KEY);
			if (saved) {
				const parsed = JSON.parse(saved);
				// Merge with defaults to handle missing keys
				iconPositions = { ...defaultPositions, ...parsed };
			}
		} catch (e) {
			// Ignore parse errors, use defaults
		}
	});

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
	<!-- Wallpaper - Dark with wavy lines -->
	<div
		class="absolute inset-0 bg-[radial-gradient(ellipse_at_top,_var(--tw-gradient-stops))] from-zinc-800/30 via-zinc-950 to-zinc-950"
	></div>
	<div
		class="absolute inset-0 bg-[radial-gradient(ellipse_at_bottom_right,_var(--tw-gradient-stops))] from-zinc-800/15 via-transparent to-transparent"
	></div>
	<!-- Wavy horizontal lines pattern -->
	<svg class="absolute inset-0 w-full h-full opacity-[0.04] pointer-events-none" xmlns="http://www.w3.org/2000/svg">
		<defs>
			<pattern id="wavy-lines" x="0" y="0" width="200" height="40" patternUnits="userSpaceOnUse">
				<path d="M0 20 Q25 10, 50 20 T100 20 T150 20 T200 20" fill="none" stroke="#a1a1aa" stroke-width="0.8"/>
			</pattern>
		</defs>
		<rect width="100%" height="100%" fill="url(#wavy-lines)" />
	</svg>

	<!-- Menu Bar -->
	<MenuBar bind:brightness />

	<!-- Desktop Icons — freely draggable, macOS style -->
	<div class="absolute inset-0 z-10">
		<!-- Dossier (Resume) -->
		<DesktopIcon
			label="Farhan"
			onclick={() => openWindow('dossier')}
			x={iconPositions.dossier.x}
			y={iconPositions.dossier.y}
			onmove={(nx, ny) => moveIcon('dossier', nx, ny)}
		>
			<svg class="w-7 h-7 text-red-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m2.25 0H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Z" />
			</svg>
		</DesktopIcon>

		<!-- Projects -->
		<DesktopIcon
			label="Projects"
			onclick={() => openWindow('projects')}
			x={iconPositions.projects.x}
			y={iconPositions.projects.y}
			onmove={(nx, ny) => moveIcon('projects', nx, ny)}
		>
			<svg class="w-7 h-7 text-amber-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="M2.25 12.75V12A2.25 2.25 0 0 1 4.5 9.75h15A2.25 2.25 0 0 1 21.75 12v.75m-8.69-6.44-2.12-2.12a1.5 1.5 0 0 0-1.061-.44H4.5A2.25 2.25 0 0 0 2.25 6v12a2.25 2.25 0 0 0 2.25 2.25h15A2.25 2.25 0 0 0 21.75 18V9a2.25 2.25 0 0 0-2.25-2.25h-5.379a1.5 1.5 0 0 1-1.06-.44Z" />
			</svg>
		</DesktopIcon>

		<!-- Blog -->
		<DesktopIcon
			label="Blog"
			onclick={() => openWindow('blog')}
			x={iconPositions.blog.x}
			y={iconPositions.blog.y}
			onmove={(nx, ny) => moveIcon('blog', nx, ny)}
		>
			<svg class="w-7 h-7 text-emerald-400" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
				<path stroke-linecap="round" stroke-linejoin="round" d="m16.862 4.487 1.687-1.688a1.875 1.875 0 1 1 2.652 2.652L10.582 16.07a4.5 4.5 0 0 1-1.897 1.13L6 18l.8-2.685a4.5 4.5 0 0 1 1.13-1.897l8.932-8.931Zm0 0L19.5 7.125M18 14v4.75A2.25 2.25 0 0 1 15.75 21H5.25A2.25 2.25 0 0 1 3 18.75V8.25A2.25 2.25 0 0 1 5.25 6H10" />
			</svg>
		</DesktopIcon>

		<!-- Spotify -->
		<DesktopIcon
			label="Spotify"
			onclick={() => openWindow('spotify')}
			x={iconPositions.spotify.x}
			y={iconPositions.spotify.y}
			onmove={(nx, ny) => moveIcon('spotify', nx, ny)}
		>
			<svg class="w-7 h-7 text-[#1DB954]" viewBox="0 0 24 24" fill="currentColor">
				<path d="M12 0C5.4 0 0 5.4 0 12s5.4 12 12 12 12-5.4 12-12S18.66 0 12 0zm5.521 17.34c-.24.359-.66.48-1.021.24-2.82-1.74-6.36-2.101-10.561-1.141-.418.122-.779-.179-.899-.539-.12-.421.18-.78.54-.9 4.56-1.021 8.52-.6 11.64 1.32.42.18.479.659.301 1.02zm1.44-3.3c-.301.42-.841.6-1.262.3-3.239-1.98-8.159-2.58-11.939-1.38-.479.12-1.02-.12-1.14-.6-.12-.48.12-1.021.6-1.141C9.6 9.9 15 10.561 18.72 12.84c.361.181.54.78.241 1.2zm.12-3.36C15.24 8.4 8.82 8.16 5.16 9.301c-.6.179-1.2-.181-1.38-.721-.18-.601.18-1.2.72-1.381 4.26-1.26 11.28-1.02 15.721 1.621.539.3.719 1.02.419 1.56-.299.421-1.02.599-1.559.3z"/>
			</svg>
		</DesktopIcon>
	</div>

	<!-- Dock -->
	<Dock onopen={openWindow} />

	<!-- Brightness overlay -->
	{#if brightness < 100}
		<div
			class="absolute inset-0 bg-black pointer-events-none z-[45]"
			style="opacity: {(100 - brightness) / 100}"
		></div>
	{/if}
</div>

<!-- Window Modals -->
{#if activeWindow === 'projects'}
	<DesktopWindow title="Projects" close={closeWindow}>
		{#snippet children(search)}
			<ProjectsContent {search} />
		{/snippet}
	</DesktopWindow>
{/if}

{#if activeWindow === 'blog'}
	<DesktopWindow title="Blog" close={closeWindow}>
		{#snippet children(search)}
			<BlogContent {posts} {search} />
		{/snippet}
	</DesktopWindow>
{/if}

{#if activeWindow === 'dossier'}
	<ResumeModal close={closeWindow} />
{/if}

{#if activeWindow === 'spotify'}
	<DesktopWindow title="Spotify" close={closeWindow}>
		{#snippet children(search)}
			<SpotifyContent {search} />
		{/snippet}
	</DesktopWindow>
{/if}
