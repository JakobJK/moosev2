import type { PageLoad } from './$types';
export const ssr = false;
export const load: PageLoad = async ({ fetch }) => {
    // 1. Fetch from your Go API (Localhost because this runs in the browser)
    const response = await fetch('https://api.moose.dev/forums/');

    if (!response.ok) {
        throw new Error(`Failed to fetch forums: ${response.statusText}`);
    }

    // 2. Parse the clean JSON from your Go backend
    const categories = await response.json();
    // 3. Return the data to your +page.svelte
    return {
        categories
    };
};
