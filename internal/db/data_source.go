package db

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"real-time-ranking/internal/config"
	"real-time-ranking/internal/models"
)

var _dataSource *datasource

type (
	ServiceDataSource interface {
		User
	}
	User interface {
		CreateUser(ctx context.Context, username, email string) (string, error)
		GetUser(ctx context.Context, id string) (*models.User, error)
		UpdateUser(ctx context.Context, id string, username, email string) error
		DeleteUser(ctx context.Context, id string) error
	}
)

type datasource struct {
	db *bun.DB
}

func InitDataSource(ctx context.Context, cfg config.Database) {
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DSN))), pgdialect.New())
	if cfg.Debug {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatalln("failed to connect to database:", err)
	}

	log.Println("Connected to database")
	_dataSource = &datasource{db: db}
}

func DataSource() ServiceDataSource {
	return _dataSource
}

func Close() {
	err := _dataSource.db.Close()
	if err != nil {
		log.Fatalf("close database failed: %v", err)
	}
	_ = _dataSource.db.Close()
	log.Println("Closed database")
}
