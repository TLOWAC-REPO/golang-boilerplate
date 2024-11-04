package db

import (
	"golang-boilerplate/ent"
	"log"
	"sync"
)

var (
	Client *ent.Client
	once   sync.Once
)

// GetClient는 데이터베이스 클라이언트를 반환합니다.
// 클라이언트가 아직 초기화되지 않았다면, 초기화합니다.
func GetClient() *ent.Client {
	once.Do(func() {
		var err error
		Client, err = ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres123 sslmode=disable")
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
	})

	return Client
}
