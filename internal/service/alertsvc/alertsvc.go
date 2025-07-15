package alertsvc

import (
	"context"
	"fmt"

	"github.com/axilock/axilock-backend/internal/axierr"
	"github.com/axilock/axilock-backend/internal/constants"
	db "github.com/axilock/axilock-backend/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type AlertServiceInterface interface {
	CreateSecretAlert(context.Context, CreateSecretAlertReq) error
	// GetAlertData(context.Context, int64, int32, string) ([]db.Alert, error)
	// GetWeeklyStats(context.Context, int64) (db.GetWeeklyStatsRow, error)
	GetTop10Repo(context.Context, int64) ([]db.Top10RepoBySecretCountRow, error)
	GetProtectedSecretsOverTime(context.Context, int64) (GraphDataProtected, error)
	GetAlertSecretTypeCount(context.Context, int64) ([]db.GetSecretTypeCountRow, error)
	GetCommitsHealth(context.Context, int64) (db.GetCommitsHealthRow, error)
}

type AlertService struct {
	store db.Store
}

func NewAlertService(store db.Store) AlertServiceInterface {
	return &AlertService{
		store: store,
	}
}

func (s *AlertService) CreateSecretAlert(ctx context.Context, req CreateSecretAlertReq) error {

	var alert_config_id int64
	alert_config_id, err := s.store.GetAlertConfigByName(ctx, db.GetAlertConfigByNameParams{
		Type: req.Secrettype,
		Org:  req.Org,
	})

	if err != nil {
		if axierr.Is(err, axierr.ErrRecordNotFound) {
			alert_config_id, err = s.store.CreateAlertConfig(ctx, db.CreateAlertConfigParams{
				Type:      req.Secrettype,
				Org:       req.Org,
				Severity:  "high",
				Matcher:   SECRET_DEFAULT,
				Regex:     pgtype.Text{String: "", Valid: true},
				Desc:      pgtype.Text{String: req.Secrettype, Valid: true},
				IsActive:  pgtype.Bool{Bool: true, Valid: true},
				AlertType: constants.ALERT_TYPE_SECRET,
			})
			if err != nil {
				return fmt.Errorf("cannot create alert config")
			}
		} else {
			return fmt.Errorf("cannot get alert config")
		}
	}

	commitID, err := s.store.GetCommitByCommitID(ctx, req.Commitid)
	if err != nil {
		return fmt.Errorf("cannot get commit id")
	}

	args := db.CreateSecretAlertParams{
		FileName:      req.Filename,
		CommitID:      commitID,
		Source:        constants.SOURCE_AXI_CLI,
		Status:        ALERT_OPEN,
		FilePath:      req.Filepath,
		AlertConfigID: alert_config_id,
		OrgID:         req.Org,
	}

	return s.store.CreateSecretAlert(ctx, args)
}

// func (s *AlertService) GetAlertData(ctx context.Context, org int64, count int32, state string) ([]db.Alert, error) {
// 	return s.store.GetAlertForEntity(ctx, db.GetAlertForEntityParams{
// 		Org: pgtype.Int8{
// 			Int64: org,
// 			Valid: true,
// 		},
// 		Status: pgtype.Text{
// 			String: state,
// 			Valid:  true,
// 		},
// 		Count: count,
// 	})
// }

// func (s *AlertService) GetWeeklyStats(ctx context.Context, org int64) (db.GetWeeklyStatsRow, error) {
// 	return s.store.GetWeeklyStats(ctx, org)
// }

func (s *AlertService) GetTop10Repo(ctx context.Context, org int64) ([]db.Top10RepoBySecretCountRow, error) {
	return s.store.Top10RepoBySecretCount(ctx, org)
}

func (s *AlertService) GetProtectedSecretsOverTime(
	ctx context.Context,
	org int64,
) (GraphDataProtected, error) {
	total, err := s.store.GetTotalAlertCount(ctx, org)
	if err != nil {
		return GraphDataProtected{}, fmt.Errorf("cannot get total stats")
	}

	buckets, err := s.store.GetAlertBuckets(ctx, org)
	if err != nil {
		return GraphDataProtected{}, fmt.Errorf("cannot get alert buckets")
	}

	firstBucketDate := buckets[0].BucketEnd
	lastBucketDate := buckets[len(buckets)-1].BucketEnd

	allSeverityAlertsInFirstBucket := int64(0)
	allSeverityAlertsInLastBucket := int64(0)

	for i := range buckets {
		if buckets[i].BucketEnd == firstBucketDate {
			allSeverityAlertsInFirstBucket += buckets[i].CumulativeCount
		} else {
			break // buckets are ordered by BucketEnd
		}
	}

	for i := len(buckets) - 1; i >= 0; i-- {
		if buckets[i].BucketEnd == lastBucketDate {
			allSeverityAlertsInLastBucket += buckets[i].CumulativeCount
		} else {
			break // buckets are ordered by BucketEnd
		}
	}

	trend := float64(0)
	if allSeverityAlertsInFirstBucket != allSeverityAlertsInLastBucket {
		trend = 100.0 * float64(allSeverityAlertsInLastBucket-allSeverityAlertsInFirstBucket) /
			float64(allSeverityAlertsInLastBucket)
	}

	return GraphDataProtected{
		Total:   total,
		Buckets: buckets,
		Trend:   trend,
	}, nil
}

func (s *AlertService) GetAlertSecretTypeCount(ctx context.Context, org int64) ([]db.GetSecretTypeCountRow, error) {
	return s.store.GetSecretTypeCount(ctx, org)
}

func (s *AlertService) GetCommitsHealth(ctx context.Context, org int64) (db.GetCommitsHealthRow, error) {
	return s.store.GetCommitsHealth(ctx, org)
}
