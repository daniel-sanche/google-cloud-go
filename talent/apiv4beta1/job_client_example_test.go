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

package talent_test

import (
	"context"

	talent "cloud.google.com/go/talent/apiv4beta1"
	"google.golang.org/api/iterator"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewJobClient() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewJobRESTClient() {
	ctx := context.Background()
	c, err := talent.NewJobRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleJobClient_CreateJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.CreateJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#CreateJobRequest.
	}
	resp, err := c.CreateJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_BatchCreateJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.BatchCreateJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#BatchCreateJobsRequest.
	}
	op, err := c.BatchCreateJobs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_GetJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.GetJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#GetJobRequest.
	}
	resp, err := c.GetJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_UpdateJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.UpdateJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#UpdateJobRequest.
	}
	resp, err := c.UpdateJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_BatchUpdateJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.BatchUpdateJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#BatchUpdateJobsRequest.
	}
	op, err := c.BatchUpdateJobs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_DeleteJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.DeleteJobRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#DeleteJobRequest.
	}
	err = c.DeleteJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleJobClient_BatchDeleteJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.BatchDeleteJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#BatchDeleteJobsRequest.
	}
	err = c.BatchDeleteJobs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleJobClient_ListJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.ListJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#ListJobsRequest.
	}
	it := c.ListJobs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleJobClient_SearchJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.SearchJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#SearchJobsRequest.
	}
	resp, err := c.SearchJobs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_SearchJobsForAlert() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.SearchJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/talent/v4beta1#SearchJobsRequest.
	}
	it := c.SearchJobsForAlert(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleJobClient_GetOperation() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
