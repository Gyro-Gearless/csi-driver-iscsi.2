module github.com/kubernetes-csi/csi-driver-iscsi

go 1.16

require (
	github.com/container-storage-interface/spec v1.2.0
	github.com/kubernetes-csi/csi-lib-iscsi v0.0.0-20190415173011-c545557492f4
	github.com/kubernetes-csi/csi-lib-utils v0.2.0
	github.com/onsi/gomega v1.8.1 // indirect
	github.com/spf13/cobra v1.0.0
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	google.golang.org/grpc v1.27.0
	k8s.io/klog/v2 v2.2.0
	k8s.io/kubernetes v1.19.4
	k8s.io/utils v0.0.0-20200729134348-d5654de09c73
)

replace (
	k8s.io/api => k8s.io/api v0.19.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.19.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.4
	k8s.io/apiserver => k8s.io/apiserver v0.19.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.19.4
	k8s.io/client-go => k8s.io/client-go v0.19.4
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.19.4
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.19.4
	k8s.io/code-generator => k8s.io/code-generator v0.19.4
	k8s.io/component-base => k8s.io/component-base v0.19.4
	k8s.io/cri-api => k8s.io/cri-api v0.19.4
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.19.4
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.19.4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.19.4
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.19.4
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.19.4
	k8s.io/kubectl => k8s.io/kubectl v0.19.4
	k8s.io/kubelet => k8s.io/kubelet v0.19.4
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.19.4
	k8s.io/metrics => k8s.io/metrics v0.19.4
	k8s.io/node-api => k8s.io/node-api v0.19.4
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.19.4
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.19.4
	k8s.io/sample-controller => k8s.io/sample-controller v0.19.4
)