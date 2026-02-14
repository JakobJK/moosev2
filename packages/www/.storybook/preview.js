import '../src/app.css';

/** @type { import('@storybook/sveltekit').Preview } */
const preview = {
    parameters: {
        options: {
            storySort: {
                order: ['Elements', 'Components', '*']
            },
        },
        controls: {
            matchers: {
                color: /(background|color)$/i,
                date: /Date$/i,
            },
        },
    },
};

export default preview;
