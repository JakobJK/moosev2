DROP INDEX IF EXISTS idx_posts_thread_id;
DROP INDEX IF EXISTS idx_threads_forum_id;
DROP INDEX IF EXISTS idx_forums_category_id;

DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS threads;
DROP TABLE IF EXISTS forums;
DROP TABLE IF EXISTS categories;
