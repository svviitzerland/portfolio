<script>
	import { onMount } from 'svelte';
	import { env } from '$env/dynamic/public';

	let { search = '' } = $props();

	const SPOTIFY_USER_ID = env.PUBLIC_SPOTIFY_USER_ID;
	const SPOTIFY_USER_IMAGE = env.PUBLIC_SPOTIFY_USER_IMAGE;

	/** @type {{ name: string, userId: string } | null} */
	let profile = $state(null);

	/** @type {Array<{ id: string, name: string, coverArt: string, trackCount: number, tracks: Array<{ title: string, artist: string, duration: number, isExplicit: boolean }> }>} */
	let playlists = $state([]);

	let activePlaylistId = $state('');
	let loading = $state(true);
	let error = $state(false);

	const activePlaylist = $derived(playlists.find((p) => p.id === activePlaylistId) || playlists[0]);

	const filteredTracks = $derived(() => {
		if (!activePlaylist?.tracks) return [];
		if (!search) return activePlaylist.tracks;
		const q = search.toLowerCase();
		return activePlaylist.tracks.filter(
			(t) => t.title.toLowerCase().includes(q) || t.artist.toLowerCase().includes(q)
		);
	});

	function formatDuration(ms) {
		const minutes = Math.floor(ms / 60000);
		const seconds = Math.floor((ms % 60000) / 1000);
		return `${minutes}:${seconds.toString().padStart(2, '0')}`;
	}

	function getTotalDuration(tracks) {
		const totalMs = tracks.reduce((sum, t) => sum + (t.duration || 0), 0);
		const minutes = Math.floor(totalMs / 60000);
		if (minutes >= 60) {
			const hours = Math.floor(minutes / 60);
			const mins = minutes % 60;
			return `${hours} hr ${mins} min`;
		}
		return `${minutes} min`;
	}

	onMount(async () => {
		try {
			const res = await fetch('/api/spotify');
			if (!res.ok) throw new Error('Failed to fetch');
			const data = await res.json();
			profile = data.profile;
			playlists = data.playlists;
			if (playlists.length > 0) {
				activePlaylistId = playlists[0].id;
			}
		} catch {
			error = true;
		} finally {
			loading = false;
		}
	});
</script>

