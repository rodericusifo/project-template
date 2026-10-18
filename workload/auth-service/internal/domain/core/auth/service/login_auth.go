// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"auth-service/internal/domain/core/auth/service/dto/input"
	"auth-service/internal/domain/core/auth/service/dto/output"
	"auth-service/internal/pkg/types"
	"auth-service/internal/pkg/util/comparer"
	"auth-service/internal/pkg/util/generator"
)

func (s *AuthService) LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error) {
	userModelRes, err := s.UserDatabaseSQLRepository.FirstUser(&types.QuerySQL{
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, err
	}

	match := comparer.CompareHashAndPassword(userModelRes.Password, payload.Password)
	if !match {
		return nil, status.Error(codes.Unauthenticated, "email and password not match")
	}

	claims := &types.JwtCustomClaims{
		XID: userModelRes.XID,
	}

	token, err := generator.GenerateJWTTokenFromClaims(claims)
	if err != nil {
		return nil, err
	}

	loginAuthDto := &output.LoginAuthDTO{
		Token: token,
	}

	return loginAuthDto, nil
}
