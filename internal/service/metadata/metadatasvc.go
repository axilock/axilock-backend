package metadata

import (
	"context"
	"encoding/json"
	"fmt"

	db "github.com/axilock/axilock-backend/internal/db/sqlc"
)

type MetaDataServiceInterface interface {
	CreateMetadata(ctx context.Context, req CreatetMetaDataReq) error
}

type MetaDataService struct {
	store db.Store
}

func NewMetaDataService(store db.Store) MetaDataServiceInterface {
	return &MetaDataService{
		store: store,
	}
}

func (s *MetaDataService) CreateMetadata(ctx context.Context, req CreatetMetaDataReq) error {
	args := db.CreateMetaDataParams{
		UserID:    req.UserID,
		Org:       req.OrgID,
		DataType:  req.MetaType,
		DataValue: json.RawMessage(req.MetaData),
	}

	err := s.store.CreateMetaData(ctx, args)
	if err != nil {
		return fmt.Errorf("cannot insert metadata in table%w", err)
	}
	return nil
}
