module istio.io/libistio

go 1.20

// Client-go does not handle different versions of mergo due to some breaking changes - use the matching version
replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

require (
	github.com/cncf/xds/go v0.0.0-20230607035331-e9ce68804cb4
	github.com/envoyproxy/go-control-plane v0.11.2-0.20230725211550-11bfe846bcd4
	github.com/fsnotify/fsnotify v1.7.0
	github.com/go-logr/logr v1.3.0
	github.com/gogo/protobuf v1.3.2
	github.com/golang/protobuf v1.5.3
	github.com/google/cel-go v0.18.2
	github.com/google/go-cmp v0.6.0
	github.com/hashicorp/go-multierror v1.0.0
	github.com/lestrrat-go/jwx v1.2.27
	github.com/prometheus/client_golang v1.17.0
	github.com/spf13/cobra v1.7.0
	github.com/spf13/pflag v1.0.5
	go.opentelemetry.io/otel v1.17.0
	go.opentelemetry.io/otel/exporters/prometheus v0.39.1-0.20230714155235-03b8c47770f2
	go.opentelemetry.io/otel/metric v1.17.0
	go.opentelemetry.io/otel/sdk/metric v0.39.0
	go.uber.org/zap v1.26.0
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
	google.golang.org/grpc v1.58.3
	google.golang.org/protobuf v1.31.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v3 v3.0.1
	istio.io/api v1.19.1
	istio.io/client-go v1.19.1
	k8s.io/api v0.28.3
	k8s.io/apiextensions-apiserver v0.28.3
	k8s.io/apimachinery v0.28.3
	k8s.io/client-go v0.28.3
	k8s.io/klog/v2 v2.110.1
	sigs.k8s.io/gateway-api v1.0.0
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/antlr4-go/antlr/v4 v4.13.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.0.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/lestrrat-go/backoff/v2 v2.0.8 // indirect
	github.com/lestrrat-go/blackmagic v1.0.2 // indirect
	github.com/lestrrat-go/httpcc v1.0.1 // indirect
	github.com/lestrrat-go/iter v1.0.2 // indirect
	github.com/lestrrat-go/option v1.0.1 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	go.opentelemetry.io/otel/sdk v1.16.0 // indirect
	go.opentelemetry.io/otel/trace v1.17.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/goleak v1.2.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.16.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231009173412-8bfb1ae86b6c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
)

replace golang.org/x/exp => golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1
