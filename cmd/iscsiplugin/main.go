/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"os"

	iscsi "github.com/kubernetes-csi/csi-driver-iscsi/pkg/iscsi"
	iscsiLib "github.com/kubernetes-csi/csi-lib-iscsi/iscsi"
	klog "k8s.io/klog/v2"
)

var (
	endpoint = flag.String("endpoint", "unix:///csi/csi.sock", "CSI endpoint")
	nodeID   = flag.String("nodeid", "", "node id")
)

func init() {
	_ = flag.Set("logtostderr", "true")
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	iscsiLib.EnableDebugLogging(os.Stdout)
	klog.Infof("+++ HERE WE GO _______.oOo._\\|/_.oOo._______________ %d", 6*7)
	handle()
	os.Exit(0)
}

func handle() {
	d := iscsi.NewDriver(*nodeID, *endpoint)
	d.Run()
}
