package db

import (
	"fmt"
	"golang-boilerplate/ent"
	"golang-boilerplate/internal/config"
	"log"
	"sync"

	"entgo.io/ent/dialect"
)

var (
	Client *ent.Client
	once   sync.Once
)

// GetClient는 데이터베이스 클라이언트를 반환합니다.
// 클라이언트가 아직 초기화되지 않았다면, 초기화합니다.
func GetClient() *ent.Client {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.Env.DBUser,
		config.Env.DBPwd,
		config.Env.DBHost,
		config.Env.DBPort,
		config.Env.DBName,
	)

	once.Do(func() {
		var err error
		Client, err = ent.Open(dialect.Postgres, dsn)
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}
	})

	return Client
}
