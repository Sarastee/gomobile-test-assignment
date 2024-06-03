package exchange

import (
	"context"
	"fmt"
)

func (s *Service) SetCache(ctx context.Context, val string, data string, content string) error {
	key := fmt.Sprintf("%s-%s", val, data)

	return s.exchangeCacheRepo.SetCache(ctx, key, content)
}
