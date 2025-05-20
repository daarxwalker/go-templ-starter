package example_repository

import (
	"context"
	
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	
	"common/pkg/service/database_service"
)

func GetExample(c context.Context, db database_service.Client) (int, error) {
	var result int
	sql, args, createSqlErr := squirrel.Select("1").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if createSqlErr != nil {
		return result, createSqlErr
	}
	if scanErr := pgxscan.Get(c, db, &result, sql, args...); scanErr != nil {
		return result, scanErr
	}
	return result, nil
}
