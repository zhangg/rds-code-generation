package main

import (
	"flag"
	"fmt"


	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	rdsclientset "github.com/rds-operator/pkg/generated/clientset/versioned"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %v", err)
	}

	rdsClient, err := rdsclientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %v", err)
	}

	list, err := rdsClient.RdsV1alpha1().Mysqls("default").List(metav1.ListOptions{})
	if err != nil {
		klog.Fatalf("Error listing all databases: %v", err)
	}

	for _, db := range list.Items {
		fmt.Printf("database %s with replicas %q\n", db.Name, db.Spec.Replicas)
	}
}
