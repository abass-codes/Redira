-- Redira Database Indexes
-- Migration: 002


CREATE INDEX IF NOT EXISTS idx_links_short_code

ON links(short_code);



CREATE INDEX IF NOT EXISTS idx_links_user_id

ON links(user_id);



CREATE INDEX IF NOT EXISTS idx_links_created_at

ON links(created_at);



CREATE INDEX IF NOT EXISTS idx_analytics_link_id

ON analytics_events(link_id);



CREATE INDEX IF NOT EXISTS idx_analytics_created_at

ON analytics_events(created_at);