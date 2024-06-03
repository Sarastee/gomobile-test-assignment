package exchange

import (
	"context"
	"fmt"
)

func (s *Service) GetCache(ctx context.Context, val string, data string) (string, error) {
	key := fmt.Sprintf("%s-%s", val, data)

	return s.exchangeCacheRepo.GetCache(ctx, key)
}
