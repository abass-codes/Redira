package cache

import (
	"context"
	"encoding/json"
	"time"

	db "github.com/abass-codes/redira/internal/database/db"
)

const LinkTTL = 24 * time.Hour

func (c *Cache) GetLink(ctx context.Context, shortCode string) (*db.Link, error) {
	value, err := c.Get(ctx, LinkKey(shortCode))
	if err != nil {
		return nil, err
	}

	var link db.Link

	if err := json.Unmarshal([]byte(value), &link); err != nil {
		return nil, err
	}

	return &link, nil
}

func (c *Cache) StoreLink(ctx context.Context, link *db.Link) error {
	value, err := json.Marshal(link)
	if err != nil {
		return err
	}

	return c.Set(
		ctx,
		LinkKey(link.ShortCode),
		string(value),
		LinkTTL,
	)
}

func (c *Cache) DeleteLink(ctx context.Context, shortCode string) error {
	return c.Delete(ctx, LinkKey(shortCode))
}
