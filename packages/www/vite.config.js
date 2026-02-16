import { sveltekit } from '@sveltejs/kit/vite'
import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
    plugins: [
        tailwindcss(),
        sveltekit(),
    ],
    server: {
        // Allows Vite to accept connections from outside the container
        host: '0.0.0.0',
        port: 5173,
        strictPort: true,
        allowedHosts: [
            'moose.dev',
            'api.moose.dev',
            'storybook.moose.dev'
        ],
        // This ensures Hot Module Replacement (HMR) works over SSL via Nginx
        hmr: {
            clientPort: 443,
            protocol: 'wss'
        },
        // Fix for Windows/WSL2/Podman not detecting file changes
        watch: {
            usePolling: true,
            interval: 100
        }
    }
});
