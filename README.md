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

- Nodes
- Pods
- Deployments

### ex2-deployment

This example shows how to create a new deployment in the default namespace, list all deployments in the default namespace, and finally delete the created deployment.
