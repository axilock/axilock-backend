-- name: GetRepoCountForOrg :one
select COUNT(*) FROM repos WHERE 
org = $1;



-- name: GetRepoByEntity :one
SELECT * FROM repos
WHERE repourl = COALESCE(sqlc.narg(repourl), repourl) AND
provider = COALESCE(sqlc.narg(provider), provider) AND 
org = COALESCE(sqlc.narg(org), org) AND
vcs_repo_id = COALESCE(sqlc.narg(vcs_repo_id), vcs_repo_id)
 LIMIT 1;


-- name: CreateRepoWithProvider :copyfrom
INSERT INTO repos ( 
    name, repourl, provider, org, vcs_repo_id)
VALUES ( $1, $2, $3, $4, $5);

