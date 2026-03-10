import { fetchMarkdownPosts } from '$lib/utils/posts';

export const load = async () => {
	const posts = await fetchMarkdownPosts();
	return {
		posts
	};
};
