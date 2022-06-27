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

package dataqna

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	dataqnapb "google.golang.org/genproto/googleapis/cloud/dataqna/v1alpha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newQuestionClientHook clientHook

// QuestionCallOptions contains the retry settings for each method of QuestionClient.
type QuestionCallOptions struct {
	GetQuestion        []gax.CallOption
	CreateQuestion     []gax.CallOption
	ExecuteQuestion    []gax.CallOption
	GetUserFeedback    []gax.CallOption
	UpdateUserFeedback []gax.CallOption
}

func defaultQuestionGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataqna.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataqna.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataqna.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultQuestionCallOptions() *QuestionCallOptions {
	return &QuestionCallOptions{
		GetQuestion: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateQuestion:  []gax.CallOption{},
		ExecuteQuestion: []gax.CallOption{},
		GetUserFeedback: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateUserFeedback: []gax.CallOption{},
	}
}

func defaultQuestionRESTCallOptions() *QuestionCallOptions {
	return &QuestionCallOptions{
		GetQuestion: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		CreateQuestion:  []gax.CallOption{},
		ExecuteQuestion: []gax.CallOption{},
		GetUserFeedback: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		UpdateUserFeedback: []gax.CallOption{},
	}
}

// internalQuestionClient is an interface that defines the methods available from Data QnA API.
type internalQuestionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetQuestion(context.Context, *dataqnapb.GetQuestionRequest, ...gax.CallOption) (*dataqnapb.Question, error)
	CreateQuestion(context.Context, *dataqnapb.CreateQuestionRequest, ...gax.CallOption) (*dataqnapb.Question, error)
	ExecuteQuestion(context.Context, *dataqnapb.ExecuteQuestionRequest, ...gax.CallOption) (*dataqnapb.Question, error)
	GetUserFeedback(context.Context, *dataqnapb.GetUserFeedbackRequest, ...gax.CallOption) (*dataqnapb.UserFeedback, error)
	UpdateUserFeedback(context.Context, *dataqnapb.UpdateUserFeedbackRequest, ...gax.CallOption) (*dataqnapb.UserFeedback, error)
}

