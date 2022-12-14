// GENERATED FILE -- DO NOT EDIT
//

package collections

import (
	"istio.io/libistio/pkg/config/schema/collection"
	"istio.io/libistio/pkg/config/schema/resource"
	"istio.io/libistio/pkg/config/validation"
)

var (

	// IstioConfigV1Alpha2Adapters describes the collection
	// istio/config/v1alpha2/adapters
	IstioConfigV1Alpha2Adapters = collection.Builder{
		Name:         "istio/config/v1alpha2/adapters",
		VariableName: "IstioConfigV1Alpha2Adapters",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "config.istio.io",
			Kind:          "adapter",
			Plural:        "adapters",
			Version:       "v1alpha2",
			Proto:         "google.protobuf.Struct",
			ProtoPackage:  "github.com/gogo/protobuf/types",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioConfigV1Alpha2Templates describes the collection
	// istio/config/v1alpha2/templates
	IstioConfigV1Alpha2Templates = collection.Builder{
		Name:         "istio/config/v1alpha2/templates",
		VariableName: "IstioConfigV1Alpha2Templates",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "config.istio.io",
			Kind:          "template",
			Plural:        "templates",
			Version:       "v1alpha2",
			Proto:         "google.protobuf.Struct",
			ProtoPackage:  "github.com/gogo/protobuf/types",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioMeshV1Alpha1MeshConfig describes the collection
	// istio/mesh/v1alpha1/MeshConfig
	IstioMeshV1Alpha1MeshConfig = collection.Builder{
		Name:         "istio/mesh/v1alpha1/MeshConfig",
		VariableName: "IstioMeshV1Alpha1MeshConfig",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "MeshConfig",
			Plural:        "meshconfigs",
			Version:       "v1alpha1",
			Proto:         "istio.mesh.v1alpha1.MeshConfig",
			ProtoPackage:  "istio.io/api/mesh/v1alpha1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioMeshV1Alpha1MeshNetworks describes the collection
	// istio/mesh/v1alpha1/MeshNetworks
	IstioMeshV1Alpha1MeshNetworks = collection.Builder{
		Name:         "istio/mesh/v1alpha1/MeshNetworks",
		VariableName: "IstioMeshV1Alpha1MeshNetworks",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "MeshNetworks",
			Plural:        "meshnetworks",
			Version:       "v1alpha1",
			Proto:         "istio.mesh.v1alpha1.MeshNetworks",
			ProtoPackage:  "istio.io/api/mesh/v1alpha1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioNetworkingV1Alpha3Destinationrules describes the collection
	// istio/networking/v1alpha3/destinationrules
	IstioNetworkingV1Alpha3Destinationrules = collection.Builder{
		Name:         "istio/networking/v1alpha3/destinationrules",
		VariableName: "IstioNetworkingV1Alpha3Destinationrules",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "DestinationRule",
			Plural:        "destinationrules",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.DestinationRule",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioNetworkingV1Alpha3Envoyfilters describes the collection
	// istio/networking/v1alpha3/envoyfilters
	IstioNetworkingV1Alpha3Envoyfilters = collection.Builder{
		Name:         "istio/networking/v1alpha3/envoyfilters",
		VariableName: "IstioNetworkingV1Alpha3Envoyfilters",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "EnvoyFilter",
			Plural:        "envoyfilters",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.EnvoyFilter",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioNetworkingV1Alpha3Gateways describes the collection
	// istio/networking/v1alpha3/gateways
	IstioNetworkingV1Alpha3Gateways = collection.Builder{
		Name:         "istio/networking/v1alpha3/gateways",
		VariableName: "IstioNetworkingV1Alpha3Gateways",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "Gateway",
			Plural:        "gateways",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.Gateway",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioNetworkingV1Alpha3Serviceentries describes the collection
	// istio/networking/v1alpha3/serviceentries
	IstioNetworkingV1Alpha3Serviceentries = collection.Builder{
		Name:         "istio/networking/v1alpha3/serviceentries",
		VariableName: "IstioNetworkingV1Alpha3Serviceentries",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "ServiceEntry",
			Plural:        "serviceentries",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.ServiceEntry",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioNetworkingV1Alpha3Sidecars describes the collection
	// istio/networking/v1alpha3/sidecars
	IstioNetworkingV1Alpha3Sidecars = collection.Builder{
		Name:         "istio/networking/v1alpha3/sidecars",
		VariableName: "IstioNetworkingV1Alpha3Sidecars",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "Sidecar",
			Plural:        "sidecars",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.Sidecar",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioNetworkingV1Alpha3Virtualservices describes the collection
	// istio/networking/v1alpha3/virtualservices
	IstioNetworkingV1Alpha3Virtualservices = collection.Builder{
		Name:         "istio/networking/v1alpha3/virtualservices",
		VariableName: "IstioNetworkingV1Alpha3Virtualservices",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "VirtualService",
			Plural:        "virtualservices",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.VirtualService",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioNetworkingV1Alpha3Workloadentries describes the collection
	// istio/networking/v1alpha3/workloadentries
	IstioNetworkingV1Alpha3Workloadentries = collection.Builder{
		Name:         "istio/networking/v1alpha3/workloadentries",
		VariableName: "IstioNetworkingV1Alpha3Workloadentries",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "WorkloadEntry",
			Plural:        "workloadentries",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.WorkloadEntry",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioSecurityV1Beta1Authorizationpolicies describes the collection
	// istio/security/v1beta1/authorizationpolicies
	IstioSecurityV1Beta1Authorizationpolicies = collection.Builder{
		Name:         "istio/security/v1beta1/authorizationpolicies",
		VariableName: "IstioSecurityV1Beta1Authorizationpolicies",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "security.istio.io",
			Kind:          "AuthorizationPolicy",
			Plural:        "authorizationpolicies",
			Version:       "v1beta1",
			Proto:         "istio.security.v1beta1.AuthorizationPolicy",
			ProtoPackage:  "istio.io/api/security/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioSecurityV1Beta1Peerauthentications describes the collection
	// istio/security/v1beta1/peerauthentications
	IstioSecurityV1Beta1Peerauthentications = collection.Builder{
		Name:         "istio/security/v1beta1/peerauthentications",
		VariableName: "IstioSecurityV1Beta1Peerauthentications",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "security.istio.io",
			Kind:          "PeerAuthentication",
			Plural:        "peerauthentications",
			Version:       "v1beta1",
			Proto:         "istio.security.v1beta1.PeerAuthentication",
			ProtoPackage:  "istio.io/api/security/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// IstioSecurityV1Beta1Requestauthentications describes the collection
	// istio/security/v1beta1/requestauthentications
	IstioSecurityV1Beta1Requestauthentications = collection.Builder{
		Name:         "istio/security/v1beta1/requestauthentications",
		VariableName: "IstioSecurityV1Beta1Requestauthentications",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "security.istio.io",
			Kind:          "RequestAuthentication",
			Plural:        "requestauthentications",
			Version:       "v1beta1",
			Proto:         "istio.security.v1beta1.RequestAuthentication",
			ProtoPackage:  "istio.io/api/security/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SApiextensionsK8SIoV1Beta1Customresourcedefinitions describes the
	// collection k8s/apiextensions.k8s.io/v1beta1/customresourcedefinitions
	K8SApiextensionsK8SIoV1Beta1Customresourcedefinitions = collection.Builder{
		Name:         "k8s/apiextensions.k8s.io/v1beta1/customresourcedefinitions",
		VariableName: "K8SApiextensionsK8SIoV1Beta1Customresourcedefinitions",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "apiextensions.k8s.io",
			Kind:          "CustomResourceDefinition",
			Plural:        "CustomResourceDefinitions",
			Version:       "v1beta1",
			Proto:         "k8s.io.apiextensions_apiserver.pkg.apis.apiextensions.v1beta1.CustomResourceDefinition",
			ProtoPackage:  "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SAppsV1Deployments describes the collection k8s/apps/v1/deployments
	K8SAppsV1Deployments = collection.Builder{
		Name:         "k8s/apps/v1/deployments",
		VariableName: "K8SAppsV1Deployments",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "apps",
			Kind:          "Deployment",
			Plural:        "Deployments",
			Version:       "v1",
			Proto:         "k8s.io.api.apps.v1.Deployment",
			ProtoPackage:  "k8s.io/api/apps/v1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SConfigIstioIoV1Alpha2Adapters describes the collection
	// k8s/config.istio.io/v1alpha2/adapters
	K8SConfigIstioIoV1Alpha2Adapters = collection.Builder{
		Name:         "k8s/config.istio.io/v1alpha2/adapters",
		VariableName: "K8SConfigIstioIoV1Alpha2Adapters",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "config.istio.io",
			Kind:          "adapter",
			Plural:        "adapters",
			Version:       "v1alpha2",
			Proto:         "google.protobuf.Struct",
			ProtoPackage:  "github.com/gogo/protobuf/types",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SConfigIstioIoV1Alpha2Templates describes the collection
	// k8s/config.istio.io/v1alpha2/templates
	K8SConfigIstioIoV1Alpha2Templates = collection.Builder{
		Name:         "k8s/config.istio.io/v1alpha2/templates",
		VariableName: "K8SConfigIstioIoV1Alpha2Templates",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "config.istio.io",
			Kind:          "template",
			Plural:        "templates",
			Version:       "v1alpha2",
			Proto:         "google.protobuf.Struct",
			ProtoPackage:  "github.com/gogo/protobuf/types",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SCoreV1Configmaps describes the collection k8s/core/v1/configmaps
	K8SCoreV1Configmaps = collection.Builder{
		Name:         "k8s/core/v1/configmaps",
		VariableName: "K8SCoreV1Configmaps",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "ConfigMap",
			Plural:        "configmaps",
			Version:       "v1",
			Proto:         "k8s.io.api.core.v1.ConfigMap",
			ProtoPackage:  "k8s.io/api/core/v1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SCoreV1Endpoints describes the collection k8s/core/v1/endpoints
	K8SCoreV1Endpoints = collection.Builder{
		Name:         "k8s/core/v1/endpoints",
		VariableName: "K8SCoreV1Endpoints",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "Endpoints",
			Plural:        "endpoints",
			Version:       "v1",
			Proto:         "k8s.io.api.core.v1.Endpoints",
			ProtoPackage:  "k8s.io/api/core/v1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SCoreV1Namespaces describes the collection k8s/core/v1/namespaces
	K8SCoreV1Namespaces = collection.Builder{
		Name:         "k8s/core/v1/namespaces",
		VariableName: "K8SCoreV1Namespaces",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "Namespace",
			Plural:        "namespaces",
			Version:       "v1",
			Proto:         "k8s.io.api.core.v1.NamespaceSpec",
			ProtoPackage:  "k8s.io/api/core/v1",
			ClusterScoped: true,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SCoreV1Nodes describes the collection k8s/core/v1/nodes
	K8SCoreV1Nodes = collection.Builder{
		Name:         "k8s/core/v1/nodes",
		VariableName: "K8SCoreV1Nodes",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "Node",
			Plural:        "nodes",
			Version:       "v1",
			Proto:         "k8s.io.api.core.v1.NodeSpec",
			ProtoPackage:  "k8s.io/api/core/v1",
			ClusterScoped: true,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SCoreV1Pods describes the collection k8s/core/v1/pods
	K8SCoreV1Pods = collection.Builder{
		Name:         "k8s/core/v1/pods",
		VariableName: "K8SCoreV1Pods",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "Pod",
			Plural:        "pods",
			Version:       "v1",
			Proto:         "k8s.io.api.core.v1.Pod",
			ProtoPackage:  "k8s.io/api/core/v1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SCoreV1Secrets describes the collection k8s/core/v1/secrets
	K8SCoreV1Secrets = collection.Builder{
		Name:         "k8s/core/v1/secrets",
		VariableName: "K8SCoreV1Secrets",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "Secret",
			Plural:        "secrets",
			Version:       "v1",
			Proto:         "k8s.io.api.core.v1.Secret",
			ProtoPackage:  "k8s.io/api/core/v1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SCoreV1Services describes the collection k8s/core/v1/services
	K8SCoreV1Services = collection.Builder{
		Name:         "k8s/core/v1/services",
		VariableName: "K8SCoreV1Services",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "",
			Kind:          "Service",
			Plural:        "services",
			Version:       "v1",
			Proto:         "k8s.io.api.core.v1.ServiceSpec",
			ProtoPackage:  "k8s.io/api/core/v1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SExtensionsV1Beta1Ingresses describes the collection
	// k8s/extensions/v1beta1/ingresses
	K8SExtensionsV1Beta1Ingresses = collection.Builder{
		Name:         "k8s/extensions/v1beta1/ingresses",
		VariableName: "K8SExtensionsV1Beta1Ingresses",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "extensions",
			Kind:          "Ingress",
			Plural:        "ingresses",
			Version:       "v1beta1",
			Proto:         "k8s.io.api.extensions.v1beta1.IngressSpec",
			ProtoPackage:  "k8s.io/api/extensions/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SNetworkingIstioIoV1Alpha3Destinationrules describes the collection
	// k8s/networking.istio.io/v1alpha3/destinationrules
	K8SNetworkingIstioIoV1Alpha3Destinationrules = collection.Builder{
		Name:         "k8s/networking.istio.io/v1alpha3/destinationrules",
		VariableName: "K8SNetworkingIstioIoV1Alpha3Destinationrules",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "DestinationRule",
			Plural:        "destinationrules",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.DestinationRule",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SNetworkingIstioIoV1Alpha3Envoyfilters describes the collection
	// k8s/networking.istio.io/v1alpha3/envoyfilters
	K8SNetworkingIstioIoV1Alpha3Envoyfilters = collection.Builder{
		Name:         "k8s/networking.istio.io/v1alpha3/envoyfilters",
		VariableName: "K8SNetworkingIstioIoV1Alpha3Envoyfilters",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "EnvoyFilter",
			Plural:        "envoyfilters",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.EnvoyFilter",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SNetworkingIstioIoV1Alpha3Gateways describes the collection
	// k8s/networking.istio.io/v1alpha3/gateways
	K8SNetworkingIstioIoV1Alpha3Gateways = collection.Builder{
		Name:         "k8s/networking.istio.io/v1alpha3/gateways",
		VariableName: "K8SNetworkingIstioIoV1Alpha3Gateways",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "Gateway",
			Plural:        "gateways",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.Gateway",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SNetworkingIstioIoV1Alpha3Serviceentries describes the collection
	// k8s/networking.istio.io/v1alpha3/serviceentries
	K8SNetworkingIstioIoV1Alpha3Serviceentries = collection.Builder{
		Name:         "k8s/networking.istio.io/v1alpha3/serviceentries",
		VariableName: "K8SNetworkingIstioIoV1Alpha3Serviceentries",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "ServiceEntry",
			Plural:        "serviceentries",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.ServiceEntry",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SNetworkingIstioIoV1Alpha3Sidecars describes the collection
	// k8s/networking.istio.io/v1alpha3/sidecars
	K8SNetworkingIstioIoV1Alpha3Sidecars = collection.Builder{
		Name:         "k8s/networking.istio.io/v1alpha3/sidecars",
		VariableName: "K8SNetworkingIstioIoV1Alpha3Sidecars",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "Sidecar",
			Plural:        "sidecars",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.Sidecar",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SNetworkingIstioIoV1Alpha3Virtualservices describes the collection
	// k8s/networking.istio.io/v1alpha3/virtualservices
	K8SNetworkingIstioIoV1Alpha3Virtualservices = collection.Builder{
		Name:         "k8s/networking.istio.io/v1alpha3/virtualservices",
		VariableName: "K8SNetworkingIstioIoV1Alpha3Virtualservices",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "VirtualService",
			Plural:        "virtualservices",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.VirtualService",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SNetworkingIstioIoV1Alpha3Workloadentries describes the collection
	// k8s/networking.istio.io/v1alpha3/workloadentries
	K8SNetworkingIstioIoV1Alpha3Workloadentries = collection.Builder{
		Name:         "k8s/networking.istio.io/v1alpha3/workloadentries",
		VariableName: "K8SNetworkingIstioIoV1Alpha3Workloadentries",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.istio.io",
			Kind:          "WorkloadEntry",
			Plural:        "workloadentries",
			Version:       "v1alpha3",
			Proto:         "istio.networking.v1alpha3.WorkloadEntry",
			ProtoPackage:  "istio.io/api/networking/v1alpha3",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SSecurityIstioIoV1Beta1Authorizationpolicies describes the collection
	// k8s/security.istio.io/v1beta1/authorizationpolicies
	K8SSecurityIstioIoV1Beta1Authorizationpolicies = collection.Builder{
		Name:         "k8s/security.istio.io/v1beta1/authorizationpolicies",
		VariableName: "K8SSecurityIstioIoV1Beta1Authorizationpolicies",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "security.istio.io",
			Kind:          "AuthorizationPolicy",
			Plural:        "authorizationpolicies",
			Version:       "v1beta1",
			Proto:         "istio.security.v1beta1.AuthorizationPolicy",
			ProtoPackage:  "istio.io/api/security/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SSecurityIstioIoV1Beta1Peerauthentications describes the collection
	// k8s/security.istio.io/v1beta1/peerauthentications
	K8SSecurityIstioIoV1Beta1Peerauthentications = collection.Builder{
		Name:         "k8s/security.istio.io/v1beta1/peerauthentications",
		VariableName: "K8SSecurityIstioIoV1Beta1Peerauthentications",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "security.istio.io",
			Kind:          "PeerAuthentication",
			Plural:        "peerauthentications",
			Version:       "v1beta1",
			Proto:         "istio.security.v1beta1.PeerAuthentication",
			ProtoPackage:  "istio.io/api/security/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SSecurityIstioIoV1Beta1Requestauthentications describes the
	// collection k8s/security.istio.io/v1beta1/requestauthentications
	K8SSecurityIstioIoV1Beta1Requestauthentications = collection.Builder{
		Name:         "k8s/security.istio.io/v1beta1/requestauthentications",
		VariableName: "K8SSecurityIstioIoV1Beta1Requestauthentications",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "security.istio.io",
			Kind:          "RequestAuthentication",
			Plural:        "requestauthentications",
			Version:       "v1beta1",
			Proto:         "istio.security.v1beta1.RequestAuthentication",
			ProtoPackage:  "istio.io/api/security/v1beta1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SServiceApisV1Alpha1Gatewayclasses describes the collection
	// k8s/service_apis/v1alpha1/gatewayclasses
	K8SServiceApisV1Alpha1Gatewayclasses = collection.Builder{
		Name:         "k8s/service_apis/v1alpha1/gatewayclasses",
		VariableName: "K8SServiceApisV1Alpha1Gatewayclasses",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.x.k8s.io",
			Kind:          "GatewayClass",
			Plural:        "gatewayclasses",
			Version:       "v1alpha1",
			Proto:         "k8s.io.service_apis.api.v1alpha1.GatewayClassSpec",
			ProtoPackage:  "sigs.k8s.io/service-apis/api/v1alpha1",
			ClusterScoped: true,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SServiceApisV1Alpha1Gateways describes the collection
	// k8s/service_apis/v1alpha1/gateways
	K8SServiceApisV1Alpha1Gateways = collection.Builder{
		Name:         "k8s/service_apis/v1alpha1/gateways",
		VariableName: "K8SServiceApisV1Alpha1Gateways",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.x.k8s.io",
			Kind:          "Gateway",
			Plural:        "gateways",
			Version:       "v1alpha1",
			Proto:         "k8s.io.service_apis.api.v1alpha1.GatewaySpec",
			ProtoPackage:  "sigs.k8s.io/service-apis/api/v1alpha1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SServiceApisV1Alpha1Httproutes describes the collection
	// k8s/service_apis/v1alpha1/httproutes
	K8SServiceApisV1Alpha1Httproutes = collection.Builder{
		Name:         "k8s/service_apis/v1alpha1/httproutes",
		VariableName: "K8SServiceApisV1Alpha1Httproutes",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.x.k8s.io",
			Kind:          "HTTPRoute",
			Plural:        "httproutes",
			Version:       "v1alpha1",
			Proto:         "k8s.io.service_apis.api.v1alpha1.HTTPRouteSpec",
			ProtoPackage:  "sigs.k8s.io/service-apis/api/v1alpha1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SServiceApisV1Alpha1Tcproutes describes the collection
	// k8s/service_apis/v1alpha1/tcproutes
	K8SServiceApisV1Alpha1Tcproutes = collection.Builder{
		Name:         "k8s/service_apis/v1alpha1/tcproutes",
		VariableName: "K8SServiceApisV1Alpha1Tcproutes",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.x.k8s.io",
			Kind:          "TcpRoute",
			Plural:        "tcproutes",
			Version:       "v1alpha1",
			Proto:         "k8s.io.service_apis.api.v1alpha1.TcpRouteSpec",
			ProtoPackage:  "sigs.k8s.io/service-apis/api/v1alpha1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// K8SServiceApisV1Alpha1Trafficsplits describes the collection
	// k8s/service_apis/v1alpha1/trafficsplits
	K8SServiceApisV1Alpha1Trafficsplits = collection.Builder{
		Name:         "k8s/service_apis/v1alpha1/trafficsplits",
		VariableName: "K8SServiceApisV1Alpha1Trafficsplits",
		Disabled:     false,
		Resource: resource.Builder{
			Group:         "networking.x.k8s.io",
			Kind:          "TrafficSplit",
			Plural:        "trafficsplits",
			Version:       "v1alpha1",
			Proto:         "k8s.io.service_apis.api.v1alpha1.TrafficSplitSpec",
			ProtoPackage:  "sigs.k8s.io/service-apis/api/v1alpha1",
			ClusterScoped: false,
			ValidateProto: validation.EmptyValidate,
		}.MustBuild(),
	}.MustBuild()

	// All contains all collections in the system.
	All = collection.NewSchemasBuilder().
		MustAdd(IstioConfigV1Alpha2Adapters).
		MustAdd(IstioConfigV1Alpha2Templates).
		MustAdd(IstioMeshV1Alpha1MeshConfig).
		MustAdd(IstioMeshV1Alpha1MeshNetworks).
		MustAdd(IstioNetworkingV1Alpha3Destinationrules).
		MustAdd(IstioNetworkingV1Alpha3Envoyfilters).
		MustAdd(IstioNetworkingV1Alpha3Gateways).
		MustAdd(IstioNetworkingV1Alpha3Serviceentries).
		MustAdd(IstioNetworkingV1Alpha3Sidecars).
		MustAdd(IstioNetworkingV1Alpha3Virtualservices).
		MustAdd(IstioNetworkingV1Alpha3Workloadentries).
		MustAdd(IstioSecurityV1Beta1Authorizationpolicies).
		MustAdd(IstioSecurityV1Beta1Peerauthentications).
		MustAdd(IstioSecurityV1Beta1Requestauthentications).
		MustAdd(K8SApiextensionsK8SIoV1Beta1Customresourcedefinitions).
		MustAdd(K8SAppsV1Deployments).
		MustAdd(K8SConfigIstioIoV1Alpha2Adapters).
		MustAdd(K8SConfigIstioIoV1Alpha2Templates).
		MustAdd(K8SCoreV1Configmaps).
		MustAdd(K8SCoreV1Endpoints).
		MustAdd(K8SCoreV1Namespaces).
		MustAdd(K8SCoreV1Nodes).
		MustAdd(K8SCoreV1Pods).
		MustAdd(K8SCoreV1Secrets).
		MustAdd(K8SCoreV1Services).
		MustAdd(K8SExtensionsV1Beta1Ingresses).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Destinationrules).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Envoyfilters).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Gateways).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Serviceentries).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Sidecars).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Virtualservices).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Workloadentries).
		MustAdd(K8SSecurityIstioIoV1Beta1Authorizationpolicies).
		MustAdd(K8SSecurityIstioIoV1Beta1Peerauthentications).
		MustAdd(K8SSecurityIstioIoV1Beta1Requestauthentications).
		MustAdd(K8SServiceApisV1Alpha1Gatewayclasses).
		MustAdd(K8SServiceApisV1Alpha1Gateways).
		MustAdd(K8SServiceApisV1Alpha1Httproutes).
		MustAdd(K8SServiceApisV1Alpha1Tcproutes).
		MustAdd(K8SServiceApisV1Alpha1Trafficsplits).
		Build()

	// Istio contains only Istio collections.
	Istio = collection.NewSchemasBuilder().
		MustAdd(IstioConfigV1Alpha2Adapters).
		MustAdd(IstioConfigV1Alpha2Templates).
		MustAdd(IstioMeshV1Alpha1MeshConfig).
		MustAdd(IstioMeshV1Alpha1MeshNetworks).
		MustAdd(IstioNetworkingV1Alpha3Destinationrules).
		MustAdd(IstioNetworkingV1Alpha3Envoyfilters).
		MustAdd(IstioNetworkingV1Alpha3Gateways).
		MustAdd(IstioNetworkingV1Alpha3Serviceentries).
		MustAdd(IstioNetworkingV1Alpha3Sidecars).
		MustAdd(IstioNetworkingV1Alpha3Virtualservices).
		MustAdd(IstioNetworkingV1Alpha3Workloadentries).
		MustAdd(IstioSecurityV1Beta1Authorizationpolicies).
		MustAdd(IstioSecurityV1Beta1Peerauthentications).
		MustAdd(IstioSecurityV1Beta1Requestauthentications).
		Build()

	// Kube contains only kubernetes collections.
	Kube = collection.NewSchemasBuilder().
		MustAdd(K8SApiextensionsK8SIoV1Beta1Customresourcedefinitions).
		MustAdd(K8SAppsV1Deployments).
		MustAdd(K8SConfigIstioIoV1Alpha2Adapters).
		MustAdd(K8SConfigIstioIoV1Alpha2Templates).
		MustAdd(K8SCoreV1Configmaps).
		MustAdd(K8SCoreV1Endpoints).
		MustAdd(K8SCoreV1Namespaces).
		MustAdd(K8SCoreV1Nodes).
		MustAdd(K8SCoreV1Pods).
		MustAdd(K8SCoreV1Secrets).
		MustAdd(K8SCoreV1Services).
		MustAdd(K8SExtensionsV1Beta1Ingresses).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Destinationrules).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Envoyfilters).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Gateways).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Serviceentries).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Sidecars).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Virtualservices).
		MustAdd(K8SNetworkingIstioIoV1Alpha3Workloadentries).
		MustAdd(K8SSecurityIstioIoV1Beta1Authorizationpolicies).
		MustAdd(K8SSecurityIstioIoV1Beta1Peerauthentications).
		MustAdd(K8SSecurityIstioIoV1Beta1Requestauthentications).
		MustAdd(K8SServiceApisV1Alpha1Gatewayclasses).
		MustAdd(K8SServiceApisV1Alpha1Gateways).
		MustAdd(K8SServiceApisV1Alpha1Httproutes).
		MustAdd(K8SServiceApisV1Alpha1Tcproutes).
		MustAdd(K8SServiceApisV1Alpha1Trafficsplits).
		Build()

	// Pilot contains only collections used by Pilot.
	Pilot = collection.NewSchemasBuilder().
		MustAdd(IstioNetworkingV1Alpha3Destinationrules).
		MustAdd(IstioNetworkingV1Alpha3Envoyfilters).
		MustAdd(IstioNetworkingV1Alpha3Gateways).
		MustAdd(IstioNetworkingV1Alpha3Serviceentries).
		MustAdd(IstioNetworkingV1Alpha3Sidecars).
		MustAdd(IstioNetworkingV1Alpha3Virtualservices).
		MustAdd(IstioNetworkingV1Alpha3Workloadentries).
		MustAdd(IstioSecurityV1Beta1Authorizationpolicies).
		MustAdd(IstioSecurityV1Beta1Peerauthentications).
		MustAdd(IstioSecurityV1Beta1Requestauthentications).
		Build()

	// PilotServiceApi contains only collections used by Pilot, including experimental Service Api.
	PilotServiceApi = collection.NewSchemasBuilder().
			MustAdd(IstioNetworkingV1Alpha3Destinationrules).
			MustAdd(IstioNetworkingV1Alpha3Envoyfilters).
			MustAdd(IstioNetworkingV1Alpha3Gateways).
			MustAdd(IstioNetworkingV1Alpha3Serviceentries).
			MustAdd(IstioNetworkingV1Alpha3Sidecars).
			MustAdd(IstioNetworkingV1Alpha3Virtualservices).
			MustAdd(IstioNetworkingV1Alpha3Workloadentries).
			MustAdd(IstioSecurityV1Beta1Authorizationpolicies).
			MustAdd(IstioSecurityV1Beta1Peerauthentications).
			MustAdd(IstioSecurityV1Beta1Requestauthentications).
			MustAdd(K8SServiceApisV1Alpha1Gatewayclasses).
			MustAdd(K8SServiceApisV1Alpha1Gateways).
			MustAdd(K8SServiceApisV1Alpha1Httproutes).
			MustAdd(K8SServiceApisV1Alpha1Tcproutes).
			MustAdd(K8SServiceApisV1Alpha1Trafficsplits).
			Build()
)
