// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package borrowrecord

import (
	"errors"

	"book-service/internal/domain/model/database/sql"
)

func (r *BorrowRecordDatabaseSQLRepository) SaveBorrowRecord(payload *sql.BorrowRecord) error {
	borrowRecord := new(sql.BorrowRecord)

	q := r.Writer()
	if q == nil {
		return errors.New("database sql writer is not configured")
	}

	if payload != nil {
		borrowRecord = payload
	}

	if err := q.Table(r.model.TableName()).Save(borrowRecord).Error; err != nil {
		return err
	}

	return nil
}
