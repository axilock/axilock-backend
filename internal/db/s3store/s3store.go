package s3store

type S3StoreInterface interface{}

type S3Store struct{}

func NewS3Clinet() S3StoreInterface {
	return &S3Store{}
}
