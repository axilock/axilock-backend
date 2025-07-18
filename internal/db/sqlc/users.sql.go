// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  email, org, hash_password
) VALUES (
  $1, $2, $3
)
RETURNING id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at
`

type CreateUserParams struct {
	Email        string `json:"email"`
	Org          int64  `json:"org"`
	HashPassword string `json:"hash_password"`
}

// CreateUser
//
//	INSERT INTO users (
//	  email, org, hash_password
//	) VALUES (
//	  $1, $2, $3
//	)
//	RETURNING id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Org, arg.HashPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Email,
		&i.HashPassword,
		&i.Provider,
		&i.Org,
		&i.Role,
		&i.GithubUserID,
		&i.GithubUserName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

type CreateUserForGithubParams struct {
	Org            int64       `json:"org"`
	GithubUserID   pgtype.Int8 `json:"github_user_id"`
	GithubUserName pgtype.Text `json:"github_user_name"`
	Email          string      `json:"email"`
	HashPassword   string      `json:"hash_password"`
	Provider       string      `json:"provider"`
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at FROM users
WHERE email = $1
LIMIT 1
`

// GetUserByEmail
//
//	SELECT id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at FROM users
//	WHERE email = $1
//	LIMIT 1
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Email,
		&i.HashPassword,
		&i.Provider,
		&i.Org,
		&i.Role,
		&i.GithubUserID,
		&i.GithubUserName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEntityId = `-- name: GetUserByEntityId :one
SELECT id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at FROM users
WHERE uuid = $1
LIMIT 1
`

// GetUserByEntityId
//
//	SELECT id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at FROM users
//	WHERE uuid = $1
//	LIMIT 1
func (q *Queries) GetUserByEntityId(ctx context.Context, uuid string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEntityId, uuid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Email,
		&i.HashPassword,
		&i.Provider,
		&i.Org,
		&i.Role,
		&i.GithubUserID,
		&i.GithubUserName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByGithubId = `-- name: GetUserByGithubId :one
SELECT id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at FROM users
WHERE github_user_id = $1
LIMIT 1
`

// GetUserByGithubId
//
//	SELECT id, uuid, email, hash_password, provider, org, role, github_user_id, github_user_name, created_at, updated_at FROM users
//	WHERE github_user_id = $1
//	LIMIT 1
func (q *Queries) GetUserByGithubId(ctx context.Context, githubUserID pgtype.Int8) (User, error) {
	row := q.db.QueryRow(ctx, getUserByGithubId, githubUserID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Email,
		&i.HashPassword,
		&i.Provider,
		&i.Org,
		&i.Role,
		&i.GithubUserID,
		&i.GithubUserName,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec

UPDATE users SET hash_password = $1 WHERE email = $2
`

type UpdateUserPasswordParams struct {
	HashPassword string `json:"hash_password"`
	Email        string `json:"email"`
}

// UpdateUserPassword
//
//	UPDATE users SET hash_password = $1 WHERE email = $2
func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.HashPassword, arg.Email)
	return err
}
