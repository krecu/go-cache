package cache

import (
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	Client *redis.Client
}

func New(host string, db int) (*Redis, error){

	client := redis.NewClient(&redis.Options{
		Addr:     	host,
		Password: 	"", // no password set
		DB:       	db,  // use default DB
		MaxRetries: 	2,
		PoolSize: 	100,
		ReadTimeout: 	5*time.Second,
		WriteTimeout: 	5*time.Second,
		PoolTimeout: 	5*time.Second,
		IdleTimeout: 	5*time.Minute,
	})

	_, err := client.Ping().Result(); if err != nil {
		return &Redis{}, err
	}

	return &Redis{
		Client: client,
	}, nil
}

func (r *Redis) Close() {
	r.Client.Close()
}

func (r *Redis) Del(key string) (error){
	return r.Client.Del(key).Err()
}


func (r *Redis) Get(key string) ([]byte, error){
	val, err := r.Client.Get(key).Result()
	return []byte(val), err
}

func (r *Redis) Set(key string, val string, ttl int) (error){
	return r.Client.Set(key, val, 300 * time.Second).Err()
}