module github.com/banzaicloud/kafka-operator

go 1.16

require (
	emperror.dev/errors v0.8.0
	github.com/Shopify/sarama v1.29.0
	github.com/banzaicloud/bank-vaults/pkg/sdk v0.3.1
	github.com/banzaicloud/istio-client-go v0.0.9
	github.com/banzaicloud/istio-operator v0.0.0-20210603082335-fd31d6ff3e0d
	github.com/banzaicloud/k8s-objectmatcher v1.4.1
	github.com/banzaicloud/kafka-operator/api v0.0.0
	github.com/banzaicloud/kafka-operator/properties v0.0.0
	github.com/envoyproxy/go-control-plane v0.9.9
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-logr/logr v0.3.0
	github.com/golang/protobuf v1.4.3
	github.com/hashicorp/go-hclog v0.12.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.7 // indirect
	github.com/hashicorp/vault v1.4.0
	github.com/hashicorp/vault/api v1.0.5-0.20200317185738-82f498082f02
	github.com/hashicorp/vault/sdk v0.1.14-0.20200406173424-43a93d4a43b1
	github.com/imdario/mergo v0.3.12
	github.com/influxdata/influxdb v1.7.6 // indirect
	github.com/jetstack/cert-manager v1.3.1
	github.com/lestrrat-go/backoff v1.0.0
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-colorable v0.1.7 // indirect
	github.com/onsi/ginkgo v1.16.1
	github.com/onsi/gomega v1.11.0
	github.com/pavel-v-chernykh/keystore-go v2.1.0+incompatible
	github.com/prometheus/common v0.10.0
	github.com/shirou/gopsutil v3.20.12+incompatible // indirect
	github.com/spf13/cast v1.3.1 // indirect
	go.uber.org/zap v1.15.0
	google.golang.org/protobuf v1.25.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.20.2
	k8s.io/apiextensions-apiserver v0.20.1
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	sigs.k8s.io/controller-runtime v0.8.3
)

replace (
	github.com/banzaicloud/kafka-operator/api => ./pkg/sdk
	github.com/banzaicloud/kafka-operator/properties => ./properties
)
