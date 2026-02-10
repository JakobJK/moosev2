import type { PageLoad } from './$types';

export const load: PageLoad = () => {
    return {
        categories: [
            {
                id: 1,
                name: "Modeling & Geometry",
                order: 1,
                forums: [
                    {
                        id: 101,
                        slug: "hard-surface",
                        title: "Hard Surface Modeling",
                        description: "Booleans, sub-d, and CAD-like workflows for props and vehicles.",
                        topicCount: 3100,
                        postCount: 15400,
                        latestPost: {
                            id: 901,
                            title: "Fixing shading artifacts on curved surfaces",
                            username: "TopoWizard",
                            userId: 301,
                            createdAt: "2026-02-08T16:20:00Z"
                        }
                    },
                    {
                        id: 102,
                        slug: "organic-sculpting",
                        title: "Organic & Sculpting",
                        description: "Anatomy, creature sculpting, and high-to-low poly baking.",
                        topicCount: 2800,
                        postCount: 11000,
                        latestPost: {
                            id: 905,
                            title: "Brush settings for realistic skin pores",
                            username: "SculptLife",
                            userId: 205,
                            createdAt: "2026-02-08T15:45:00Z"
                        }
                    },
                    {
                        id: 103,
                        slug: "topology",
                        title: "Topology & Retopology",
                        description: "Optimization, quad-draw, and clean edge-flow discussion.",
                        topicCount: 1200,
                        postCount: 5200,
                        latestPost: {
                            id: 910,
                            title: "Retopo for game assets: How low is too low?",
                            username: "PolyCount",
                            userId: 102,
                            createdAt: "2026-02-08T10:10:00Z"
                        }
                    }
                ]
            },
            {
                id: 2,
                name: "Surfacing & Environment",
                order: 2,
                forums: [
                    {
                        id: 201,
                        slug: "texturing-uvs",
                        title: "Texturing & UVs",
                        description: "Unwrapping strategies, PBR texturing, and trim sheets.",
                        topicCount: 1900,
                        postCount: 8200,
                        latestPost: {
                            id: 1005,
                            title: "Smart materials vs manual painting",
                            username: "UV_Unwrapper",
                            userId: 55,
                            createdAt: "2026-02-08T17:01:00Z"
                        }
                    },
                    {
                        id: 202,
                        slug: "environment-art",
                        title: "Environment Art",
                        description: "Level assembly, modular kits, and lighting setups.",
                        topicCount: 2200,
                        postCount: 9500,
                        latestPost: {
                            id: 1020,
                            title: "Constructing a modular sci-fi corridor",
                            username: "Env_Artist",
                            userId: 12,
                            createdAt: "2026-02-07T22:30:00Z"
                        }
                    }
                ]
            },
            {
                id: 3,
                name: "Showcase & Feedback",
                order: 3,
                forums: [
                    {
                        id: 301,
                        slug: "wip-critique",
                        title: "WIP Critiques",
                        description: "Post your raw meshes and gray-box models for feedback.",
                        topicCount: 7500,
                        postCount: 62000,
                        latestPost: {
                            id: 1105,
                            title: "WIP: Mech Leg Hydraulics",
                            username: "MechHead",
                            userId: 9,
                            createdAt: "2026-02-08T11:00:00Z"
                        }
                    },
                    {
                        id: 302,
                        slug: "finished-works",
                        title: "Finished Works",
                        description: "Final renders and wireframe turnarounds.",
                        topicCount: 4300,
                        postCount: 21000,
                        latestPost: {
                            id: 1120,
                            title: "Vintage Camera - Final Render",
                            username: "CyclesLover",
                            userId: 44,
                            createdAt: "2026-02-08T08:15:00Z"
                        }
                    }
                ]
            }
        ]
    };
};
