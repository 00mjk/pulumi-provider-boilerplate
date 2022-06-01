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
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func log(ctx context.Context, severity diag.Severity, message string) {
	urn, ok := ctx.Value("urn").(resource.URN)
	contract.Assertf(ok, "context missing required value: urn")
	host, ok := ctx.Value("host").(*provider.HostClient)
	contract.Assertf(ok, "context missing required value: host")

	_ = host.LogStatus(ctx, severity, urn, message)
}

func Debug(ctx context.Context, message string) {
	log(ctx, diag.Debug, message)
}

func Info(ctx context.Context, message string) {
	log(ctx, diag.Info, message)
}

func Error(ctx context.Context, message string) {
	log(ctx, diag.Error, message)
}

// ClearStatus will clear the `Info` column of the CLI of all statuses and messages.
func ClearStatus(ctx context.Context) {
	log(ctx, diag.Info, "")
}