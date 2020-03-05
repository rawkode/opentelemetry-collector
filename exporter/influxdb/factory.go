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
	"strings"

	influxDbClient "github.com/influxdata/influxdb-client-go"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector/config/configerror"
	"github.com/open-telemetry/opentelemetry-collector/config/configmodels"
	"github.com/open-telemetry/opentelemetry-collector/exporter"
)

const (
	// The value of "type" key in configuration.
	typeStr = "influxdb"
)

// Factory is the factory for InfluxDB exporter.
type Factory struct {
}

// Type gets the type of the Exporter config created by this factory.
func (f *Factory) Type() string {
	return typeStr
}

// CreateDefaultConfig creates the default configuration for exporter.
func (f *Factory) CreateDefaultConfig() configmodels.Exporter {
	return &Config{
		ExporterSettings: configmodels.ExporterSettings{
			TypeVal: typeStr,
			NameVal: typeStr,
		},
	}
}

// CreateTraceExporter creates a trace exporter based on this config.
func (f *Factory) CreateTraceExporter(logger *zap.Logger, config configmodels.Exporter) (exporter.TraceExporter, error) {
	return nil, configerror.ErrDataTypeIsNotSupported
}

// CreateMetricsExporter creates a metrics exporter based on this config.
func (f *Factory) CreateMetricsExporter(logger *zap.Logger, cfg configmodels.Exporter) (exporter.MetricsExporter, error) {
	icfg := cfg.(*Config)

	endpoint := strings.TrimSpace(icfg.Endpoint)
	if endpoint == "" {
		return nil, errBlankInfluxDBEndpoint
	}

	token := strings.TrimSpace(icfg.Token)
	if token == "" {
		return nil, errBlankInfluxDBToken
	}

	org := strings.TrimSpace(icfg.Org)
	if org == "" {
		return nil, errBlankInfluxDBOrg
	}

	bucket := strings.TrimSpace(icfg.Bucket)
	if bucket == "" {
		return nil, errBlankInfluxDBBucket
	}

	client, err := influxDbClient.New(endpoint, token, influxDbClient.Option{})
	if err != nil {
		return nil, err
	}

	exporter := &Exporter{
		name:     cfg.Name(),
		org:      org,
		bucket:   bucket,
		client:   client,
		shutdown: client.Close,
	}

	return exporter, nil
}
