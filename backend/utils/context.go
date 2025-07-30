package utils

import "context"

type contextKey string

const UserClaimsKey = contextKey("user")

// Accessor to extract claims cleanly
func GetClaims(ctx context.Context) (*Claims, bool) {
	claims, ok := ctx.Value(UserClaimsKey).(*Claims)
	return claims, ok
}
