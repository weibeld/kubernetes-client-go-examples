# Kubernetes Go Client Examples

Example usages of the official Kubernetes Go client ([client-go](https://github.com/kubernetes/client-go)).

The example programs are adapted from the [examples](https://github.com/kubernetes/client-go/tree/master/examples) in the client-go GitHub repository.

## Usage

Compile and run the example programs as follows:

~~~bash
go build example.go
./example
~~~

### *kubeconfig* File

One of the first thing every program using a [Kubernetes client library](https://kubernetes.io/docs/reference/using-api/client-libraries/) does is reading a [*kubeconfig*](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/) file (the same type of file read by `kubectl`).

All example programs in this repository read the default *kubeconfig* file `~/.kube/config`. So, before running the programs, make sure that this file exists ad that the context is set to the Kubernetes cluster that you want to work on.

You can check the active context in the *kubeconfig* file like this:

~~~bash
kubectl config current-context
~~~

### Go Client Library

To be able to build the example programs, you first need to install the Go client library according to the instructions [here](https://github.com/kubernetes/client-go/blob/master/INSTALL.md#installing-client-go).

## Examples

### ex1-list

This example shows how to list different types of resources across all namespaces. The listed resources include:

- [Nodes](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.12/#node-v1-core)
- [Pods](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.12/#pod-v1-core)
- [Deployments](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.12/#deployment-v1-apps)

### ex2-deployment

This example shows how to create a new [deployment](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.12/#deployment-v1-apps) in the default namespace, list all deployments in the default namespace, and finally delete the created deployment.

The carried out steps are:

1. Create deployment
2. List deployments
3. Delete deployment

The execution of the program stops after each steps until you press *Enter*. This allows you to inspect the resources in the cluster with `kubectl` after each step.

### ex3-deployment-and-service

This example shows how to create a [deployment](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.12/#deployment-v1-apps) and a [service](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.12/#service-v1-core) that exposes this deployment. All resources are created in the default namespace.

The service is of type *load balancer* and listens on port 80. The deployment manages multiple replicas of [NGINX](https://hub.docker.com/_/nginx/) containers.

The steps carried out are:

1. Create deployment
2. Create service
3. Delete service
4. Delete deployment

The execution of the program stops after each steps until you press *Enter*. This allows you to inspect the resources in the cluster with `kubectl` after each step.

Note that after step 2, it might take some minutes until the service is accessible through the load balander IP address or DNS name (this depends on the cloud provider). Once the load balancer has been initialised, you should see the NGINX welcome page when you access the service.

