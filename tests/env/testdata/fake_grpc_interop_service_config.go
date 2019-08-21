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
	"strings"

	"cloudesf.googlesource.com/gcpproxy/src/go/util"
	"github.com/golang/glog"
	"github.com/golang/protobuf/jsonpb"

	conf "google.golang.org/genproto/googleapis/api/serviceconfig"
)

var (
	FakeGRPCInteropConfig = createFakeGRPCInteropConfig()
)

func createFakeGRPCInteropConfig() *conf.Service {
	unmarshaler := &jsonpb.Unmarshaler{
		AllowUnknownFields: true,
		AnyResolver:        util.Resolver,
	}
	var serviceConfig conf.Service
	if err := unmarshaler.Unmarshal(strings.NewReader(grpcInteropServiceConfigJsonStr), &serviceConfig); err != nil {
		glog.Errorf("fail to unmarshal serviceConfig, %s", err)
		return nil
	}
	return &serviceConfig
}
