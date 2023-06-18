package db

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Echo is a single echo item in DDB.
type Echo struct {
	ID    string
	Value string
}

func (e Echo) attributes() map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		"id": &types.AttributeValueMemberS{
			Value: e.ID,
		},
		"value": &types.AttributeValueMemberS{
			Value: e.Value,
		},
	}
}
