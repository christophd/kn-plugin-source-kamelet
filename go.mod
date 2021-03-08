module knative.dev/kn-plugin-source-kamelet

go 1.15

require (
	github.com/apache/camel-k v1.3.1
	github.com/apache/camel-k/pkg/apis/camel v1.3.1
	github.com/apache/camel-k/pkg/client/camel v1.3.1
	github.com/spf13/cobra v1.1.3
	gotest.tools v2.2.0+incompatible
	k8s.io/apimachinery v0.18.12
	k8s.io/client-go v12.0.0+incompatible
	knative.dev/client v0.20.0
	knative.dev/hack v0.0.0-20210120165453-8d623a0af457
)
