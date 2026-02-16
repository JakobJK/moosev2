CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    sort_order INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE forums (
    id SERIAL PRIMARY KEY,
    category_id INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    slug TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    topic_count INTEGER NOT NULL DEFAULT 0,
    post_count INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE threads (
    id SERIAL PRIMARY KEY,
    forum_id INTEGER NOT NULL REFERENCES forums(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id),
    title TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    replies_count INTEGER NOT NULL DEFAULT 0,
    last_post_id INTEGER, 
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    thread_id INTEGER NOT NULL REFERENCES threads(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Add the foreign key constraint after both tables exist
ALTER TABLE threads 
ADD CONSTRAINT fk_last_post 
FOREIGN KEY (last_post_id) REFERENCES posts(id) ON DELETE SET NULL;

-- Performance Indexes
CREATE INDEX idx_forums_category_id ON forums(category_id);
CREATE INDEX idx_threads_forum_id ON threads(forum_id);
CREATE INDEX idx_posts_thread_id ON posts(thread_id);
