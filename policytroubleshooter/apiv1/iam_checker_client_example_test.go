// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package policytroubleshooter_test

import (
	"context"

	policytroubleshooter "cloud.google.com/go/policytroubleshooter/apiv1"
	policytroubleshooterpb "cloud.google.com/go/policytroubleshooter/apiv1/policytroubleshooterpb"
)

func ExampleNewIamCheckerClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := policytroubleshooter.NewIamCheckerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleIamCheckerClient_TroubleshootIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := policytroubleshooter.NewIamCheckerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &policytroubleshooterpb.TroubleshootIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/policytroubleshooter/apiv1/policytroubleshooterpb#TroubleshootIamPolicyRequest.
	}
	resp, err := c.TroubleshootIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
