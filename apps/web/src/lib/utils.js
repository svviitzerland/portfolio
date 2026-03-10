import { clsx } from 'clsx';
import { twMerge } from 'tailwind-merge';

/**
 * Merges class names with tailwind-merge
 * @param {...(import('clsx').ClassValue)} inputs
 * @returns {string}
 */
export function cn(...inputs) {
	return twMerge(clsx(inputs));
}
