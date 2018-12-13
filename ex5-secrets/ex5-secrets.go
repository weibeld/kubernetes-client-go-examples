package main

import (
	"bufio"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	coreV1Types "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

// API client for managing secrets
var secretsClient coreV1Types.SecretInterface

func initClient() {
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secretsClient = clientset.CoreV1().Secrets("default")
}

func main() {

	initClient()

	//-------------------------------------------------------------------------//
	// Read secret specification from YAML file
	//-------------------------------------------------------------------------//

	bytes, err := ioutil.ReadFile("secret.yml")
	if err != nil {
		panic(err.Error())
	}

	var secretSpec coreV1.Secret
	err = yaml.Unmarshal(bytes, &secretSpec)
	if err != nil {
		panic(err.Error())
	}

	secretName := secretSpec.ObjectMeta.Name
	var secret *coreV1.Secret

	//-------------------------------------------------------------------------//
	// Create secret
	//-------------------------------------------------------------------------//

	_, err = secretsClient.Create(&secretSpec)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created secret %s\n", secretName)

	prompt()

	//-------------------------------------------------------------------------//
	// Read the secret and print its data
	//-------------------------------------------------------------------------//

	secret = readAndPrintSecret(secretName)
	prompt()

	//-------------------------------------------------------------------------//
	// Add new secret key/value pair to 'data' (as a byte array)
	//-------------------------------------------------------------------------//

	// Evaluates to true if "data" AND "stringData" field in YAML file are empty
	if secret.Data == nil {
		secret.Data = map[string][]byte{}
	}

	fmt.Println("Adding new key/value pair to secret as a byte array (Data)")
	secret.Data["newKey1"] = []byte("newValue1")

	_, err = secretsClient.Update(secret)
	if err != nil {
		panic(err.Error())
	}

	secret = readAndPrintSecret(secretName)
	prompt()

	//-------------------------------------------------------------------------//
	// Add new secret key/value pair to 'stringData' (as string)
	//-------------------------------------------------------------------------//

	// Evaluates to true if "stringData" field in YAML file is empty
	if secret.StringData == nil {
		secret.StringData = map[string]string{}
	}

	fmt.Println("Adding new key/value pair to secret as a string (StringData)")
	secret.StringData["newKey2"] = "newValue2"

	_, err = secretsClient.Update(secret)
	if err != nil {
		panic(err.Error())
	}

	readAndPrintSecret(secretName)
	prompt()

	//-------------------------------------------------------------------------//
	// Delete secret
	//-------------------------------------------------------------------------//

	err = secretsClient.Delete(secretName, &metaV1.DeleteOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deleted secret %s\n", secretName)

}

// Read the specified secret and print its data
func readAndPrintSecret(name string) *coreV1.Secret {
	fmt.Printf("Reading secret %s\n", name)

	secret, err := secretsClient.Get(name, metaV1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Data:")
	// When reading a secret, only Data contains any data, StringData is empty
	for key, value := range secret.Data {
		// key is string, value is []byte
		fmt.Printf("    %s: %s\n", key, value)
	}
	return secret
}

// Block until the users presses the Enter key
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
