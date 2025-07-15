package commitsvc

type CreateCommitGrpcReq struct {
	RepoURL    string
	CommitData []CreateCommmitSvcDBParamas
	Org        int64
	UserID     int64
}

type CreateCommmitSvcDBParamas struct {
	CommitID     string `json:"commit_id"`
	AuthorName   string `json:"author_name"`
	AuthorEmail  string `json:"author_email"`
	CommitTime   string `json:"commit_time"`
	PushTime     string `json:"push_time"`
	ScannedByCli bool   `json:"scanned_by_cli"`
}

type CreateCommitGithubReq struct {
	RepoID     int64
	CommitData []CreateCommmitSvcDBParamas
	Org        int64
}
