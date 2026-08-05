CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS links (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    original_url TEXT NOT NULL,

    short_code VARCHAR(10) NOT NULL UNIQUE,

    title TEXT,

    click_count BIGINT NOT NULL DEFAULT 0,

    expires_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_links_short_code
ON links(short_code);

CREATE INDEX idx_links_created_at
ON links(created_at);