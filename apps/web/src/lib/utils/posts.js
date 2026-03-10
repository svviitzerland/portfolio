export const fetchMarkdownPosts = async () => {
	const allPostFiles = import.meta.glob('/src/lib/content/posts/*.md');
	const iterablePostFiles = Object.entries(allPostFiles);

	const allPosts = await Promise.all(
		iterablePostFiles.map(async ([path, resolver]) => {
			try {
				const { metadata } = await resolver();
				const postPath = path.slice(11, -3);
				const slug = path.split('/').pop().slice(0, -3);

				return {
					meta: metadata,
					path: postPath,
					slug
				};
			} catch (error) {
				console.error('Error loading post:', path, error);
				return null;
			}
		})
	);

	const validPosts = allPosts.filter((post) => post && post.meta && post.meta.date);

	return validPosts.sort((a, b) => {
		return new Date(b.meta.date) - new Date(a.meta.date);
	});
};
