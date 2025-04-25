package connections

import (
	"fmt"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/w0ikid/world-map-tracker/internal/app/config"
)

type Connections struct {
	DB *pgx.Conn
}

// close закрывает все соединения
func (c *Connections) Close() {
	c.DB.Close(context.Background())
}

func NewConnections(cfg *config.Config) (*Connections, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DB.GetDBConnString())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &Connections{DB: conn}, nil
}