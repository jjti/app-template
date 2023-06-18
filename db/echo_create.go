package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// Create adds an echo to DDB.
func (e echoImpl) Create(ctx context.Context, echo Echo) error {
	_, err := e.ddb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &e.table,
		Item:      echo.attributes(),
	})
	if err != nil {
		return fmt.Errorf("failed to insert echo item: %w", err)
	}

	return nil
}
