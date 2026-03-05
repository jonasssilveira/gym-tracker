package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gym-tracker/app/series"
	config "gym-tracker/infra"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	client *redis.Client
}

func NewRedisConfig(cfg config.Redis) Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DbName,
	})
	return Redis{
		client,
	}
}

func (r Redis) Get(ctx context.Context, seriesID int) (series.Series, error) {
	result, err := r.client.GetDel(ctx, string(rune(seriesID))).Result()
	if errors.Is(err, redis.Nil) {
		return series.Series{}, fmt.Errorf(`series "%d" not found`, seriesID)
	} else if errors.Is(err, redis.Nil) {
		return series.Series{}, fmt.Errorf("failed to get user from redis: %w", err)
	}
	var s series.Series
	err = json.Unmarshal([]byte(result), &s)
	if err != nil {
		return series.Series{}, fmt.Errorf("failed to unmarshal result: %w", err)
	}
	return s, nil
}

func (r Redis) Set(ctx context.Context, series series.Series) (string, error) {
	return r.client.Set(ctx, string(rune(series.ID)), series, 0).Result()
}
