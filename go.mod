module knative.dev/kn-plugin-source-kamelet

go 1.15

require (
	github.com/apache/camel-k/pkg/apis/camel v1.3.1
	github.com/apache/camel-k/pkg/client/camel v1.3.1
	github.com/spf13/cobra v1.1.3
	gotest.tools v2.2.0+incompatible
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/client v0.21.0
	knative.dev/hack v0.0.0-20210203173706-8368e1f6eacf
)
