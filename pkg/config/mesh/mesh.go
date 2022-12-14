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

package mesh

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/gogo/protobuf/proto"

	"istio.io/api/networking/v1alpha3"

	"github.com/gogo/protobuf/types"
	"github.com/hashicorp/go-multierror"

	"istio.io/pkg/filewatcher"
	"istio.io/pkg/log"

	meshconfig "istio.io/api/mesh/v1alpha1"

	"istio.io/libistio/pkg/config/constants"
	"istio.io/libistio/pkg/util/gogoprotomarshal"
)

// DefaultProxyConfig for individual proxies
func DefaultProxyConfig() meshconfig.ProxyConfig {
	// TODO: include revision based on REVISION env
	// TODO: set default namespace based on POD_NAMESPACE env
	return meshconfig.ProxyConfig{
		// missing: ConnectTimeout: 10 * time.Second,
		ConfigPath:             constants.ConfigPathDir,
		ServiceCluster:         constants.ServiceClusterName,
		DrainDuration:          types.DurationProto(45 * time.Second),
		ParentShutdownDuration: types.DurationProto(60 * time.Second),
		ProxyAdminPort:         15000,
		Concurrency:            &types.Int32Value{Value: 2},
		ControlPlaneAuthPolicy: meshconfig.AuthenticationPolicy_NONE,
		DiscoveryAddress:       "istiod.istio-system.svc:15012",
		Tracing: &meshconfig.Tracing{
			Tracer: &meshconfig.Tracing_Zipkin_{
				Zipkin: &meshconfig.Tracing_Zipkin{
					Address: "zipkin.istio-system:9411",
				},
			},
		},

		// Code defaults
		BinaryPath:            constants.BinaryPathFilename,
		StatsdUdpAddress:      "",
		EnvoyMetricsService:   &meshconfig.RemoteService{Address: ""},
		EnvoyAccessLogService: &meshconfig.RemoteService{Address: ""},
		CustomConfigFile:      "",
		StatNameLength:        189,
		StatusPort:            15020,
	}
}

// DefaultMeshConfig returns the default mesh config.
// This is merged with values from the mesh config map.
func DefaultMeshConfig() meshconfig.MeshConfig {
	proxyConfig := DefaultProxyConfig()

	// Defaults matching the standard install
	// order matches the generated mesh config.
	return meshconfig.MeshConfig{
		EnableTracing:               true,
		AccessLogFile:               "",
		AccessLogEncoding:           meshconfig.MeshConfig_TEXT,
		AccessLogFormat:             "",
		EnableEnvoyAccessLogService: false,
		ProtocolDetectionTimeout:    types.DurationProto(100 * time.Millisecond),
		IngressService:              "istio-ingressgateway",
		IngressControllerMode:       meshconfig.MeshConfig_STRICT,
		IngressClass:                "istio",
		TrustDomain:                 "cluster.local",
		TrustDomainAliases:          []string{},
		EnableAutoMtls:              &types.BoolValue{Value: true},
		OutboundTrafficPolicy:       &meshconfig.MeshConfig_OutboundTrafficPolicy{Mode: meshconfig.MeshConfig_OutboundTrafficPolicy_ALLOW_ANY},
		LocalityLbSetting: &v1alpha3.LocalityLoadBalancerSetting{
			Enabled: &types.BoolValue{Value: true},
		},
		Certificates:  []*meshconfig.Certificate{},
		DefaultConfig: &proxyConfig,

		// Not set in the default mesh config - code defaults.
		RootNamespace:                  constants.IstioSystemNamespace,
		ProxyListenPort:                15001,
		ConnectTimeout:                 types.DurationProto(10 * time.Second),
		DefaultServiceExportTo:         []string{"*"},
		DefaultVirtualServiceExportTo:  []string{"*"},
		DefaultDestinationRuleExportTo: []string{"*"},
		DnsRefreshRate:                 types.DurationProto(5 * time.Second), // 5 seconds is the default refresh rate used in Envoy
		ThriftConfig:                   &meshconfig.MeshConfig_ThriftConfig{},
		ServiceSettings:                make([]*meshconfig.MeshConfig_ServiceSettings, 0),
	}
}

// ApplyProxyConfig applies the give proxy config yaml to a mesh config object. The passed in mesh config
// will not be modified.
func ApplyProxyConfig(yaml string, meshConfig meshconfig.MeshConfig) (*meshconfig.MeshConfig, error) {
	mc := proto.Clone(&meshConfig).(*meshconfig.MeshConfig)
	if err := gogoprotomarshal.ApplyYAML(yaml, mc.DefaultConfig); err != nil {
		return nil, fmt.Errorf("could not parse proxy config: %v", err)
	}
	return mc, nil
}

