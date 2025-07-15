package orgsvc

type CreateRegexReq struct {
	RegexStr  string
	RegexType string
	Desc      string
	Org       int64
	Name      string
}
