// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fs_config_provider

import (
	"github.com/rskvp/necode/server/config"
)

type (
	fsConfigProvider struct {
		configDir string
		env       string
	}
)

var _ config.ConfigProvider = (*fsConfigProvider)(nil)

// NewFSConfigProvider creates a default file system based Config Provider
func NewFSConfigProvider(configDir string, env string) config.ConfigProvider {
	return &fsConfigProvider{
		configDir,
		env,
	}
}

func (a *fsConfigProvider) GetConfig() (*config.Config, error) {
	cfg, err := LoadConfig(a.configDir, a.env)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}