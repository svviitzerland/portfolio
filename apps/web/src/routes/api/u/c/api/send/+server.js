import { error } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function POST({ request, fetch }) {
	const body = await request.json();
	const userAgent = request.headers.get('user-agent') || '';

	const response = await fetch('https://api-gateway.umami.dev/api/send', {
		method: 'POST',
		headers: {
			'content-type': 'application/json',
			'user-agent': userAgent
		},
		body: JSON.stringify(body)
	});

	if (!response.ok) {
		const text = await response.text();
		throw error(response.status, text);
	}

	return new Response(await response.text(), {
		status: response.status,
		headers: {
			'content-type': 'text/plain'
		}
	});
}
