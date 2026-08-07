CREATE INDEX IF NOT EXISTS idx_links_redirect_lookup
ON links(short_code, active);

CREATE INDEX IF NOT EXISTS idx_links_expiration
ON links(expires_at);