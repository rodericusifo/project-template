// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"errors"

	"auth-service/internal/domain/model/database/sql"
)

func (r *RolePermissionDatabaseSQLRepository) SaveRolePermission(payload *sql.RolePermission) error {
	rolePermission := new(sql.RolePermission)

	q := r.Writer()
	if q == nil {
		return errors.New("database sql writer is not configured")
	}

	if payload != nil {
		rolePermission = payload
	}

	if err := q.Table(r.model.TableName()).Save(rolePermission).Error; err != nil {
		return err
	}

	return nil
}
