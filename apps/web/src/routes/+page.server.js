import { fetchMarkdownPosts } from '$lib/utils/posts';

export async function load({ request, platform }) {
	// Get geo data from Cloudflare headers
	let country = null;
	let city = null;
	let region = null;

	// Try Cloudflare headers first
	if (platform?.cf) {
		country = platform.cf.country;
		city = platform.cf.city;
		region = platform.cf.region;
	}

	// Fallback to CF headers
	if (!country) {
		country = request.headers.get('cf-ipcountry');
	}

	// Country name mapping for common codes
	const countryNames = {
		ID: 'Indonesia',
		US: 'United States',
		GB: 'United Kingdom',
		JP: 'Japan',
		SG: 'Singapore',
		MY: 'Malaysia',
		AU: 'Australia',
		DE: 'Germany',
		FR: 'France',
		NL: 'Netherlands',
		IN: 'India',
		KR: 'South Korea',
		CN: 'China',
		TH: 'Thailand',
		VN: 'Vietnam',
		PH: 'Philippines',
		CA: 'Canada',
		BR: 'Brazil',
		AE: 'UAE',
		SA: 'Saudi Arabia'
	};

	const countryName = country ? countryNames[country] || country : null;

	// Fetch blog posts for the desktop Blog modal
	const posts = await fetchMarkdownPosts();

	return {
		geo: {
			country,
			countryName,
			city,
			region
		},
		posts
	};
}
