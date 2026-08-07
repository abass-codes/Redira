CREATE INDEX IF NOT EXISTS idx_click_events_link_id_clicked_at
ON click_events(link_id, clicked_at);

CREATE INDEX IF NOT EXISTS idx_click_events_clicked_at
ON click_events(clicked_at);