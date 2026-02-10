export const load = () => {
    return {
        tree: {
            id: 'root',
            label: '3D Modeling Root',
            children: [
                {
                    id: 'modeling',
                    label: 'Modeling & Geometry',
                    children: [
                        { id: 'hard-surface', label: 'Hard Surface' },
                        { id: 'organic', label: 'Organic/Sculpting' },
                        { id: 'topology', label: 'Topology' }
                    ]
                },
                {
                    id: 'surfacing',
                    label: 'Surfacing & Env',
                    children: [
                        { id: 'texturing', label: 'Texturing & UVs' },
                        { id: 'env-art', label: 'Environment Art' }
                    ]
                }
            ]
        }
    };
};