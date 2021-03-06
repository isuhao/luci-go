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

package buildbucketpb

import (
	"fmt"
	"strconv"
	"strings"
)

// BuildSet is a parsed buildset tag value.
// It is implemented by *GerritChange, *GitilesCommit.
type BuildSet interface {
	// BuildSetString returns a tag value in buildset format.
	BuildSetString() string
}

// e.g. "patch/gerrit/chromium-review.googlesource.com/677784/5".
func (c *GerritChange) BuildSetString() string {
	return fmt.Sprintf("patch/gerrit/%s/%d/%d", c.Host, c.Change, c.Patchset)
}

// URL returns URL of the change.
func (c *GerritChange) URL() string {
	return fmt.Sprintf("https://%s/c/%d/%d", c.Host, c.Change, c.Patchset)
}

// BuildSetString encodes the commit in buildset tag format,
// e.g. "commit/gitiles/chromium.googlesource.com/infra/luci/luci-go/+/b7a757f457487cd5cfe2dae83f65c5bc10e288b7"
func (c *GitilesCommit) BuildSetString() string {
	return fmt.Sprintf("commit/gitiles/%s/%s/+/%s", c.Host, c.Project, c.Id)
}

// RepoURL returns the URL for the gitiles repo.
// e.g. "https://chromium.googlesource.com/chromium/src"
func (c *GitilesCommit) RepoURL() string {
	return fmt.Sprintf("https://%s/%s", c.Host, c.Project)
}

// URL returns the URL for the gitiles commit.
// e.g. "https://chromium.googlesource.com/chromium/src/+/b7a757f457487cd5cfe2dae83f65c5bc10e288b7"
func (c *GitilesCommit) URL() string {
	return fmt.Sprintf("%s/+/%s", c.RepoURL(), c.Id)
}

// ParseBuildSet tries to parse buildset as one of the known formats.
// If buildSet was not recognized, returns nil.
func ParseBuildSet(buildSet string) BuildSet {
	// fmt.Sscanf cannot be used for this parsing because
	//   var a, b string
	//   fmt.Scanf("a/b", "%s/%s", &a, &b)
	//   a == "a/b", b == ""

	p := strings.Split(buildSet, "/")
	for _, c := range p {
		if c == "" {
			return nil
		}
	}
	n := len(p)
	switch {
	case n == 5 && p[0] == "patch" && p[1] == "gerrit":
		gerrit := &GerritChange{
			Host: p[2],
		}
		var err error
		if gerrit.Change, err = strconv.ParseInt(p[3], 10, 64); err != nil {
			return nil
		}
		if gerrit.Patchset, err = strconv.ParseInt(p[4], 10, 64); err != nil {
			return nil
		}
		return gerrit

	case n >= 5 && p[0] == "commit" && p[1] == "gitiles":
		if p[n-2] != "+" || !looksLikeSha1(p[n-1]) {
			return nil
		}
		return &GitilesCommit{
			Host:    p[2],
			Project: strings.Join(p[3:n-2], "/"), // exclude plus
			Id:      p[n-1],
		}

	default:
		return nil
	}
}

func looksLikeSha1(s string) bool {
	if len(s) != 40 {
		return false
	}
	for _, c := range s {
		switch {
		case '0' <= c && c <= '9':
		case 'a' <= c && c <= 'f':
		default:
			return false
		}
	}
	return true
}
