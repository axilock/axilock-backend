-- name: GetOrgByEntity :one
SELECT * FROM organisations
WHERE domain = COALESCE(sqlc.narg(domain), domain) AND
id = COALESCE(sqlc.narg(id), id) AND 
name = COALESCE(sqlc.narg(name), name) AND
github_org_id = COALESCE(sqlc.narg(github_org_id), github_org_id)
 LIMIT 1;


-- name: CreateOrganiation :one
INSERT INTO organisations (
  name, domain, provider, github_org_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;


-- name: GetOrgByID :one
SELECT * FROM organisations
WHERE id = COALESCE(sqlc.narg(id), id)
 LIMIT 1;


-- name: CreateRegexForOrg :exec
-- INSERT INTO regexs (
--   "org", regstring, description, version, name
-- ) VALUES (
--   $1, $2, $3, $4, $5
-- );

-- name: GetRegexesForOrg :many
-- SELECT * FROM regexs
-- WHERE org = $1;