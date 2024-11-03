package db

import (
	"context"
	"golang-boilerplate/ent"
	"golang-boilerplate/ent/migrate"
	"log"
	"sync"
)

var (
	client *ent.Client
	once   sync.Once
)

// GetClient는 데이터베이스 클라이언트를 반환합니다.
// 클라이언트가 아직 초기화되지 않았다면, 초기화합니다.
func GetClient() *ent.Client {
	once.Do(func() {
		var err error
		client, err = ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=golang-boilerplate-db password=postgres123 sslmode=disable")
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	})

	return client
}
