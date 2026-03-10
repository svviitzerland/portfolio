import { error } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function GET({ fetch }) {
	const response = await fetch('https://cloud.umami.is/script.js');
	if (!response.ok) {
		throw error(response.status, 'Failed to fetch script');
	}
	const script = await response.text();

	return new Response(script, {
		headers: {
			'content-type': 'application/javascript',
			'cache-control': 'public, max-age=3600'
		}
	});
}
