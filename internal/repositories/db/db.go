package db

import (
	"os"

	"github.com/pkg/errors"
	"github.com/upikoth/tg-bot-mm/internal/config"
	"github.com/upikoth/tg-bot-mm/internal/pkg/logger"
	"github.com/upikoth/tg-bot-mm/internal/repositories/db/dbmodels"
	"github.com/upikoth/tg-bot-mm/internal/repositories/db/users"
	ydb "github.com/ydb-platform/gorm-driver"
	environ "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"gorm.io/gorm"
)

type DB struct {
	db     *gorm.DB
	config *config.Config
	logger logger.Logger
	Users  *users.Users
}

func New(
	logger logger.Logger,
	cfg *config.Config,
) (*DB, error) {
	db := &gorm.DB{}

	return &DB{
		db:     db,
		config: cfg,
		logger: logger,
		Users:  users.New(db, logger),
	}, nil
}

func (y *DB) Connect() error {
	filePath := y.config.AuthFileDirName + "/" + y.config.AuthFileName
	err := os.Setenv("YDB_SERVICE_ACCOUNT_KEY_FILE_CREDENTIALS", filePath)

	if err != nil {
		return errors.WithStack(err)
	}

	db, err := gorm.Open(
		ydb.Open(y.config.Dsn, ydb.With(environ.WithEnvironCredentials())),
	)

	if err != nil {
		return errors.WithStack(err)
	}

	*y.db = *db

	return y.AutoMigrate()
}

func (y *DB) Disconnect() error {
	if y.db == nil {
		return nil
	}

	db, err := y.db.DB()

	if err != nil {
		return errors.WithStack(err)
	}

	err = db.Close()

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (y *DB) AutoMigrate() error {
	err := y.db.AutoMigrate(
		&dbmodels.User{},
	)

	return errors.WithStack(err)
}
