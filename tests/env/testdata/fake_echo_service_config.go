// Copyright 2018 Google Cloud Platform Proxy Authors
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

package testdata

import (
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/protobuf/api"

	conf "google.golang.org/genproto/googleapis/api/serviceconfig"
	ptype "google.golang.org/genproto/protobuf/ptype"
)

var (
	FakeEchoConfig = &conf.Service{
		Name:              "echo-api.endpoints.cloudesf-testing.cloud.goog",
		Title:             "Endpoints Example",
		ProducerProjectId: "producer-project",
		Apis: []*api.Api{
			{
				Name: "1.echo_api_endpoints_cloudesf_testing_cloud_goog",
				Methods: []*api.Method{
					{
						Name:            "Auth_info_google_jwt",
						RequestTypeUrl:  "type.googleapis.com/google.protobuf.Empty",
						ResponseTypeUrl: "type.googleapis.com/AuthInfoResponse",
					},
					{
						Name:            "Echo",
						RequestTypeUrl:  "type.googleapis.com/EchoRequest",
						ResponseTypeUrl: "type.googleapis.com/EchoMessage",
					},
					{
						Name:            "Simplegetcors",
						RequestTypeUrl:  "type.googleapis.com/google.protobuf.Empty",
						ResponseTypeUrl: "type.googleapis.com/SimpleCorsMessage",
					},
					{
						Name:            "Auth_info_firebase",
						RequestTypeUrl:  "type.googleapis.com/google.protobuf.Empty",
						ResponseTypeUrl: "type.googleapis.com/AuthInfoResponse",
					},
				},
				Version: "1.0.0",
			},
		},
		Http: &annotations.Http{
			Rules: []*annotations.HttpRule{
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Auth_info_google_jwt",
					Pattern: &annotations.HttpRule_Get{
						Get: "/auth/info/googlejwt",
					},
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Auth0",
					Pattern: &annotations.HttpRule_Get{
						Get: "/auth/info/auth0",
					},
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Echo",
					Pattern: &annotations.HttpRule_Post{
						Post: "/echo",
					},
					Body: "message",
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Echo_nokey",
					Pattern: &annotations.HttpRule_Post{
						Post: "/echo/nokey",
					},
					Body: "message",
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Simplegetcors",
					Pattern: &annotations.HttpRule_Get{
						Get: "/simplegetcors",
					},
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog._post_anypath",
					Pattern: &annotations.HttpRule_Post{
						Post: "/**",
					},
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Auth_info_firebase",
					Pattern: &annotations.HttpRule_Get{
						Get: "/auth/info/firebase",
					},
				},
			},
		},
		Types: []*ptype.Type{
			{
				Fields: []*ptype.Field{
					&ptype.Field{
						JsonName: "BOOK",
						Name:     "b_o_o_k",
					},
					&ptype.Field{
						JsonName: "SHELF",
						Name:     "s_h_e_l_f",
					},
				},
			},
		},
		Authentication: &conf.Authentication{
			Rules: []*conf.AuthenticationRule{
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Auth_info_google_jwt",
					Requirements: []*conf.AuthRequirement{
						{
							ProviderId: GoogleJwtProvider,
						},
					},
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Auth0",
					Requirements: []*conf.AuthRequirement{
						{
							ProviderId: GoogleJwtProvider,
							Audiences:  "admin.cloud.goog",
						},
					},
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Echo",
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Simplegetcors",
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog._post_anypath",
				},
				{
					Selector: "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Auth_info_firebase",
					Requirements: []*conf.AuthRequirement{
						{
							ProviderId: GoogleJwtProvider,
						},
					},
				},
			},
		},
		Usage: &conf.Usage{
			Rules: []*conf.UsageRule{
				{
					Selector:               "1.echo_api_endpoints_cloudesf_testing_cloud_goog.Echo_nokey",
					AllowUnregisteredCalls: true,
				},
				{
					Selector:               "1.echo_api_endpoints_cloudesf_testing_cloud_goog._post_anypath",
					AllowUnregisteredCalls: true,
				},
			},
		},
		Endpoints: []*conf.Endpoint{
			{
				Name: "echo-api.endpoints.cloudesf-testing.cloud.goog",
			},
		},
	}
)
