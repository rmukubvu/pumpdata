package cache

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

type Service struct {
	conn redis.Conn
	pool *redis.Pool
}

func New() *Service {
	var _pool *redis.Pool
	_pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}
	// Get a connection
	conn := _pool.Get()
	defer conn.Close()
	// Test the connection
	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalf("can't connect to the redis database, got error:\n%v", err)
	}

	return &Service{
		pool: _pool,
		conn: conn,
	}
}

func (s *Service) Set(key string, v string) error {
	conn := s.pool.Get()
	_, err := conn.Do("SET", key, v)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) SetWithExpiry(key string, v string, expirySeconds int) error {
	conn := s.pool.Get()
	_, err := conn.Do("SET", key, v, "EX", expirySeconds)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get(key string) (string, error) {
	conn := s.pool.Get()
	item, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return item, nil
}
