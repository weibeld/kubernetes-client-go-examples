package main

import (
	"bufio"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func main() {

	// Location of kubeconfig file
	kubeconfig := os.Getenv("HOME") + "/.kube/config"

	// Create a Config (k8s.io/client-go/rest)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create an API Clientset (k8s.io/client-go/kubernetes)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	appsV1Client := clientset.AppsV1()

	//-------------------------------------------------------------------------//
	// Read deployment object from YAML file
	//-------------------------------------------------------------------------//

	bytes, err := ioutil.ReadFile("deployment.yml")
	if err != nil {
		panic(err.Error())
	}

	var deploymentSpec appsV1.Deployment
	err = yaml.Unmarshal(bytes, &deploymentSpec)
	if err != nil {
		panic(err.Error())
	}

	//-------------------------------------------------------------------------//
	// Create deployment
	//-------------------------------------------------------------------------//

	deployment, err := appsV1Client.Deployments("default").Create(&deploymentSpec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %s\n", deployment.ObjectMeta.Name)

	prompt()

	//-------------------------------------------------------------------------//
	// Delete deployment
	//-------------------------------------------------------------------------//

	err = appsV1Client.Deployments("default").Delete(deployment.ObjectMeta.Name, &metaV1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted deployment %s\n", deployment.ObjectMeta.Name)

}

// prompt blocks until the users presses the Enter key
func prompt() {
	fmt.Printf("> Press Enter to continue ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
