package metrics

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	// PutItemCounter counts up how many puts are made to DDB.
	PutItemCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "echo",
			Name:      "ddb_put_item_total",
			Help:      "Number of DDB put item calls.",
		},
		[]string{"result"},
	)
)

// Register all the metrics in Prometheus.
func Register() error {
	for _, m := range []prometheus.Collector{
		PutItemCounter,
	} {
		if err := prometheus.Register(m); err != nil {
			return fmt.Errorf("failed to register metric: %w", err)
		}
	}

	return nil
}
