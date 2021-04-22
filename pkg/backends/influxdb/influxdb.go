/*
 * Copyright 2018 The Trickster Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package influxdb provides the InfluxDB Backend provider
package influxdb

import (
	"net/http"

	"github.com/tricksterproxy/trickster/pkg/backends"
	modelflux "github.com/tricksterproxy/trickster/pkg/backends/influxdb/model"
	bo "github.com/tricksterproxy/trickster/pkg/backends/options"
	"github.com/tricksterproxy/trickster/pkg/backends/providers/registration/types"
	"github.com/tricksterproxy/trickster/pkg/cache"
)

var _ backends.TimeseriesBackend = (*Client)(nil)

// Client Implements the Proxy Client Interface
type Client struct {
	backends.TimeseriesBackend
}

var _ types.NewBackendClientFunc = NewClient

// NewClient returns a new Client Instance
func NewClient(name string, o *bo.Options, router http.Handler,
	cache cache.Cache, _ backends.Backends, _ types.Lookup) (backends.Backend, error) {
	if o != nil {
		o.FastForwardDisable = true
	}
	c := &Client{}
	b, err := backends.NewTimeseriesBackend(name, o, c.RegisterHandlers,
		router, cache, modelflux.NewModeler())
	c.TimeseriesBackend = b
	return c, err
}
