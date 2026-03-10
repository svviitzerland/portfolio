import { redirect } from '@sveltejs/kit';

export const load = async () => {
	// Redirect /blog to homepage with blog folder open
	throw redirect(307, '/?open=blog');
};