// QuestionClient is a client for interacting with Data QnA API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to interpret natural language queries.
// The service allows to create Question resources that are interpreted and
// are filled with one or more interpretations if the question could be
// interpreted. Once a Question resource is created and has at least one
// interpretation, an interpretation can be chosen for execution, which
// triggers a query to the backend (for BigQuery, it will create a job).
// Upon successful execution of that interpretation, backend specific
// information will be returned so that the client can retrieve the results
// from the backend.
//
// The Question resources are named projects/*/locations/*/questions/*.
//
// The Question resource has a singletion sub-resource UserFeedback named
// projects/*/locations/*/questions/*/userFeedback, which allows access to
// user feedback.
type QuestionClient struct {
	// The internal transport-dependent client.
	internalClient internalQuestionClient

	// The call options for this service.
	CallOptions *QuestionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *QuestionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *QuestionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *QuestionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetQuestion gets a previously created question.
func (c *QuestionClient) GetQuestion(ctx context.Context, req *dataqnapb.GetQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	return c.internalClient.GetQuestion(ctx, req, opts...)
}

// CreateQuestion creates a question.
func (c *QuestionClient) CreateQuestion(ctx context.Context, req *dataqnapb.CreateQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	return c.internalClient.CreateQuestion(ctx, req, opts...)
}

// ExecuteQuestion executes an interpretation.
func (c *QuestionClient) ExecuteQuestion(ctx context.Context, req *dataqnapb.ExecuteQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	return c.internalClient.ExecuteQuestion(ctx, req, opts...)
}

// GetUserFeedback gets previously created user feedback.
func (c *QuestionClient) GetUserFeedback(ctx context.Context, req *dataqnapb.GetUserFeedbackRequest, opts ...gax.CallOption) (*dataqnapb.UserFeedback, error) {
	return c.internalClient.GetUserFeedback(ctx, req, opts...)
}

// UpdateUserFeedback updates user feedback. This creates user feedback if there was none before
// (upsert).
func (c *QuestionClient) UpdateUserFeedback(ctx context.Context, req *dataqnapb.UpdateUserFeedbackRequest, opts ...gax.CallOption) (*dataqnapb.UserFeedback, error) {
	return c.internalClient.UpdateUserFeedback(ctx, req, opts...)
}

// questionGRPCClient is a client for interacting with Data QnA API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type questionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing QuestionClient
	CallOptions **QuestionCallOptions

	// The gRPC API client.
	questionClient dataqnapb.QuestionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewQuestionClient creates a new question service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to interpret natural language queries.
// The service allows to create Question resources that are interpreted and
// are filled with one or more interpretations if the question could be
// interpreted. Once a Question resource is created and has at least one
// interpretation, an interpretation can be chosen for execution, which
// triggers a query to the backend (for BigQuery, it will create a job).
// Upon successful execution of that interpretation, backend specific
// information will be returned so that the client can retrieve the results
// from the backend.
//
// The Question resources are named projects/*/locations/*/questions/*.
//
// The Question resource has a singletion sub-resource UserFeedback named
// projects/*/locations/*/questions/*/userFeedback, which allows access to
// user feedback.
func NewQuestionClient(ctx context.Context, opts ...option.ClientOption) (*QuestionClient, error) {
	clientOpts := defaultQuestionGRPCClientOptions()
	if newQuestionClientHook != nil {
		hookOpts, err := newQuestionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := QuestionClient{CallOptions: defaultQuestionCallOptions()}

	c := &questionGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		questionClient:   dataqnapb.NewQuestionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *questionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *questionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *questionGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type questionRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing QuestionClient
	CallOptions **QuestionCallOptions
}

// NewQuestionRESTClient creates a new question service rest client.
//
// Service to interpret natural language queries.
// The service allows to create Question resources that are interpreted and
// are filled with one or more interpretations if the question could be
// interpreted. Once a Question resource is created and has at least one
// interpretation, an interpretation can be chosen for execution, which
// triggers a query to the backend (for BigQuery, it will create a job).
// Upon successful execution of that interpretation, backend specific
// information will be returned so that the client can retrieve the results
// from the backend.
//
// The Question resources are named projects/*/locations/*/questions/*.
//
// The Question resource has a singletion sub-resource UserFeedback named
// projects/*/locations/*/questions/*/userFeedback, which allows access to
// user feedback.
func NewQuestionRESTClient(ctx context.Context, opts ...option.ClientOption) (*QuestionClient, error) {
	clientOpts := append(defaultQuestionRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultQuestionRESTCallOptions()
	c := &questionRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &QuestionClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultQuestionRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dataqna.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://dataqna.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://dataqna.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *questionRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *questionRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *questionRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *questionGRPCClient) GetQuestion(ctx context.Context, req *dataqnapb.GetQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetQuestion[0:len((*c.CallOptions).GetQuestion):len((*c.CallOptions).GetQuestion)], opts...)
	var resp *dataqnapb.Question
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.questionClient.GetQuestion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *questionGRPCClient) CreateQuestion(ctx context.Context, req *dataqnapb.CreateQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateQuestion[0:len((*c.CallOptions).CreateQuestion):len((*c.CallOptions).CreateQuestion)], opts...)
	var resp *dataqnapb.Question
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.questionClient.CreateQuestion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *questionGRPCClient) ExecuteQuestion(ctx context.Context, req *dataqnapb.ExecuteQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ExecuteQuestion[0:len((*c.CallOptions).ExecuteQuestion):len((*c.CallOptions).ExecuteQuestion)], opts...)
	var resp *dataqnapb.Question
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.questionClient.ExecuteQuestion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *questionGRPCClient) GetUserFeedback(ctx context.Context, req *dataqnapb.GetUserFeedbackRequest, opts ...gax.CallOption) (*dataqnapb.UserFeedback, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetUserFeedback[0:len((*c.CallOptions).GetUserFeedback):len((*c.CallOptions).GetUserFeedback)], opts...)
	var resp *dataqnapb.UserFeedback
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.questionClient.GetUserFeedback(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *questionGRPCClient) UpdateUserFeedback(ctx context.Context, req *dataqnapb.UpdateUserFeedbackRequest, opts ...gax.CallOption) (*dataqnapb.UserFeedback, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "user_feedback.name", url.QueryEscape(req.GetUserFeedback().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateUserFeedback[0:len((*c.CallOptions).UpdateUserFeedback):len((*c.CallOptions).UpdateUserFeedback)], opts...)
	var resp *dataqnapb.UserFeedback
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.questionClient.UpdateUserFeedback(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetQuestion gets a previously created question.
func (c *questionRESTClient) GetQuestion(ctx context.Context, req *dataqnapb.GetQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetName())

	params := url.Values{}
	if req.GetReadMask().GetPaths() != nil {
		params.Add("readMask.paths", fmt.Sprintf("%v", req.GetReadMask().GetPaths()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetQuestion[0:len((*c.CallOptions).GetQuestion):len((*c.CallOptions).GetQuestion)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataqnapb.Question{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// CreateQuestion creates a question.
func (c *questionRESTClient) CreateQuestion(ctx context.Context, req *dataqnapb.CreateQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetQuestion()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v/questions", req.GetParent())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateQuestion[0:len((*c.CallOptions).CreateQuestion):len((*c.CallOptions).CreateQuestion)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataqnapb.Question{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ExecuteQuestion executes an interpretation.
func (c *questionRESTClient) ExecuteQuestion(ctx context.Context, req *dataqnapb.ExecuteQuestionRequest, opts ...gax.CallOption) (*dataqnapb.Question, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v:execute", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).ExecuteQuestion[0:len((*c.CallOptions).ExecuteQuestion):len((*c.CallOptions).ExecuteQuestion)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataqnapb.Question{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetUserFeedback gets previously created user feedback.
func (c *questionRESTClient) GetUserFeedback(ctx context.Context, req *dataqnapb.GetUserFeedbackRequest, opts ...gax.CallOption) (*dataqnapb.UserFeedback, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).GetUserFeedback[0:len((*c.CallOptions).GetUserFeedback):len((*c.CallOptions).GetUserFeedback)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataqnapb.UserFeedback{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// UpdateUserFeedback updates user feedback. This creates user feedback if there was none before
// (upsert).
func (c *questionRESTClient) UpdateUserFeedback(ctx context.Context, req *dataqnapb.UpdateUserFeedbackRequest, opts ...gax.CallOption) (*dataqnapb.UserFeedback, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetUserFeedback()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1alpha/%v", req.GetUserFeedback().GetName())

	params := url.Values{}
	if req.GetUpdateMask().GetPaths() != nil {
		params.Add("updateMask.paths", fmt.Sprintf("%v", req.GetUpdateMask().GetPaths()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "user_feedback.name", url.QueryEscape(req.GetUserFeedback().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).UpdateUserFeedback[0:len((*c.CallOptions).UpdateUserFeedback):len((*c.CallOptions).UpdateUserFeedback)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataqnapb.UserFeedback{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
