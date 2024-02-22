module github.com/kubernetes-csi/csi-driver-iscsi

go 1.19

require (
	github.com/container-storage-interface/spec v1.8.0
	github.com/kubernetes-csi/csi-lib-utils v0.14.1
	golang.org/x/net v0.21.0
	google.golang.org/grpc v1.62.0
	k8s.io/klog/v2 v2.120.1
	k8s.io/kubernetes v1.26.11
	k8s.io/mount-utils v0.26.11 // indirect
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emicklei/go-restful/v3 v3.10.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-openapi/jsonpointer v0.19.6 // indirect
	github.com/go-openapi/jsonreference v0.20.1 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/gnostic v0.6.9 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/moby/sys/mountinfo v0.6.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/opencontainers/selinux v1.10.2 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.38.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/oauth2 v0.16.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/term v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/api v0.27.0 // indirect
	k8s.io/apimachinery v0.27.0 // indirect
	k8s.io/apiserver v0.26.11 // indirect
	k8s.io/client-go v0.27.0 // indirect
	k8s.io/cloud-provider v0.26.11 // indirect
	k8s.io/component-base v0.26.11 // indirect
	k8s.io/component-helpers v0.26.11 // indirect
	k8s.io/kube-openapi v0.0.0-20230308215209-15aac26d736a // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

require k8s.io/utils v0.0.0-20230209194617-a36077c30491

replace k8s.io/api => k8s.io/api v0.26.11

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.26.11

replace k8s.io/apimachinery => k8s.io/apimachinery v0.26.11

replace k8s.io/apiserver => k8s.io/apiserver v0.26.11

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.26.11

replace k8s.io/client-go => k8s.io/client-go v0.26.11

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.26.11

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.26.11

replace k8s.io/code-generator => k8s.io/code-generator v0.26.11

replace k8s.io/component-base => k8s.io/component-base v0.26.11

replace k8s.io/component-helpers => k8s.io/component-helpers v0.26.11

replace k8s.io/controller-manager => k8s.io/controller-manager v0.26.11

replace k8s.io/cri-api => k8s.io/cri-api v0.26.11

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.26.11

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.26.11

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.26.11

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.26.11

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.26.11

replace k8s.io/kubectl => k8s.io/kubectl v0.26.11

replace k8s.io/kubelet => k8s.io/kubelet v0.26.11

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.26.11

replace k8s.io/metrics => k8s.io/metrics v0.26.11

replace k8s.io/mount-utils => k8s.io/mount-utils v0.26.11

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.26.11

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.26.11

replace k8s.io/sample-controller => k8s.io/sample-controller v0.26.11

replace k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.26.11
