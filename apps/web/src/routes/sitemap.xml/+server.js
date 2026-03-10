import { fetchMarkdownPosts } from '$lib/utils/posts.js';

export async function GET({ url }) {
	const baseUrl = url.origin;
	const posts = await fetchMarkdownPosts();

	// Static pages
	const staticPages = [
		{ url: '', priority: '1.0', changefreq: 'weekly' },
		{ url: '/blog', priority: '0.8', changefreq: 'daily' }
	];

	// Dynamic blog posts
	const blogPosts = posts.map((post) => ({
		url: `/blog/${post.slug}`,
		lastmod: post.meta.date,
		priority: '0.6',
		changefreq: 'monthly'
	}));

	const allPages = [...staticPages, ...blogPosts];

	const sitemap = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
${allPages
	.map(
		(page) => `  <url>
    <loc>${baseUrl}${page.url}</loc>${page.lastmod ? `\n    <lastmod>${page.lastmod}</lastmod>` : ''}
    <changefreq>${page.changefreq}</changefreq>
    <priority>${page.priority}</priority>
  </url>`
	)
	.join('\n')}
</urlset>`;

	return new Response(sitemap, {
		headers: {
			'Content-Type': 'application/xml',
			'Cache-Control': 'public, max-age=3600'
		}
	});
}
