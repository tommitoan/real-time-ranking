package db

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
	"real-time-ranking/internal/models"
)

var _dataSource *datasource

type (
	ServiceDataSource interface {
		User
		Video
	}

	User interface {
		CreateUser(ctx context.Context, username, email string) (string, error)
		GetUser(ctx context.Context, id string) (*models.User, error)
		UpdateUser(ctx context.Context, id string, username, email string) error
		DeleteUser(ctx context.Context, id string) error
	}

	Video interface {
		CreateVideo(ctx context.Context, v *models.Video) (string, error)
		GetVideo(ctx context.Context, id string) (*models.Video, error)
		UpdateVideo(ctx context.Context, req models.UpdateVideoRequest, id string) error
		UpdateVideoReactions(ctx context.Context, id string, r *models.UpdateVideoReactionsRequest) (*models.Video, error)
		DeleteVideo(ctx context.Context, id string) error
	}
)

type datasource struct {
	db *bun.DB
}

func InitDataSource(ctx context.Context, dsn string, debug bool) {
	db := bun.NewDB(sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn))), pgdialect.New())
	if debug {
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
