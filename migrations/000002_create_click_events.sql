CREATE TABLE click_events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    link_id UUID NOT NULL REFERENCES links(id) ON DELETE CASCADE,

    clicked_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    ip_address TEXT,
    user_agent TEXT,
    referer TEXT
);

CREATE INDEX idx_click_events_link_id
ON click_events(link_id);

CREATE INDEX idx_click_events_clicked_at
ON click_events(clicked_at);