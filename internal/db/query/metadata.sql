-- name: CreateMetaData :exec
INSERT INTO metadata (
  "user_id", "org", data_type, data_value
) VALUES (
  $1, $2, $3, $4
);