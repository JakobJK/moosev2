import type { PageLoad } from './$types';

export const load: PageLoad = ({ params }) => {
    return {
        slug: params.slug,
        title: "Hard Surface"
    };
};
