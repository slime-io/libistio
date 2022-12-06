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

package validation

import (
	"github.com/gogo/protobuf/proto"
	"istio.io/pkg/log"
)

var (
	scope = log.RegisterScope("validation", "CRD validation debugging", 0)

	_ ValidateFunc = EmptyValidate

	// EmptyValidate is a Validate that does nothing and returns no error.
	EmptyValidate = registerValidateFunc("EmptyValidate",
		func(string, string, proto.Message) error {
			return nil
		})

	validateFuncs = make(map[string]ValidateFunc)
)

// Validate defines a validation func for an API proto.
type ValidateFunc func(name, namespace string, config proto.Message) error

func registerValidateFunc(name string, f ValidateFunc) ValidateFunc {
	validateFuncs[name] = f
	return f
}

// IsValidateFunc indicates whether there is a validation function with the given name.
func IsValidateFunc(name string) bool {
	return GetValidateFunc(name) != nil
}

// GetValidateFunc returns the validation function with the given name, or null if it does not exist.
func GetValidateFunc(name string) ValidateFunc {
	return validateFuncs[name]
}
