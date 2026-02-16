/** @type { import('@storybook/sveltekit').StorybookConfig } */
const config = {
  features: { onboarding: false },
  stories: ["../src/**/*.stories.@(js|ts|svelte)"],
  addons: ["@storybook/addon-svelte-csf"],
  framework: "@storybook/sveltekit",
  staticDirs: ['../static'],
  // Add the viteFinal configuration below
  async viteFinal(config) {
    const { mergeConfig } = await import('vite');
    return mergeConfig(config, {
      server: {
        // Required for Vite 6+ to allow the custom domain
        allowedHosts: ['storybook.moose.dev'],
        hmr: {
          host: 'storybook.moose.dev',
          protocol: 'wss',
        },
      },
    });
  },
};

export default config;
