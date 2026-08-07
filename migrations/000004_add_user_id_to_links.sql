ALTER TABLE links
ADD COLUMN user_id UUID REFERENCES users(id);

CREATE INDEX links_user_id_idx
ON links(user_id);