// ApplyMeshConfig returns a new MeshConfig decoded from the
// input YAML with the provided defaults applied to omitted configuration values.
func ApplyMeshConfig(yaml string, defaultConfig meshconfig.MeshConfig) (*meshconfig.MeshConfig, error) {
	if err := gogoprotomarshal.ApplyYAML(yaml, &defaultConfig); err != nil {
		return nil, multierror.Prefix(err, "failed to convert to proto.")
	}

	// Reset the default ProxyConfig as jsonpb.UnmarshalString doesn't
	// handled nested decode properly for our use case.
	prevDefaultConfig := defaultConfig.DefaultConfig
	defaultProxyConfig := DefaultProxyConfig()
	defaultConfig.DefaultConfig = &defaultProxyConfig

	// Re-apply defaults to ProxyConfig if they were defined in the
	// original input MeshConfig.ProxyConfig.
	if prevDefaultConfig != nil {
		origProxyConfigYAML, err := gogoprotomarshal.ToYAML(prevDefaultConfig)
		if err != nil {
			return nil, multierror.Prefix(err, "failed to re-encode default proxy config")
		}
		if err := gogoprotomarshal.ApplyYAML(origProxyConfigYAML, defaultConfig.DefaultConfig); err != nil {
			return nil, multierror.Prefix(err, "failed to convert to proto.")
		}
	}

	return &defaultConfig, nil
}

// ApplyMeshConfigDefaults returns a new MeshConfig decoded from the
// input YAML with defaults applied to omitted configuration values.
func ApplyMeshConfigDefaults(yaml string) (*meshconfig.MeshConfig, error) {
	return ApplyMeshConfig(yaml, DefaultMeshConfig())
}

// EmptyMeshNetworks configuration with no networks
func EmptyMeshNetworks() meshconfig.MeshNetworks {
	return meshconfig.MeshNetworks{
		Networks: map[string]*meshconfig.Network{},
	}
}

// ParseMeshNetworks returns a new MeshNetworks decoded from the
// input YAML.
func ParseMeshNetworks(yaml string) (*meshconfig.MeshNetworks, error) {
	out := EmptyMeshNetworks()
	if err := gogoprotomarshal.ApplyYAML(yaml, &out); err != nil {
		return nil, multierror.Prefix(err, "failed to convert to proto.")
	}

	return &out, nil
}

// ReadMeshNetworks gets mesh networks configuration from a config file
func ReadMeshNetworks(filename string) (*meshconfig.MeshNetworks, error) {
	yaml, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, multierror.Prefix(err, "cannot read networks config file")
	}
	return ParseMeshNetworks(string(yaml))
}

// ReadMeshConfig gets mesh configuration from a config file
func ReadMeshConfig(filename string) (*meshconfig.MeshConfig, error) {
	yaml, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, multierror.Prefix(err, "cannot read mesh config file")
	}
	return ApplyMeshConfigDefaults(string(yaml))
}

// ResolveHostsInNetworksConfig will go through the Gateways addresses for all
// networks in the config and if it's not an IP address it will try to lookup
// that hostname and replace it with the IP address in the config
func ResolveHostsInNetworksConfig(config *meshconfig.MeshNetworks) {
	if config == nil {
		return
	}
	for _, n := range config.Networks {
		for _, gw := range n.Gateways {
			gwAddr := gw.GetAddress()
			gwIP := net.ParseIP(gwAddr)
			if gwIP == nil && len(gwAddr) != 0 {
				addrs, err := net.LookupHost(gwAddr)
				if err != nil {
					log.Warnf("error resolving host %#v: %v", gw.GetAddress(), err)
				} else {
					gw.Gw = &meshconfig.Network_IstioNetworkGateway_Address{
						Address: addrs[0],
					}
				}
			}
		}
	}
}

// Add to the FileWatcher the provided file and execute the provided function
// on any change event for this file.
// Using a debouncing mechanism to avoid calling the callback multiple times
// per event.
func addFileWatcher(fileWatcher filewatcher.FileWatcher, file string, callback func()) {
	_ = fileWatcher.Add(file)
	go func() {
		var timerC <-chan time.Time
		for {
			select {
			case <-timerC:
				timerC = nil
				callback()
			case <-fileWatcher.Events(file):
				// Use a timer to debounce configuration updates
				if timerC == nil {
					timerC = time.After(100 * time.Millisecond)
				}
			}
		}
	}()
}
