package main

import (
	"encoding/json"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/serializer/yaml"
)

const dsManifest = `
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: my-example-daemon-set
  namespace: default
spec:
  selector:
    matchLabels:
      name: nginx-ds
  template:
    metadata:
      labels:
        name: nginx-ds
    spec:
      containers:
      - name: nginx
        image: nginx:latest
`

func main() {
	obj := &unstructured.Unstructured{}

	// decode YAML into unstructured.Unstructured
	dec := yaml.NewDecodingSerializer(unstructured.UnstructuredJSONScheme)
	_, gvk, err := dec.Decode([]byte(dsManifest), nil, obj)

	if err != nil {
		fmt.Println("received error %s", err)
	}
	// Get the common metadata, and show GVK
	fmt.Println(obj.GetName(), gvk.String())

	// encode back to JSON
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(obj)
}
