import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
    return {
        username: params.username,
        bio: "3D Environment Artist specializing in modular workflows and trim sheets.",
        avatarUrl: null, // Placeholder
        joinedDate: "2024-11-12",
        stats: {
            posts: 154,
            topics: 12,
            reputation: 850
        },
        recentActivity: [
            { id: 101, title: "Low Poly Sci-Fi Door", type: "WIP", date: "2 hours ago" },
            { id: 105, title: "Modular Building Kit", type: "Showcase", date: "3 days ago" }
        ]
    };
};