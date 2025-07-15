-- name: CreateGithubInstallation :exec
INSERT INTO git_provider (
    org_name, org_id, type, token , vcs_org_id, install_id)
VALUES ( $1, $2, $3, $4 , $5, $6);

-- name: GetGithubInstallation :one
SELECT * FROM git_provider
WHERE org_name = COALESCE(sqlc.narg(org_name), org_name) AND
 org_id = COALESCE(sqlc.narg(org_id), org_id) AND 
 vcs_org_id = COALESCE(sqlc.narg(vcs_org_id), vcs_org_id)
 LIMIT 1;