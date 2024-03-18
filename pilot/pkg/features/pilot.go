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

package features

import (
	"istio.io/libistio/pkg/env"
)

var (
	EnableAlphaGatewayAPI = env.Register("PILOT_ENABLE_ALPHA_GATEWAY_API", false,
		"If this is set to true, support for alpha APIs in the Kubernetes gateway-api (github.com/kubernetes-sigs/gateway-api) will "+
			" be enabled. In addition to this being enabled, the gateway-api CRDs need to be installed.").Get()

	EnableDestinationRuleInheritance = env.Register(
		"PILOT_ENABLE_DESTINATION_RULE_INHERITANCE",
		false,
		"If set, workload specific DestinationRules will inherit configurations settings from mesh and namespace level rules",
	).Get()

	ResolveHostnameGateways = env.Register("RESOLVE_HOSTNAME_GATEWAYS", true,
		"If true, hostnames in the LoadBalancer addresses of a Service will be resolved at the control plane for use in cross-network gateways.").Get()

	EnableQUICListeners = env.Register("PILOT_ENABLE_QUIC_LISTENERS", false,
		"If true, QUIC listeners will be generated wherever there are listeners terminating TLS on gateways "+
			"if the gateway service exposes a UDP port with the same number (for example 443/TCP and 443/UDP)").Get()
)
