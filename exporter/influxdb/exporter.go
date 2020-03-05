// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package influxdb

import (
	"context"
	"errors"
	"log"

	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	influxDbClient "github.com/influxdata/influxdb-client-go"

	"github.com/open-telemetry/opentelemetry-collector/component"
	"github.com/open-telemetry/opentelemetry-collector/consumer"
	"github.com/open-telemetry/opentelemetry-collector/consumer/consumerdata"
	"github.com/open-telemetry/opentelemetry-collector/exporter/exporterhelper"
)

// Exporter struct
type Exporter struct {
	name     string
	client   *influxDbClient.Client
	org      string
	bucket   string
	shutdown exporterhelper.Shutdown
}

var errBlankInfluxDBEndpoint = errors.New("expecting a non-blank endpoint for the InfluxDB configuration")
var errBlankInfluxDBToken = errors.New("expecting a non-blank token for the InfluxDB configuration")
var errBlankInfluxDBOrg = errors.New("expecting a non-blank org for the InfluxDB configuration")
var errBlankInfluxDBBucket = errors.New("expecting a non-blank bucket for the InfluxDB configuration")

var _ consumer.MetricsConsumer = (*Exporter)(nil)

// Start starts
func (e *Exporter) Start(host component.Host) error {
	return nil
}

// ConsumeMetricsData writes metrics to InfluxDB
func (e *Exporter) ConsumeMetricsData(ctx context.Context, md consumerdata.MetricsData) error {
	myMetrics := []influxDbClient.Metric{}

	for _, metric := range md.Metrics {
		switch metric.GetMetricDescriptor().GetType() {
		case metricspb.MetricDescriptor_SUMMARY:

		}
		// append(myMetrics, influxDbClient.NewRowMetric())
		// fmt.Printf("%v", metric.String())
	}

	if _, err := e.client.Write(context.Background(), e.bucket, e.org, myMetrics...); err != nil {
		log.Fatal(err)
	}

	return nil
}

// Shutdown stops the exporter and is invoked during shutdown.
func (e *Exporter) Shutdown() error {
	return e.shutdown()
}
