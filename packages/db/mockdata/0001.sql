BEGIN;

-- 1. 20 Users
INSERT INTO users (id, username, email, password_hash) VALUES
(1, 'admin', 'admin@moose.dev', '$2a$10$dummyhash'),
(2, 'topoviz', 'topo@moose.dev', '$2a$10$dummyhash'),
(3, 'gopher_one', 'go1@moose.dev', '$2a$10$dummyhash'),
(4, 'svelte_master', 'sm@moose.dev', '$2a$10$dummyhash'),
(5, 'db_guru', 'db@moose.dev', '$2a$10$dummyhash'),
(6, 'rust_fan', 'rust@moose.dev', '$2a$10$dummyhash'),
(7, 'zig_user', 'zig@moose.dev', '$2a$10$dummyhash'),
(8, 'pixel_art', 'pa@moose.dev', '$2a$10$dummyhash'),
(9, 'low_poly', 'lp@moose.dev', '$2a$10$dummyhash'),
(10, 'shader_guy', 'sg@moose.dev', '$2a$10$dummyhash'),
(11, 'bevy_pro', 'bp@moose.dev', '$2a$10$dummyhash'),
(12, 'docker_dev', 'dd@moose.dev', '$2a$10$dummyhash'),
(13, 'linux_user', 'lu@moose.dev', '$2a$10$dummyhash'),
(14, 'vim_lord', 'vl@moose.dev', '$2a$10$dummyhash'),
(15, 'java_script', 'js@moose.dev', '$2a$10$dummyhash'),
(16, 'react_refugee', 'rr@moose.dev', '$2a$10$dummyhash'),
(17, 'htmx_fan', 'hf@moose.dev', '$2a$10$dummyhash'),
(18, 'frontend_gal', 'fg@moose.dev', '$2a$10$dummyhash'),
(19, 'backend_guy', 'bg@moose.dev', '$2a$10$dummyhash'),
(20, 'qa_bot', 'qa@moose.dev', '$2a$10$dummyhash')
ON CONFLICT (id) DO NOTHING;

-- 2. Categories
INSERT INTO categories (id, name, sort_order) VALUES
(1, 'General Discussion', 1),
(2, 'Development', 2),
(3, 'Showcase', 3)
ON CONFLICT (id) DO NOTHING;

-- 3. Forums
INSERT INTO forums (id, category_id, slug, title, description) VALUES
(1, 1, 'announcements', 'Announcements', 'Official news.'),
(2, 2, 'golang', 'Go / Backend', 'Backend discussion.'),
(3, 2, 'svelte', 'Svelte / Frontend', 'Frontend discussion.'),
(4, 3, 'projects', 'Project Showcase', 'Show us what you built.')
ON CONFLICT (id) DO NOTHING;

-- 4. Threads
INSERT INTO threads (id, forum_id, user_id, title, slug, is_pinned) VALUES
(1, 1, 1, 'Welcome to MooseV2', 'welcome-to-moosev2', TRUE),
(2, 2, 3, 'How to use pgx with Go?', 'go-pgx-usage', FALSE),
(3, 3, 4, 'SvelteKit 5 is out!', 'sveltekit-5-release', FALSE),
(4, 4, 2, 'Showcase: Procedural Terrain', 'procedural-terrain', FALSE)
ON CONFLICT (id) DO NOTHING;

-- 5. Posts
INSERT INTO posts (id, thread_id, user_id, content) VALUES
(1, 1, 1, 'Welcome to the forum!'),
(2, 2, 3, 'Does anyone have a good example for pgx connection pools?'),
(3, 3, 4, 'The new runes system is amazing.'),
(4, 4, 2, 'Check out these heightmaps I generated in Go.')
ON CONFLICT (id) DO NOTHING;

-- 6. Link last_post_id back to threads
UPDATE threads SET last_post_id = id;

-- 7. Sync Sequences
SELECT setval(pg_get_serial_sequence('users', 'id'), coalesce(max(id), 1)) FROM users;
SELECT setval(pg_get_serial_sequence('categories', 'id'), coalesce(max(id), 1)) FROM categories;
SELECT setval(pg_get_serial_sequence('forums', 'id'), coalesce(max(id), 1)) FROM forums;
SELECT setval(pg_get_serial_sequence('threads', 'id'), coalesce(max(id), 1)) FROM threads;
SELECT setval(pg_get_serial_sequence('posts', 'id'), coalesce(max(id), 1)) FROM posts;

COMMIT;
