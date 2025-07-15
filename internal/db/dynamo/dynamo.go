package dynamo

type NoSQLStoreInterface interface{}

type DynamoDBStore struct{}

func NewDynamoDBStore() NoSQLStoreInterface {
	return &DynamoDBStore{}
}
