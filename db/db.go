package db

import (
	"app/config"
	"app/db/engineers"
	"context"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
)

type DB struct {
	driver    *ydb.Driver
	Engineers *engineers.Engineers
}

func New(
	ctx context.Context,
	cfg *config.Config,
) (*DB, error) {
	filePath := fmt.Sprintf("%s/%s", cfg.AuthFileDirName, cfg.AuthFileName)

	driver, err := ydb.Open(
		ctx,
		cfg.Dsn,
		yc.WithServiceAccountKeyFileCredentials(filePath),
	)

	if err != nil {
		return nil, err
	}

	return &DB{
		driver:    driver,
		Engineers: engineers.New(driver),
	}, nil
}

func (d *DB) Close(ctx context.Context) error {
	return d.driver.Close(ctx)
}
