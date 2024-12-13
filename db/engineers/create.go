package engineers

import (
	"app/models"
	"context"
	"log"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func (e *Engineers) Create(
	ctx context.Context,
	engineer models.Engineer,
) error {
	log.Println("Create", engineer.TelegramUsername)

	qRes, qErr := e.db.Query().Query(ctx,
		`declare $telegram_username as text;

			insert into engineers
			(telegram_username)
			values ($telegram_username);`,
		query.WithParameters(
			ydb.ParamsBuilder().Param("$telegram_username").Text(engineer.TelegramUsername).Build(),
		),
	)

	if qErr != nil {
		log.Println("DeleteByTelegramUsername error", qErr.Error())
		return qErr
	}

	defer func() { _ = qRes.Close(ctx) }()

	return nil
}
