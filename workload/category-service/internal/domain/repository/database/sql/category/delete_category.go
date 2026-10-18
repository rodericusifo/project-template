// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"errors"

	"category-service/internal/domain/model/database/sql"
)

func (r *CategoryDatabaseSQLRepository) DeleteCategory(payload *sql.Category) error {
	category := new(sql.Category)

	q := r.Writer()
	if q == nil {
		return errors.New("database sql writer is not configured")
	}

	if payload != nil {
		category = payload
	}

	if err := q.Table(r.model.TableName()).Unscoped().Delete(category).Error; err != nil {
		return err
	}

	return nil
}
