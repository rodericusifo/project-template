// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"errors"

	"author-service/internal/domain/model/database/sql"
	"author-service/internal/pkg/constant"
	"author-service/internal/pkg/types"
	"author-service/internal/pkg/util/builder"
)

func (r *AuthorDatabaseSQLRepository) FirstAuthor(query *types.QuerySQL) (*sql.Author, error) {
	author := new(sql.Author)

	q := r.Reader()
	if q == nil {
		return nil, errors.New("database sql reader is not configured")
	}

	if query != nil {
		q = builder.BuildQuerySQL(r.model.TableName(), q, query, constant.DialectDatabaseSQL(q.Dialector.Name()))
	}

	if err := q.Table(r.model.TableName()).First(author).Error; err != nil {
		return nil, err
	}

	return author, nil
}
