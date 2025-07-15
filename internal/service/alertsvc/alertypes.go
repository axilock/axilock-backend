package alertsvc

import (
	"strings"

	db "github.com/axilock/axilock-backend/internal/db/sqlc"
)

type CreateSecretAlertReq struct {
	Filename   string
	Repo       string
	Org        int64
	Lineno     int64
	Commitid   string
	Source     string
	Secrettype string
	Filepath   string
	IsVerified bool
}

type GraphDataProtected struct {
	Total   int64                   `json:"total"`
	Buckets []db.GetAlertBucketsRow `json:"buckets"`
	Trend   float64                 `json:"trend"`
}

const (
	ALERT_OPEN   = "open"
	ALERT_CLOSED = "closed"
)

const (
	SECRET_DEFAULT = "default"
	SECRET_CUSTOM  = "custom"
)

const (
	INTERNAL_SECRET_PREFIX = "CUSTOM_"
)

func getAlertType(t string) string {
	if strings.HasPrefix(t, INTERNAL_SECRET_PREFIX) {
		return SECRET_CUSTOM
	}
	return SECRET_DEFAULT
}

const (
	AlertSecret = "secret"
)

const (
	SeverityLow      = "low"
	SeverityMedium   = "medium"
	SeverityHigh     = "high"
	SeverityCritical = "critical"
)
