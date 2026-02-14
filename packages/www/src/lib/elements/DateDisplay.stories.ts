import type { Meta, StoryObj } from '@storybook/svelte';
import DateDisplay from './DateDisplay.svelte';

const meta = {
    title: 'Elements/DateDisplay',
    component: DateDisplay,
    tags: ['autodocs'],
    argTypes: {
        date: {
            control: 'text',
            description: 'ISO string or Date object',
        },
        className: {
            control: 'text',
            description: 'CSS classes for the span element',
        },
    },
} satisfies Meta<DateDisplay>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
    args: {
        date: '2026-02-13',
    },
};

export const FromDateObject: Story = {
    args: {
        date: new Date(),
    },
};

export const InvalidDate: Story = {
    args: {
        date: 'not-a-date',
    },
};

export const CustomStyled: Story = {
    args: {
        date: '2026-12-25',
        className: 'text-blue-600 font-bold italic',
    },
};
