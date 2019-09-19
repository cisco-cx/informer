// Copyright 2019 Cisco Systems, Inc.
//
// This work incorporates works covered by the following notice:
//
// Copyright 2018 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package v1_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/stretchr/testify/require"

	"github.com/cisco-cx/informer"
	"github.com/cisco-cx/informer/v1"
)

// Confirm that v1.InfoService implements the informer.InfoService interface.
func TestInfoService_Interface(t *testing.T) {
	var _ informer.InfoService = &v1.InfoService{}
	require.Nil(t, nil) // If we get this far, the test passed.
}

// Confirm that v1.InfoService implements the informer.InfoCollector interface.
func TestInfoCollector_Interface(t *testing.T) {
	var _ informer.InfoCollector = &v1.InfoService{}
	require.Nil(t, nil) // If we get this far, the test passed.
}

// Confirm basic informer.InfoService functionality of v1.InfoService.
func TestInfoService_Basic(t *testing.T) {
	info := v1.NewInfoService(
		"program",
		"license",
		"url",
		"user",
		"date",
		"language",
		"languageVersion",
		"version",
		"revision",
		"branch",
	)
	infoExpect := "(metadata=(program=program, license=license, url=url), versionInfo=(version=version, branch=branch, revision=revision), buildInfo=(language=language, languageVersion=languageVersion, user=user, date=date))"
	require.Equal(t, infoExpect, info.String())
}

// Confirm basic informer.InfoCollector functionality of v1.InfoService.
func TestInfoCollector_Basic(t *testing.T) {
	// Prepare to require multiple times.
	require := require.New(t)

	// Get a new service.
	info := v1.NewInfoService(
		"program",
		"license",
		"url",
		"user",
		"date",
		"language",
		"languageVersion",
		"version",
		"revision",
		"branch",
	)

	// Get a new collector.
	c := info.NewCollector()
	err := prometheus.Register(c)
	require.NoError(err)
	defer prometheus.Unregister(c)

	// Parse metrics.
	require.Contains(getMetrics(t), `program_info{branch="branch",build_date="date",build_user="user",language="language",language_version="languageVersion",license="license",program="program",revision="revision",url="url",version="version"} 1`)
}

// getMetrics bootstraps an http server and fetches current metrics.
func getMetrics(t *testing.T) string {
	ts := httptest.NewServer(promhttp.Handler())
	defer ts.Close()
	res, err := http.Get(ts.URL)
	require.NoError(t, err)
	defer res.Body.Close()
	metrics, err := ioutil.ReadAll(res.Body)
	require.NoError(t, err)
	return string(metrics)
}
