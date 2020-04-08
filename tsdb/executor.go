package tsdb

import "github.com/FernandoMorais/influxdb/models"

// Executor is an interface for a query executor.
type Executor interface {
	Execute() <-chan *models.Row
}
