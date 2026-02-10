import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
    return {
        slug: params.slug,
        title: "Hard Surface",
        threadId: params.thread, 
        threadData: {
            title: "Fixing shading artifacts on curved surfaces",
            posts: [
                {
                    id: 1,
                    author: "TopoWizard",
                    role: "Elite Modeler",
                    avatar: "TW",
                    content: "I'm having trouble with N-gons causing pinching on this cylinder.",
                    timestamp: "Feb 08, 2026"
                },
                {
                    id: 2,
                    author: "BevelKing",
                    role: "Community Mod",
                    avatar: "BK",
                    content: "Have you tried using weighted normals?",
                    timestamp: "Feb 08, 2026"
                }
            ]
        }
    };
};
