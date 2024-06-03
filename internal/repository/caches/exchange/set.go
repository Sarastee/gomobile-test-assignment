package exchange

import "context"

func (e *ExchangeCacheRepo) SetCache(ctx context.Context, key string, content string) error {
	_, err := e.client.DB().DoContext(ctx, setCommand, key, content, exCommand, e.redisConfig.TTL.Seconds())
	if err != nil {
		return err
	}

	return nil
}
