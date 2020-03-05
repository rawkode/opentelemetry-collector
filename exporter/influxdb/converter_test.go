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
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	influxDbClient "github.com/influxdata/influxdb-client-go"

	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"github.com/stretchr/testify/assert"
)

func TestConvertToInfluxPoint(t *testing.T) {
	metrics := []*metricspb.Metric{{
		Resource: &v1.Resource{Type: "type", Labels: map[string]string{}},
		MetricDescriptor: &metricspb.MetricDescriptor{
			Description: "Description",
			Name:        "name",
			Type:        metricspb.MetricDescriptor_GAUGE_DOUBLE,
			Unit:        "ss",
			LabelKeys:   []*metricspb.LabelKey{{Key: "k1"}},
		},
		Timeseries: []*metricspb.TimeSeries{
			{
				LabelValues: []*metricspb.LabelValue{
					{Value: "v1"},
				},
				Points: []*metricspb.Point{
					{
						Timestamp: &timestamp.Timestamp{Seconds: time.Now().Unix()},
						Value:     &metricspb.Point_Int64Value{Int64Value: 1},
					},
				},
			},
		},
	}}

	got := toInfluxDbMetric(metrics)

	want := influxDbClient.NewRowMetric(map[string]interface{}{
		"field": 1,
	}, "name", map[string]string{
		"tag": "value",
	}, time.Now())

	assert.Equal(t, want, got)
}
