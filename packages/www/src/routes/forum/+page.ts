import type { PageLoad } from './$types';
export const ssr = false;
export const load: PageLoad = async ({ fetch }) => {
    const response = await fetch('https://api.moose.dev/forums/');

    if (!response.ok) {
        throw new Error(`Failed to fetch forums: ${response.statusText}`);
    }

    const categories = await response.json();
    return {
        categories
    };
};
