// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storagereceiver

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

const componentType = component.Type("jaeger_storage_receiver")

func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		componentType,
		createDefaultConfig,
		receiver.WithTraces(createTraces, component.StabilityLevelDevelopment),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTraces(ctx context.Context, set receiver.CreateSettings, config component.Config, nextConsumer consumer.Traces) (receiver.Traces, error) {
	cfg := config.(*Config)

	return newReceiver(cfg, set.TelemetrySettings, nextConsumer)
}
