// Copyright 2016-2022, Pulumi Corporation.
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

package logging

import (
	"context"

	"github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func Log(ctx context.Context, severity diag.Severity, message string) {
	host, ok := ctx.Value("host").(*provider.HostClient)
	if !ok {
		return
	}
	urn, ok := ctx.Value("urn").(resource.URN)
	if !ok {
		return
	}
	_ = host.LogStatus(ctx, severity, urn, message)
}

// ClearStatus will clear the `Info` column of the CLI of all statuses and messages.
func ClearStatus(ctx context.Context) {
	Log(ctx, diag.Info, "")
}