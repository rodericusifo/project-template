// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"errors"

	"auth-service/internal/domain/model/database/sql"
)

func (r *UserDatabaseSQLRepository) SaveUser(payload *sql.User) error {
	user := new(sql.User)

	q := r.Writer()
	if q == nil {
		return errors.New("database sql writer is not configured")
	}

	if payload != nil {
		user = payload
	}

	if err := q.Table(r.model.TableName()).Save(user).Error; err != nil {
		return err
	}

	return nil
}
