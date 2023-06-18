package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/jjtimmons/sblast/pkg/metrics"
)

// Create adds an echo to DDB.
func (e echoImpl) Create(ctx context.Context, echo Echo) error {
	_, err := e.ddb.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &e.table,
		Item:      echo.attributes(),
	})
	if err != nil {
		metrics.PutItemCounter.With(prometheus.Labels{"result": "failure"}).Inc()
		return fmt.Errorf("failed to insert echo item: %w", err)
	}

	metrics.PutItemCounter.With(prometheus.Labels{"result": "success"}).Inc()
	return nil
}
