import type { Meta, StoryObj } from '@storybook/svelte';
import VideoPlayer from './VideoPlayer.svelte';

const meta = {
  title: 'Elements/VideoPlayer',
  component: VideoPlayer,
} satisfies Meta<VideoPlayer>;

export default meta;
type Story = StoryObj<typeof meta>;

export const Default: Story = {
  args: {
    videoId: 'dQw4w9WgXcQ'
  }
};