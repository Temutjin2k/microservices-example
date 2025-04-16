CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    email citext UNIQUE NOT NULL,
    avatar_link text CHECK (url ~ '^https?://')
    password_hash bytea NOT NULL,
    is_deleted BOOLEAN DEFAULT 'FALSE',
    version integer NOT NULL DEFAULT 1
);