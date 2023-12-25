package collections

import "istio.io/libistio/pkg/config/schema/collection"

var (
	LegacyDefault = collection.NewSchemasBuilder().
			MustAdd(MeshConfig).
			MustAdd(DestinationRule).
			MustAdd(EnvoyFilter).
			MustAdd(Gateway).
			MustAdd(ServiceEntry).
			MustAdd(Sidecar).
			MustAdd(VirtualService).
			MustAdd(WorkloadEntry).
			MustAdd(AuthorizationPolicy).
			MustAdd(PeerAuthentication).
			MustAdd(RequestAuthentication).
			MustAdd(Namespace).
			MustAdd(Service).
			Build()

	LegacyLocalAnalysis = collection.NewSchemasBuilder().
				MustAdd(MeshConfig).
				MustAdd(MeshNetworks).
				MustAdd(DestinationRule).
				MustAdd(EnvoyFilter).
				MustAdd(Gateway).
				MustAdd(ServiceEntry).
				MustAdd(Sidecar).
				MustAdd(VirtualService).
				MustAdd(CustomResourceDefinition).
				MustAdd(Deployment).
				MustAdd(ConfigMap).
				MustAdd(Namespace).
				MustAdd(Pod).
				MustAdd(Secret).
				MustAdd(Service).
				Build()

	LegacyDefaultExcludeKubeResourceKinds = func() []string {
		return []string{
			Endpoints.Kind(),
			Namespace.Kind(),
			Node.Kind(),
			Pod.Kind(),
			Service.Kind(),
		}
	}
)
