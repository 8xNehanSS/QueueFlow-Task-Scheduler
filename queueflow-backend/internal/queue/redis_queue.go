package queue

import (
	"context"
	"encoding/json"
	"queueflow/internal/models"

	"github.com/redis/go-redis/v9"
)

const QueueKey = "jobs_queue"

type RedisQueue struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisQueue() *RedisQueue {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &RedisQueue{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (q *RedisQueue) Push(job models.Job) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return q.client.LPush(q.ctx, QueueKey, data).Err()
}

func (q *RedisQueue) Pop() (models.Job, error) {
	result, err := q.client.BRPop(q.ctx, 0, QueueKey).Result()
	if err != nil {
		return models.Job{}, err
	}

	var job models.Job
	err = json.Unmarshal([]byte(result[1]), &job)
	return job, err
}
