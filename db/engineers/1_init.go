package engineers

import (
	"github.com/ydb-platform/ydb-go-sdk/v3"
)

type Engineers struct {
	db *ydb.Driver
}

func New(
	db *ydb.Driver,
) *Engineers {
	return &Engineers{
		db: db,
	}
}
