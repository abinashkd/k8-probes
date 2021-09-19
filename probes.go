package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"

	"sigs.k8s.io/yaml"
)

func probes() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pod, _ := clientset.CoreV1().Pods("abinashkd-dev").List(context.TODO(), metav1.ListOptions{})

	for i := 0; i < len(pod.Items); i++ {

		for j := 0; j < len(pod.Items[i].Spec.Containers); j++ {
			fmt.Println(pod.Items[i].Spec.Containers[j].Image)
			fmt.Println(pod.Items[i].Spec.Containers[j].LivenessProbe)
			newL, _ := json.MarshalIndent(pod.Items[i].Spec.Containers[j].LivenessProbe, "", "  ")
			fmt.Println(string(newL))

			yl, _ := yaml.JSONToYAML(newL)
			fmt.Println(string(yl))
			fmt.Println(pod.Items[i].Spec.Containers[j].ReadinessProbe)
			fmt.Println(pod.Items[i].Spec.Containers[j].StartupProbe)

		}

	}

}
