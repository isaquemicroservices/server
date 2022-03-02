package redis

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Service struct {
	pool *redis.Pool
	conn redis.Conn
}

// NewConn get new connection with database redis
func NewConn() *Service {
	var redisPoll = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		MaxConnLifetime: time.Second * 10,
	}

	// Get connection with redis
	var conn = redisPoll.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		log.Fatalf("can't connect to the redis database, got error:\n%v", err)
	}

	return &Service{
		pool: redisPoll,
		conn: conn,
	}
}

// Publish publish key with value
func (s *Service) Publish(key string, value ...string) (err error) {
	conn := s.pool.Get()
	// defer conn.Close()

	if _, err = conn.Do("PUBLISH", key, value); err != nil {
		return err
	}

	return nil
}

// Subscribe ...
func (s *Service) Subscribe(key string, msg chan []byte) (err error) {
	pubSubConn := redis.PubSubConn{Conn: s.pool.Get()}
	// defer pubSubConn.Close()

	if err = pubSubConn.PSubscribe(key); err != nil {
		return err
	}

	go func() {
		for {
			switch v := pubSubConn.Receive().(type) {
			case redis.PMessage:
				msg <- v.Data
			}
		}
	}()

	return nil
}
