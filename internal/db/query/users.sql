-- name: CreateUser :one
INSERT INTO users (
  email, org, hash_password
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserByGithubId :one
SELECT * FROM users
WHERE github_user_id = $1
LIMIT 1;

-- name: GetUserByEntityId :one
SELECT * FROM users
WHERE uuid = $1
LIMIT 1;


-- name: CreateUserForGithub :copyfrom
INSERT INTO users (
  org, github_user_id, github_user_name, email, hash_password, provider
) VALUES (
  $1, $2, $3, $4, $5, $6
);

-- name: UpdateUserPassword :exec

UPDATE users SET hash_password = $1 WHERE email = $2;
