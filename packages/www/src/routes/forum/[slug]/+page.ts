import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
    // In a real app, you'd fetch from Go using params.slug
    return {
        forumName: "Hard Surface Modeling",
        slug: params.slug,
        description: "Booleans, sub-d, and CAD-like workflows for props and vehicles.",
        threads: [
            {
                id: 501,
                title: "Fixing shading artifacts on curved surfaces",
                author: "TopoWizard",
                replies: 24,
                views: 1200,
                lastPost: "2 hours ago",
                isPinned: true
            },
            {
                id: 502,
                title: "Showcase: Sci-fi Grenade (Sub-D)",
                author: "BevelKing",
                replies: 8,
                views: 450,
                lastPost: "5 hours ago",
                isPinned: false
            },
            {
                id: 503,
                title: "How to handle massive boolean operations without crashing?",
                author: "CutterOps",
                replies: 15,
                views: 890,
                lastPost: "Yesterday",
                isPinned: false
            }
        ]
    };
};