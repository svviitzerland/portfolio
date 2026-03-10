import { error } from '@sveltejs/kit';
import { fetchMarkdownPosts } from '$lib/utils/posts';

export const load = async ({ params }) => {
	// Vite requires a glob for dynamic imports to work reliably with variables
	const posts = import.meta.glob('/src/lib/content/posts/*.md');
	const match = posts[`/src/lib/content/posts/${params.slug}.md`];

	if (!match) {
		throw error(404, `Could not find ${params.slug}`);
	}

	const post = await match();

	// Get all posts to find previous and next
	const allPosts = await fetchMarkdownPosts();
	const currentIndex = allPosts.findIndex((p) => p.slug === params.slug);

	const prevPost = currentIndex < allPosts.length - 1 ? allPosts[currentIndex + 1] : null;
	const nextPost = currentIndex > 0 ? allPosts[currentIndex - 1] : null;

	return {
		content: post.default,
		meta: post.metadata,
		prevPost,
		nextPost
	};
};
