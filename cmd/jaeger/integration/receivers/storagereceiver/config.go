// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package storagereceiver

import (
	"github.com/asaskevich/govalidator"

	badgerCfg "github.com/jaegertracing/jaeger/plugin/storage/badger"
)

type Config struct {
	Badger badgerCfg.NamespaceConfig `mapstructure:"badger"`
}

func (cfg *Config) Validate() error {
	_, err := govalidator.ValidateStruct(cfg)
	return err
}
