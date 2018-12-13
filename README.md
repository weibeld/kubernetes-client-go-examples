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

### ex1-list-resources

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

### ex3-deployment-service

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

### ex5-secrets

This example shows how to create and update [secrets](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#secret-v1-core). [Kubernetes secrets](https://kubernetes.io/docs/concepts/configuration/secret/) are similar to [config maps](https://kubernetes.io/docs/concepts/configuration/secret/) except that secrets are intended to contain sensitive data (and are stored in encrypted form in the Kubernetes control plane). The purpose of secrets is to securely [distribute sensitive data to pods](https://kubernetes.io/docs/tasks/inject-data-application/distribute-credentials-secure/).

A single secret objects contains a `data` field that contains any number of key/value pairs. The keys are strings and the values are [Base64](https://tools.ietf.org/html/rfc4648#section-4) encodings of arbitrary bytes. The decoded values may be strings or non-strings (e.g. binary data).

There is a small complication: the secret object also provides a `stringData` field which allows to specify values as unencoded strings rather than Base64-encoded bytes. The values are then automatically Base64-encoded and the key/value pairs are saved in the `data` field of the object (overwriting any existing values with identical keys). The `stringData` field exists for convenience (to more easily specify secret data with string values) and it is **write-only**. That means, when reading a secret, only the `data` field contains any data, no matter whether the data has been specified through the `data` or `stringData` field.

When reading a secret specification from a YAML file, client-go (as well as `kubectl`) expect values under `data` to be Base64-encodings and values under `stringData` to be plain strings. However, when defining new secret data programmatically, client-go does the Base64 encoding automatically, and it simply requires values for `data` to be byte arrays and values for `stringData` to be strings. Similarly, when reading a secret, client-go returns the values of the `data` field as non-Base64-encoded byte arrays (regarding the `stringData` field, remember that it is never read). However, when you read a secret with `kubectl`, the `data` values are always output as Base64 encodings. 

All of this is summarised in the  following.

Secret specified in YAML file:

~~~
           | client-go | kubectl |
-----------+-----------+---------+
data       | Base64    | Base64  |
stringData | string    | string  |
~~~

Adding and reading secret values in client-go, and reading secret values with `kubectl`:

~~~
           |                client-go                |             |
           | add                | read               | kubectl get |
-----------+--------------------+--------------------+-------------|
data       | []byte (unencoded) | []byte (unencoded) | Base64      |
stringData | string             | -                  | -           |
~~~

To read the data of a secret with `kubectl`, use the following command:

~~~bash
kubectl get -o yaml secret/my-secret
~~~

The example reads an initial secret specification from a YAML file and then adds new key/value pairs to it, both through the `data` and the `stringData` fields. After each step the execution pauses until you press Enter. It is helpful to inspect the secret after each step with the above command.

