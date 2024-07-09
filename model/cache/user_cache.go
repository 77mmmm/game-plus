package cache

import (
	"context"
	"game-pro/lib/redis"
)

type RedisRepository struct {
	boardclient redis.BoardClient
}

func NewRedisRepository() *RedisRepository {
	return &RedisRepository{
		boardclient: redis.InitClient(),
	}
}

func (repo *RedisRepository) Zadd(ctx context.Context, key string, username int64, userscore int64) (int64, error) {
	reply, err := repo.boardclient.ZAdd(ctx, key, userscore, userscore)
	if err != nil || reply == -1 {
		return -1, err
	}
	return reply, nil
}
func (repo *RedisRepository) Zrank(ctx context.Context, key string, username int64) (int64, error) {
	reply, err := repo.boardclient.ZRank(ctx, key, username)
	if err != nil {
		return -1, err
	}
	return reply, nil
}

func (repo *RedisRepository) Zscore(ctx context.Context, key string, username int64) (int64, error) {
	reply, err := repo.boardclient.ZScore(ctx, key, username)
	if err != nil {
		return -1, err
	}
	return reply, nil
}
func (repo *RedisRepository) Zcount(ctx context.Context, key string) (int64, error) {
	reply, err := repo.boardclient.ZCount(ctx, key)
	if err != nil {
		return -1, err
	}
	return reply, nil
}

func (repo *RedisRepository) Zrange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	reply, err := repo.boardclient.ZRange(ctx, key, start, stop)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
