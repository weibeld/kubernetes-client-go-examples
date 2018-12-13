package main

import (
	"bufio"
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
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

	// Create a CoreV1Client (k8s.io/client-go/kubernetes/typed/core/v1)
	coreV1Client := clientset.CoreV1()
	// Create an AppsV1Client (k8s.io/client-go/kubernetes/typed/apps/v1)
	appsV1Client := clientset.AppsV1()

	//-------------------------------------------------------------------------//
	// Create a deployment (default namespace)
	//-------------------------------------------------------------------------//

	// Specification of the Deployment (k8s.io/api/apps/v1)
	deploymentSpec := &appsV1.Deployment{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsV1.DeploymentSpec{
			Replicas: func() *int32 { i := int32(2); return &i }(),
			Selector: &metaV1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: coreV1.PodTemplateSpec{
				ObjectMeta: metaV1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: coreV1.PodSpec{
					Containers: []coreV1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []coreV1.ContainerPort{
								{
									Name:          "http",
									Protocol:      coreV1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	deployment, err := appsV1Client.Deployments("default").Create(deploymentSpec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %s\n", deployment.ObjectMeta.Name)

	prompt()

	//-------------------------------------------------------------------------//
	// Create a load balancer service (default namespace)
	//-------------------------------------------------------------------------//

	// Specification of the Service (k8s.io/api/core/v1)
	serviceSpec := &coreV1.Service{
		ObjectMeta: metaV1.ObjectMeta{
			Name: "demo-service",
		},
		Spec: coreV1.ServiceSpec{
			Selector: map[string]string{
				"app": "demo",
			},
			Type: "LoadBalancer",
			Ports: []coreV1.ServicePort{
				{
					Port: 80,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 80,
					},
				},
			},
		},
	}

	service, err := coreV1Client.Services("default").Create(serviceSpec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created service %s\n", service.ObjectMeta.Name)

	prompt()

	//-------------------------------------------------------------------------//
	// Delete service
	//-------------------------------------------------------------------------//

	err = coreV1Client.Services("default").Delete(service.ObjectMeta.Name, &metaV1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted service %s\n", service.ObjectMeta.Name)

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
