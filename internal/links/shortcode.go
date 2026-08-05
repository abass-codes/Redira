package links

import (
	"context"

	"github.com/abass-codes/redira/internal/utils"
)

func GenerateUniqueCode(ctx context.Context) (string, error) {
	// Later we'll check PostgreSQL for collisions.
	// For now, generate a secure random code.
	return utils.GenerateShortCode(6)
}
