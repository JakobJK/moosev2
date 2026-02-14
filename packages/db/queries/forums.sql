-- name: ListForums :many
WITH latest_posts AS (
    SELECT DISTINCT ON (t.forum_id)
        p.id,
        p.created_at,
        t.forum_id,
        t.title AS thread_title,
        u.username,
        u.id AS user_id
    FROM posts p
    JOIN threads t ON p.thread_id = t.id
    JOIN users u ON p.user_id = u.id
    ORDER BY t.forum_id, p.created_at DESC
)
SELECT 
    c.id AS category_id,
    c.name AS category_name,
    c.sort_order AS category_order,
    f.id AS forum_id,
    f.slug AS forum_slug,
    f.title AS forum_title,
    f.description AS forum_description,
    f.topic_count,
    f.post_count,
    lp.id AS last_post_id,
    lp.thread_title AS last_post_title,
    lp.username AS last_post_username,
    lp.user_id AS last_post_user_id,
    lp.created_at AS last_post_created_at
FROM categories c
LEFT JOIN forums f ON c.id = f.category_id
LEFT JOIN latest_posts lp ON f.id = lp.forum_id
ORDER BY c.sort_order, f.id;
