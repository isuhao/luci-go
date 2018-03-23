// Copyright 2015 The LUCI Authors.
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

// Package signing provides interfaces to sign arbitrary small blobs with
// RSA-SHA256 signature (PKCS1v15) and verify such signatures.
//
// Each service has its own private keys it uses for signing, with public
// certificates served over HTTPS. Other services may use the public keys
// to authenticate data generated by the service. It is useful, for example, for
// authenticating PubSub messages payload.
package signing