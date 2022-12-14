// Copyright Istio Authors
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

package collections

import (
	"errors"

	"github.com/gogo/protobuf/proto"

	"istio.io/libistio/pkg/config/schema/collection"
	"istio.io/libistio/pkg/config/schema/resource"
	"istio.io/libistio/pkg/test/config"
)

var (
	// Mock is used purely for testing
	Mock = collection.Builder{
		Name:         "mock",
		VariableName: "Mock",
		Resource: resource.Builder{
			ClusterScoped: false,
			Kind:          "MockConfig",
			Plural:        "mockconfigs",
			Group:         "test.istio.io",
			Version:       "v1",
			Proto:         "test.MockConfig",
			ProtoPackage:  "istio.io/istio/pkg/test/config",
			ValidateProto: func(name, namespace string, msg proto.Message) error {
				if msg.(*config.MockConfig).Key == "" {
					return errors.New("empty key")
				}
				return nil
			},
		}.MustBuild(),
	}.MustBuild()

	// Mocks is a Schemas containing the Mock Schema.
	Mocks = collection.NewSchemasBuilder().MustAdd(Mock).Build()
)
