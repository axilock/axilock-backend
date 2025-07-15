package metadata

const (
	InitMeta = "INIT_META"
	RepoMeta = "REPO_META"
)

type CreatetMetaDataReq struct {
	MetaData string
	OrgID    int64
	UserID   int64
	MetaType string
}
