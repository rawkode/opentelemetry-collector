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
	"time"

	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	influxDbClient "github.com/influxdata/influxdb-client-go"
)

func toInfluxDbMetric(metrics []*metricspb.Metric) []influxDbClient.Metric {
	var influxMetrics []influxDbClient.Metric

	for _, metric := range metrics {
		var labelKeys []string
		name := metric.GetMetricDescriptor().GetName()

		for _, labelKey := range metric.GetMetricDescriptor().GetLabelKeys() {
			append(labelKeys, labelKey.GetKey())
		}

		metric.GetMetricDescriptor().GetType()

		for _, timeseries := range metric.GetTimeseries() {
			var labelValues []string

			for _, labelValue := range timeseries.GetLabelValues() {
				append(labelValues, labelValue.GetValue())
			}

			for _, point := range timeseries.Points {
				val := point.GetValue()
				ts := point.GetTimestamp()

				append(influxMetrics, influxDbClient.NewRowMetric(
					map[string]interface{}{"_value": val},
					name,
					map[string]string{},
					time.Now(),
				))
			}
		}
	}
	return influxMetrics
}

func SummaryToInflux(metric metricspb.Metric) influxDbClient.Metric {
	// precondition metrics.MetricDescriptor.Type == MetricDescriptor_SUMMARY
	ts := metric.GetTimeseries()
	_ = ts
	x
	// tag value
	_ = ts[0].LabelValues[0].Value
	summary := ts[0].Points[0].GetSummaryValue()
	// count  and sum field
	_ = summary.Count
	_ = summary.Sum
	// these are tag keys... values ar in the labe values.
	metric.MetricDescriptor.LabelKeys
	metric.MetricDescriptor.Type

	// where the metric originates (tags)
	_ = metric.Resource.Labels

	// host1
	// node1
	ts[0].LabelValues

	/*
		metric {
			metricresource {
				name
				type
				label kv pairs
			}

			timeseries {
				label keys
				points  {
					value
					label values
				}
			}
		}
	*/
}
