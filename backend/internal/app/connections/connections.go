package connections

import (
	"context"
	"fmt"
	"log"
	"time"
	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/w0ikid/world-map-tracker/internal/app/config"
)

type Connections struct {
    DB *pgxpool.Pool
}

// close закрывает все соединения
func (c *Connections) Close() {
    c.DB.Close()
}

// func NewConnections(cfg *config.Config) (*Connections, error) {
// 	log.Println("Connecting to database...")
// 	log.Println("DB connection string:", cfg.DB.GetDBConnString())
// 	conn, err := pgx.Connect(context.Background(), cfg.DB.GetDBConnString())
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to connect to database: %w", err)
// 	}
// 	return &Connections{DB: conn}, nil
// }

func NewConnections(cfg *config.Config) (*Connections, error) {
    log.Println("Connecting to database...")
    log.Println("DB connection string:", cfg.DB.GetDBConnString())

    config, err := pgxpool.ParseConfig(cfg.DB.GetDBConnString())
    if err != nil {
        return nil, fmt.Errorf("failed to parse database config: %w", err)
    }

    // Настройка пула (опционально, настройте под свои нужды)
    config.MaxConns = 10 // Максимальное количество соединений
    config.MinConns = 2  // Минимальное количество соединений
    config.MaxConnLifetime = 30 * time.Minute // Максимальное время жизни соединения
    config.MaxConnIdleTime = 5 * time.Minute  // Время простоя соединения

    // Создаем пул соединений
    pool, err := pgxpool.NewWithConfig(context.Background(), config)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    // Проверяем подключение
    if err := pool.Ping(context.Background()); err != nil {
        pool.Close()
        return nil, fmt.Errorf("failed to ping database: %w", err)
    }

    log.Println("Successfully connected to database")
    return &Connections{DB: pool}, nil
}