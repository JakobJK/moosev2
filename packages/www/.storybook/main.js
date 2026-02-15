/** @type { import('@storybook/sveltekit').StorybookConfig } */
const config = {
    features: { onboarding: false },
    stories: ["../src/**/*.stories.@(js|ts|svelte)"],
    addons: ["@storybook/addon-svelte-csf"],
    framework: "@storybook/sveltekit",
    staticDirs: ['../static'],
};
export default config;
