// Copyright 2016 The LUCI Authors.
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

package dmTemplate

import "fmt"

// Normalize will normalize all of the Templates in this message, returning an
// error if any are invalid.
func (f *File) Normalize() error {
	for tempName, t := range f.Template {
		if err := t.Normalize(); err != nil {
			return fmt.Errorf("template %q: %s", tempName, err)
		}
	}
	return nil
}

// Normalize will normalize this Template, returning an error if it is invalid.
func (t *File_Template) Normalize() error {
	if t.DistributorConfigName == "" {
		return fmt.Errorf("missing distributor_config_name")
	}
	if err := t.Parameters.Normalize(); err != nil {
		return err
	}
	if err := t.DistributorParameters.Normalize(); err != nil {
		return err
	}
	return nil
}
