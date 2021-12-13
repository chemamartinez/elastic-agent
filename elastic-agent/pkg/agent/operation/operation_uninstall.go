// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package operation

import (
	"context"

	"github.com/elastic/elastic-agent-poc/elastic-agent/pkg/artifact/uninstall"
	"github.com/elastic/elastic-agent-poc/elastic-agent/pkg/core/logger"
	"github.com/elastic/elastic-agent-poc/elastic-agent/pkg/core/state"
)

// operationUninstall uninstalls a artifact from predefined location
type operationUninstall struct {
	logger      *logger.Logger
	program     Descriptor
	uninstaller uninstall.Uninstaller
}

func newOperationUninstall(
	logger *logger.Logger,
	program Descriptor,
	uninstaller uninstall.Uninstaller) *operationUninstall {

	return &operationUninstall{
		logger:      logger,
		program:     program,
		uninstaller: uninstaller,
	}
}

// Name is human readable name identifying an operation
func (o *operationUninstall) Name() string {
	return "operation-uninstall"
}

// Check checks whether uninstall needs to be ran.
//
// Always true.
func (o *operationUninstall) Check(_ context.Context, _ Application) (bool, error) {
	return true, nil
}

// Run runs the operation
func (o *operationUninstall) Run(ctx context.Context, application Application) (err error) {
	defer func() {
		if err != nil {
			application.SetState(state.Failed, err.Error(), nil)
		}
	}()

	return o.uninstaller.Uninstall(ctx, o.program.Spec(), o.program.Version(), o.program.Directory())
}