{#if loading}
	<div class="flex items-center justify-center h-full">
		<div class="flex flex-col items-center gap-4">
			<div class="w-10 h-10 border-2 border-[#1DB954] border-t-transparent rounded-full animate-spin"></div>
			<span class="text-zinc-500 text-sm">Connecting to Spotify...</span>
		</div>
	</div>
{:else if error || !profile}
	<div class="flex items-center justify-center h-full">
		<div class="flex flex-col items-center gap-4 text-center px-8">
			<svg class="w-12 h-12 text-zinc-700" viewBox="0 0 24 24" fill="currentColor">
				<path d="M12 0C5.4 0 0 5.4 0 12s5.4 12 12 12 12-5.4 12-12S18.66 0 12 0zm5.521 17.34c-.24.359-.66.48-1.021.24-2.82-1.74-6.36-2.101-10.561-1.141-.418.122-.779-.179-.899-.539-.12-.421.18-.78.54-.9 4.56-1.021 8.52-.6 11.64 1.32.42.18.479.659.301 1.02zm1.44-3.3c-.301.42-.841.6-1.262.3-3.239-1.98-8.159-2.58-11.939-1.38-.479.12-1.02-.12-1.14-.6-.12-.48.12-1.021.6-1.141C9.6 9.9 15 10.561 18.72 12.84c.361.181.54.78.241 1.2zm.12-3.36C15.24 8.4 8.82 8.16 5.16 9.301c-.6.179-1.2-.181-1.38-.721-.18-.601.18-1.2.72-1.381 4.26-1.26 11.28-1.02 15.721 1.621.539.3.719 1.02.419 1.56-.299.421-1.02.599-1.559.3z"/>
			</svg>
			<p class="text-sm text-zinc-500">Could not load Spotify data.</p>
			{#if SPOTIFY_USER_ID}
				<a
					href="https://open.spotify.com/user/{SPOTIFY_USER_ID}"
					target="_blank"
					rel="noopener noreferrer"
					class="text-xs text-[#1DB954] hover:underline"
				>
					Open profile on Spotify
				</a>
			{/if}
		</div>
	</div>
{:else}
	<div class="flex flex-col h-full">
		<!-- Profile Header -->
		<div class="relative px-4 sm:px-8 pt-5 sm:pt-8 pb-4 sm:pb-6 bg-gradient-to-b from-[#1DB954]/15 via-[#1DB954]/5 to-transparent">
			<div class="flex items-center gap-4 sm:gap-5">
				<!-- Profile Picture -->
				{#if SPOTIFY_USER_IMAGE}
					<img
						src={SPOTIFY_USER_IMAGE}
						alt={profile.name}
						class="w-16 h-16 sm:w-20 sm:h-20 rounded-full object-cover shadow-xl ring-2 ring-[#1DB954]/30"
					/>
				{:else}
					<div class="w-16 h-16 sm:w-20 sm:h-20 rounded-full bg-zinc-800 flex items-center justify-center ring-2 ring-[#1DB954]/30 shadow-xl">
						<svg class="w-8 h-8 sm:w-10 sm:h-10 text-zinc-600" fill="currentColor" viewBox="0 0 24 24">
							<path d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
						</svg>
					</div>
				{/if}

				<!-- Profile Info -->
				<div class="flex-1 min-w-0">
					<p class="text-xs font-medium text-[#1DB954] uppercase tracking-widest mb-1">Profile</p>
					<h2 class="text-xl sm:text-2xl font-bold text-white truncate">{profile.name || 'Spotify User'}</h2>
					<div class="flex flex-wrap items-center gap-2 sm:gap-3 mt-2">
						<span class="text-xs text-zinc-500">{playlists.length} Public Playlists</span>
						<a
							href="https://open.spotify.com/user/{SPOTIFY_USER_ID}"
							target="_blank"
							rel="noopener noreferrer"
							class="inline-flex items-center gap-1.5 text-xs text-zinc-400 hover:text-[#1DB954] transition-colors"
						>
							<svg class="w-3.5 h-3.5" viewBox="0 0 24 24" fill="currentColor">
								<path d="M12 0C5.4 0 0 5.4 0 12s5.4 12 12 12 12-5.4 12-12S18.66 0 12 0zm5.521 17.34c-.24.359-.66.48-1.021.24-2.82-1.74-6.36-2.101-10.561-1.141-.418.122-.779-.179-.899-.539-.12-.421.18-.78.54-.9 4.56-1.021 8.52-.6 11.64 1.32.42.18.479.659.301 1.02zm1.44-3.3c-.301.42-.841.6-1.262.3-3.239-1.98-8.159-2.58-11.939-1.38-.479.12-1.02-.12-1.14-.6-.12-.48.12-1.021.6-1.141C9.6 9.9 15 10.561 18.72 12.84c.361.181.54.78.241 1.2zm.12-3.36C15.24 8.4 8.82 8.16 5.16 9.301c-.6.179-1.2-.181-1.38-.721-.18-.601.18-1.2.72-1.381 4.26-1.26 11.28-1.02 15.721 1.621.539.3.719 1.02.419 1.56-.299.421-1.02.599-1.559.3z"/>
							</svg>
							Open in Spotify
						</a>
					</div>
				</div>
			</div>
		</div>

		<!-- Main Content: Sidebar + Track list -->
		<div class="flex flex-col sm:flex-row flex-1 min-h-0">
			<!-- Left Sidebar: Playlists -->
			<div class="sm:w-56 shrink-0 border-b sm:border-b-0 sm:border-r border-zinc-800/50 flex flex-col">
				<div class="px-4 py-3">
					<h3 class="text-[11px] font-semibold text-zinc-500 uppercase tracking-wider">Playlists</h3>
				</div>

				<div class="overflow-y-auto max-h-40 sm:max-h-none sm:flex-1 spotify-scroll">
					{#each playlists as playlist (playlist.id)}
						<button
							class="w-full flex items-center gap-3 px-4 py-2.5 text-left transition-all duration-150 {activePlaylistId === playlist.id
								? 'bg-zinc-800/70 border-l-2 border-[#1DB954]'
								: 'border-l-2 border-transparent hover:bg-zinc-800/40'}"
							onclick={() => (activePlaylistId = playlist.id)}
						>
							<img
								src={playlist.coverArt}
								alt={playlist.name}
								class="w-10 h-10 rounded object-cover shadow-md shrink-0"
							/>
							<div class="min-w-0 flex-1">
								<p class="text-sm font-medium truncate {activePlaylistId === playlist.id ? 'text-[#1DB954]' : 'text-zinc-200'}">
									{playlist.name}
								</p>
								<p class="text-[11px] text-zinc-500 truncate">{playlist.trackCount} tracks</p>
							</div>
						</button>
					{/each}
				</div>
			</div>

			<!-- Right: Track List -->
			<div class="flex-1 flex flex-col min-w-0">
				{#if activePlaylist}
					<!-- Playlist Header -->
					<div class="flex items-center gap-3 sm:gap-4 px-4 sm:px-6 py-3 sm:py-4 border-b border-zinc-800/50">
						<img
							src={activePlaylist.coverArt}
							alt={activePlaylist.name}
							class="w-12 h-12 sm:w-16 sm:h-16 rounded-lg object-cover shadow-lg"
						/>
						<div class="flex-1 min-w-0">
							<h3 class="text-base sm:text-lg font-bold text-white truncate">{activePlaylist.name}</h3>
							<p class="text-xs text-zinc-500 mt-0.5">
								{activePlaylist.trackCount} tracks · {getTotalDuration(activePlaylist.tracks)}
							</p>
						</div>
						<a
							href="https://open.spotify.com/playlist/{activePlaylist.id}"
							target="_blank"
							rel="noopener noreferrer"
							class="shrink-0 flex items-center gap-2 px-3 sm:px-4 py-1.5 sm:py-2 bg-[#1DB954] hover:bg-[#1ed760] text-black text-xs sm:text-sm font-semibold rounded-full transition-all duration-200 hover:scale-105"
						>
							<svg class="w-4 h-4" viewBox="0 0 24 24" fill="currentColor">
								<path d="M8 5v14l11-7z"/>
							</svg>
							Play
						</a>
					</div>

					<!-- Track List -->
					<div class="flex-1 overflow-y-auto spotify-scroll">
						<!-- Column Headers -->
						<div class="flex items-center gap-3 sm:gap-4 px-4 sm:px-6 py-2 border-b border-zinc-800/30 text-[11px] font-medium text-zinc-600 uppercase tracking-wider sticky top-0 bg-zinc-900/95 backdrop-blur-sm z-10">
							<span class="w-6 sm:w-8 text-center shrink-0">#</span>
							<span class="flex-1">Title</span>
							<span class="w-12 text-right shrink-0">
								<svg class="w-3.5 h-3.5 inline" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24">
									<circle cx="12" cy="12" r="10"/>
									<path d="M12 6v6l4 2"/>
								</svg>
							</span>
						</div>

						{#each filteredTracks() as track, i (track.title + i)}
							<div class="group flex items-center gap-3 sm:gap-4 px-4 sm:px-6 py-2.5 hover:bg-zinc-800/40 transition-colors">
								<!-- Track Number -->
								<span class="w-6 sm:w-8 text-center text-sm text-zinc-600 group-hover:text-zinc-400 shrink-0 tabular-nums">
									{i + 1}
								</span>

								<!-- Track Info -->
								<div class="flex-1 min-w-0">
									<p class="text-sm text-zinc-200 group-hover:text-white truncate">
										{track.title}
										{#if track.isExplicit}
											<span class="inline-flex items-center justify-center w-4 h-4 text-[9px] font-bold bg-zinc-700 text-zinc-400 rounded ml-1.5 align-middle">E</span>
										{/if}
									</p>
									<p class="text-xs text-zinc-500 group-hover:text-zinc-400 truncate">{track.artist}</p>
								</div>

								<!-- Duration -->
								<span class="w-12 text-right text-xs text-zinc-600 group-hover:text-zinc-400 tabular-nums shrink-0">
									{formatDuration(track.duration)}
								</span>
							</div>
						{/each}

						{#if filteredTracks().length === 0 && search}
							<div class="flex flex-col items-center justify-center py-16 text-zinc-500">
								<svg class="w-8 h-8 mb-3 text-zinc-700" fill="none" stroke="currentColor" stroke-width="1.5" viewBox="0 0 24 24">
									<path stroke-linecap="round" stroke-linejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z"/>
								</svg>
								<p class="text-sm">No tracks matching "{search}"</p>
							</div>
						{/if}
					</div>
				{/if}
			</div>
		</div>
	</div>
{/if}

<style>
	.spotify-scroll::-webkit-scrollbar {
		width: 6px;
	}

	.spotify-scroll::-webkit-scrollbar-track {
		background: transparent;
	}

	.spotify-scroll::-webkit-scrollbar-thumb {
		background-color: #52525b;
		border-radius: 3px;
	}

	.spotify-scroll::-webkit-scrollbar-thumb:hover {
		background-color: #71717a;
	}
</style>
