// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"github.com/luci/luci-go/grpc/grpcutil"
	"github.com/luci/luci-go/luci_config/common/cfgtypes"
	milo "github.com/luci/luci-go/milo/api/proto"
	"github.com/luci/luci-go/milo/buildsource/buildbot"
	"github.com/luci/luci-go/milo/buildsource/swarming"

	"google.golang.org/grpc/codes"

	"golang.org/x/net/context"
)

// BuildInfoService is a BuildInfoServer implementation.
type BuildInfoService struct {
	// BuildBot is the BuildInfoProvider for the BuildBot service.
	BuildBot buildbot.BuildInfoProvider
	// Swarming is the BuildInfoProvider for the Swarming service.
	Swarming swarming.BuildInfoProvider
}

var _ milo.BuildInfoServer = (*BuildInfoService)(nil)

// Get implements milo.BuildInfoServer.
func (svc *BuildInfoService) Get(c context.Context, req *milo.BuildInfoRequest) (*milo.BuildInfoResponse, error) {
	projectHint := cfgtypes.ProjectName(req.ProjectHint)
	if projectHint != "" {
		if err := projectHint.Validate(); err != nil {
			return nil, grpcutil.Errf(codes.InvalidArgument, "invalid project hint: %s", err.Error())
		}
	}

	switch {
	case req.GetBuildbot() != nil:
		resp, err := svc.BuildBot.GetBuildInfo(c, req.GetBuildbot(), projectHint)
		if err != nil {
			return nil, err
		}
		return resp, nil

	case req.GetSwarming() != nil:
		resp, err := svc.Swarming.GetBuildInfo(c, req.GetSwarming(), projectHint)
		if err != nil {
			return nil, err
		}
		return resp, nil

	default:
		return nil, grpcutil.Errf(codes.InvalidArgument, "must supply a build")
	}
}