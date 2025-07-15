-- name: CreateCommmitsFromRPC :copyfrom
INSERT INTO commits_cli (
  commit_id, repo, author, commit_time, org, user_id,
  push_time, sessionid, source, user_repo_url
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
);


-- name: GetCommitByCommitID :one
SELECT id FROM commits_cli
WHERE commit_id = $1
LIMIT 1;



-- name: CreateVCSCommit :copyfrom
INSERT INTO commit_webhooks (
  commit_id, repo_id, author_name, author_email, commit_time, org, provider , scanned_by_cli,
  scanned_by_runner
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
);


-- name: LinkRepoInCommitCli :exec
UPDATE commits_cli
SET repo = $2
WHERE commit_id = $1 AND org = $3;


-- name: GetCommitsHealth :one
SELECT 
    COUNT(*) AS total_commits,
    COUNT(*) FILTER (WHERE scanned_by_cli IS NOT TRUE) AS not_scanned_count
FROM commit_webhooks
WHERE org = $1;


-- name: GetUniqueCommitUsernames :many
SELECT 
    author_name
FROM commit_webhooks
WHERE org = $1
GROUP BY author_name;


