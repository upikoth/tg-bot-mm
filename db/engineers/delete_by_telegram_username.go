package engineers

import (
	"context"
	"log"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func (e *Engineers) DeleteByTelegramUsername(
	ctx context.Context,
	telegramUsername string,
) error {
	log.Println("DeleteByTelegramUsername", telegramUsername)

	qRes, qErr := e.db.Query().Query(ctx,
		`declare $telegram_username as text;

			delete from engineers
			where telegram_username = $telegram_username`,
		query.WithParameters(
			ydb.ParamsBuilder().Param("$telegram_username").Text(telegramUsername).Build(),
		),
	)

	if qErr != nil {
		log.Println("DeleteByTelegramUsername error", qErr.Error())
		return qErr
	}

	defer func() { _ = qRes.Close(ctx) }()

	return nil
}
