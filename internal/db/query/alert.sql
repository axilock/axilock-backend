-- name: CreateSecretAlert :exec
INSERT INTO alerts (
    file_name, commit_id, source, status, file_path, alert_config_id, org_id)
VALUES ( $1, $2, $3, $4, $5, $6, $7 );


-- name: CreateAlertConfig :one
INSERT INTO alert_config (
    org, type, severity, matcher, regex, "desc", is_active, alert_type )
VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 ) RETURNING id;

-- name: GetAlertConfigByName :one
SELECT id FROM alert_config 
WHERE org = $1
AND type = $2
LIMIT 1;


-- name: GetSecretTypeCount :many
SELECT
    ac.type,
    COUNT(a.id) AS count
FROM
    alerts a
JOIN
    alert_config ac ON a.alert_config_id = ac.id
WHERE
    a.org_id = $1
GROUP BY
    ac.type
ORDER BY
    count DESC;


-- name: Top10RepoBySecretCount :many
SELECT
    r.repourl AS repo_name,
    COUNT(a.id) AS secret_count
FROM
    repos r
JOIN
    commit_webhooks cw ON r.id = cw.repo_id
JOIN
    alerts a ON cw.id = a.commit_id
JOIN
    alert_config ac ON a.alert_config_id = ac.id
WHERE
    ac.alert_type = 'secret'
    AND r.org = $1
GROUP BY
    r.repourl
ORDER BY
    secret_count DESC
LIMIT 10;


-- -- name: GetAlertForEntity :many
-- SELECT * FROM alerts 
-- WHERE org_id = COALESCE(sqlc.narg(org_id), org_id) AND
-- repo_id = COALESCE(sqlc.narg(repo_id), repo_id) AND
-- status =  COALESCE(sqlc.narg(status), status)
-- LIMIT sqlc.arg(count);

-- name: GetWeeklyStats :one
-- WITH org_metrics AS (
--     SELECT
--         COUNT(*) FILTER (WHERE created_at >= CURRENT_DATE - INTERVAL '7 days') AS recent_count,
--         COUNT(*) FILTER (WHERE source = 'axi-cli') AS protected_count,
--         COUNT(*) FILTER (WHERE source = 'axi-runner') AS bypassed_count
--     FROM
--         alerts
--     WHERE
--         org = $1
-- )
-- SELECT
--     COALESCE(m.recent_count, 0) AS alerts_last_7_days,
--     COALESCE(m.protected_count, 0) AS alerts_protected,
--     COALESCE(m.bypassed_count, 0) AS alerts_bypassed
-- FROM
--     organisations o
-- CROSS JOIN
--     org_metrics m
-- WHERE
--     o.id = $1;


-- name: Top10RepoBySecretCount :many
SELECT 
    r.id AS repo_id,
    r.name AS repo_name,
    COUNT(a.id) AS alert_count
FROM 
    repos r
JOIN 
    commits_cli cc ON r.id = cc.repo
JOIN 
    alerts a ON cc.id = a.commit_id
WHERE 
    r.org = $1
GROUP BY 
    r.id, r.name
ORDER BY 
    alert_count DESC
LIMIT 10;


-- name: GetTotalAlertCount :one
SELECT COUNT(*) AS count
FROM alerts
WHERE org_id = $1;

-- name: GetAlertBuckets :many
WITH all_severities AS (
    SELECT 'critical' AS severity
    UNION ALL
    SELECT 'high'
    UNION ALL
    SELECT 'medium'
    UNION ALL
    SELECT 'low'
),
historical_total AS (
  SELECT
    s.severity,
    (
      SELECT COUNT(*)
      FROM alerts a
      INNER JOIN alert_config ac ON a.alert_config_id = ac.id
      WHERE 
        ac.severity = s.severity
        AND a.org_id = $1
        AND a.created_at < (current_date - INTERVAL '1 MONTH')
    ) AS total_prior_alerts
  FROM all_severities s
),
all_buckets AS (
  SELECT 
    generated_time AS bucket_start,
    LEAST( -- for last bucket
      generated_time + INTERVAL '4 days' - INTERVAL '1 second',
      current_date + INTERVAL '1 days' - INTERVAL '1 second'
    ) AS bucket_end
  FROM generate_series(
    current_date - INTERVAL '1 MONTH',
    current_date,
    '4 days'::interval
  ) AS generated_time
),
recent_buckets_data AS (
  SELECT
    (date_bin(
      '4 days',
      a.created_at,
      current_date - INTERVAL '1 MONTH'
    ) + INTERVAL '4 days' - INTERVAL '1 second') AS bucket_end,
    ac.severity,
    COUNT(*) AS alert_count
  FROM alerts a
  INNER JOIN alert_config ac ON a.alert_config_id = ac.id
  WHERE
    a.org_id = $1
    AND a.created_at >= (current_date - INTERVAL '1 MONTH')
  GROUP BY bucket_end, ac.severity
),
recent_buckets AS (
  SELECT
    b.bucket_end,
    s.severity,
    COALESCE(r.alert_count, 0) AS alert_count
  FROM all_buckets b
  CROSS JOIN all_severities s
  LEFT JOIN recent_buckets_data r 
    ON b.bucket_end = r.bucket_end AND s.severity = r.severity
  WHERE b.bucket_end <= current_date + INTERVAL '1 day' - INTERVAL '1 second'  -- Ensure no future buckets
)
SELECT
  rb.bucket_end,
  rb.severity,
  rb.alert_count AS period_count,
  (
    SUM(rb.alert_count) OVER (
    PARTITION BY rb.severity
    ORDER BY rb.bucket_end
    ROWS BETWEEN UNBOUNDED PRECEDING AND CURRENT ROW
    ) + COALESCE(ht.total_prior_alerts, 0)
  )::bigint AS cumulative_count
FROM recent_buckets rb
LEFT JOIN historical_total ht USING (severity)
ORDER BY rb.bucket_end, rb.severity;


 -- name: GetCommitsHealth :one
 SELECT 
     COUNT(*) AS total_commits,
     COUNT(*) FILTER (WHERE scanned_by_cli IS NOT TRUE) AS not_scanned_count
 FROM commit_webhooks
 WHERE org = $1;
