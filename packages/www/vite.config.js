import { sveltekit } from '@sveltejs/kit/vite'
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'


export default defineConfig({
	plugins: [
        tailwindcss(),
        sveltekit(),
    ],
        server: {
		allowedHosts: [
			'moose.dev',
			'api.moose.dev'
		],
		// This ensures Hot Module Replacement (HMR) works over SSL
		hmr: {
			clientPort: 443
		}
	}
});
