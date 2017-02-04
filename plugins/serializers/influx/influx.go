package influx

import (
	"github.com/influxdata/telegraf-registry"
)

type InfluxSerializer struct {
}

func (s *InfluxSerializer) Serialize(m telegraf.Metric) ([]byte, error) {
	return m.Serialize(), nil
}
