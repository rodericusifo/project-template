// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"errors"

	"book-service/internal/domain/model/database/sql"
)

func (r *BookDatabaseSQLRepository) SaveBook(payload *sql.Book) error {
	book := new(sql.Book)

	q := r.Writer()
	if q == nil {
		return errors.New("database sql writer is not configured")
	}

	if payload != nil {
		book = payload
	}

	if err := q.Table(r.model.TableName()).Save(book).Error; err != nil {
		return err
	}

	return nil
}
