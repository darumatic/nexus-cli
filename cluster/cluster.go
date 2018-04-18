/*
Copyright 2016 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.
package cluster

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"regexp"
	"k8s.io/client-go/rest"
)

func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func ListImages(useKubeConfig bool) (map[string][]string, error) {

	var restConfig *rest.Config;

	if (useKubeConfig) {
		var kubeconfig *string

		if home := homeDir(); home != "" {
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
		restConfig = config;
	} else {

		config, err := rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}

		restConfig = config;
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		panic(err.Error())
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Gathering list of deployed pods")
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	fmt.Printf("Nexus registry path: %s\n", "localhost:5000")

	var containers = make(map[string][]string)

	for _, item := range pods.Items {
		for _, container := range item.Spec.Containers {
			if (strings.HasPrefix(container.Image, "localhost:5000")) {

				findRepoPattern := regexp.MustCompile(`(localhost:5000/)(.*):(.*)`)
				repo := findRepoPattern.FindStringSubmatch(container.Image);
				if (len((repo)) > 3) {
					if (!StringInSlice(repo[3], containers[repo[2]])) {
						containers[repo[2]] = append(containers[repo[2]], repo[3])
					}
				}
			}
		}
	}
	fmt.Println("List of deployed images from nexus registry")
	for image := range containers {
		fmt.Printf("Image: %s\n", image)
		for tag := range containers[image] {
			fmt.Printf("Tag: %s\n", containers[image][tag])
		}
	}
	return containers, nil;
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
