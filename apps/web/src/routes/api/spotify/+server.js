import { env } from '$env/dynamic/public';

/**
 * Parse the __NEXT_DATA__ entity from a Spotify embed page.
 * @param {string} url
 */
async function fetchSpotifyEntity(url) {
	const res = await fetch(url, {
		headers: { 'User-Agent': 'Mozilla/5.0 (compatible; portfolio/1.0)' }
	});
	if (!res.ok) return null;

	const html = await res.text();
	const match = html.match(
		/<script\s+id="__NEXT_DATA__"\s+type="application\/json">(.*?)<\/script>/s
	);
	if (!match) return null;

	try {
		const data = JSON.parse(match[1]);
		return data?.props?.pageProps?.state?.data?.entity ?? null;
	} catch {
		return null;
	}
}

/**
 * Fetch playlist data + derive owner profile from the embed page.
 * @param {string} playlistId
 */
async function fetchPlaylist(playlistId) {
	const entity = await fetchSpotifyEntity(
		`https://open.spotify.com/embed/playlist/${playlistId}`
	);
	if (!entity) return null;

	return {
		id: entity.id,
		name: entity.name || entity.title,
		owner: entity.subtitle || '',
		coverArt:
			entity.coverArt?.sources?.[0]?.url ?? entity.visualIdentity?.image?.[0]?.url ?? '',
		trackCount: entity.trackList?.length ?? 0,
		tracks: (entity.trackList ?? []).map((t) => ({
			title: t.title,
			artist: t.subtitle,
			duration: t.duration,
			isExplicit: t.isExplicit
		}))
	};
}

/** @type {import('./$types').RequestHandler} */
export async function GET() {
	const userId = env.PUBLIC_SPOTIFY_USER_ID;
	const playlistCsv = env.PUBLIC_SPOTIFY_PLAYLIST_IDS || '';
	const playlistIds = playlistCsv.split(',').map((s) => s.trim()).filter(Boolean);

	if (!userId || playlistIds.length === 0) {
		return new Response(
			JSON.stringify({ error: 'Spotify env vars not configured' }),
			{ status: 500, headers: { 'Content-Type': 'application/json' } }
		);
	}

	try {
		const results = await Promise.all(playlistIds.map(fetchPlaylist));
		const playlists = results.filter(Boolean);

		// Derive profile name from the first playlist's owner field
		const ownerName = playlists[0]?.owner || '';

		// Profile image: Spotify user avatar URL pattern
		const profileImage = `https://i.scdn.co/image/ab6775700000ee85${userId.length > 10 ? '' : ''}`;

		return new Response(
			JSON.stringify({
				profile: {
					name: ownerName,
					userId
				},
				playlists
			}),
			{
				headers: {
					'Content-Type': 'application/json',
					'Cache-Control': 'public, max-age=3600, s-maxage=3600'
				}
			}
		);
	} catch {
		return new Response(
			JSON.stringify({ error: 'Failed to fetch Spotify data' }),
			{ status: 500, headers: { 'Content-Type': 'application/json' } }
		);
	}
}
