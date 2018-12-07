# Kubernetes Go Client Examples

Example usages of the official Kubernetes Go client ([client-go](https://github.com/kubernetes/client-go)).

The example programs are adapted from the [examples](https://github.com/kubernetes/client-go/tree/master/examples) in the client-go GitHub repository.

## Usage

Compile and run the example programs as follows:

~~~bash
go run example.go
~~~

### *kubeconfig* File

One of the first thing every program using a [Kubernetes client library](https://kubernetes.io/docs/reference/using-api/client-libraries/) does is reading a [*kubeconfig*](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/) file (the same type of file read by `kubectl`).

All example programs in this repository read the default *kubeconfig* file `~/.kube/config`. So, before running the programs, make sure that this file exists ad that the context is set to the Kubernetes cluster that you want to work on.

You can check the active context in the *kubeconfig* file like this:

~~~bash
kubectl config current-context
~~~

### Go Client Library

To be able to build the example programs, you first need to install the Go client library according to the instructions [here](https://github.com/kubernetes/client-go/blob/master/INSTALL.md#installing-client-go):

~~~bash
go get k8s.io/client-go/...
go get -u k8s.io/apimachinery/...
~~~

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

Note that after step 2, it might take some minutes until the service is accessible through the load balancer's IP address or DNS name (this depends on the cloud provider). Once the load balancer is initialised, you should see the NGINX welcome page when you access the service.

### ex4-read-yaml

This example shows how to read a Kubernetes API object (in this case a deployment) from a YAML file, and then deploy it to the cluster.

The client-go library provides no straightforward way to read API objects from YAML files (see [here](https://github.com/kubernetes/client-go/issues/193)). We can work around this by using the [`github.com/ghodss/yaml`](https://github.com/ghodss/yaml) package to parse a YAML file and convert it to a client-go Kubernetes API object struct.

The `github.com/ghodss/yaml` package is suitable, because it converts the YAML to JSON before reading it into a struct. This is good because the client-go API object structs contain [tags](https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e) for the JSON format, that is, the structs themselves define their serialisation/deserialisation to and from JSON. In this way, the overall conversion from YAML to a client-go API object struct works correctly, even if the structs have no YAML tags.

Note that if we defined the API objects in JSON instead of YAML, we could convert them to client-go API object structs with the standard [encoding/json](https://godoc.org/encoding/json) package (see [here](https://gist.github.com/mofelee/36b996d5c161dc60d551b52f3848a464)).

