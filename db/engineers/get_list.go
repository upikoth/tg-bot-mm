package engineers

import (
	"app/models"
	"context"
	"io"

	"github.com/pkg/errors"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func (e *Engineers) GetList(
	ctx context.Context,
) (res []models.Engineer, err error) {
	var resEngineers []models.Engineer

	qRes, qErr := e.db.Query().QueryResultSet(ctx, `select telegram_username from engineers`)

	if qErr != nil {
		return resEngineers, qErr
	}

	defer func() { _ = qRes.Close(ctx) }()

	for {
		row, rErr := qRes.NextRow(ctx)

		if rErr != nil && !errors.Is(rErr, io.EOF) {
			return resEngineers, rErr
		}

		if row == nil || errors.Is(rErr, io.EOF) {
			break
		}

		var user models.Engineer
		sErr := row.ScanNamed(
			query.Named("telegram_username", &user.TelegramUsername),
		)

		if sErr != nil {
			return resEngineers, sErr
		}

		resEngineers = append(resEngineers, user)
	}

	return resEngineers, nil
}
