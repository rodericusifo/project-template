// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"errors"

	"auth-service/internal/domain/model/database/sql"
)

func (r *RoleDatabaseSQLRepository) SaveRole(payload *sql.Role) error {
	role := new(sql.Role)

	q := r.Writer()
	if q == nil {
		return errors.New("database sql writer is not configured")
	}

	if payload != nil {
		role = payload
	}

	if err := q.Table(r.model.TableName()).Save(role).Error; err != nil {
		return err
	}

	return nil
}
