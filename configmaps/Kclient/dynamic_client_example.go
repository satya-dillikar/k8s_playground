//http://blog.johandry.com/post/build-k8s-client/

package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//kubeconfigPath := "/Users/sdillikar/.config/tanzu/tkg/standalone/tce/kube.conf"
	kubeconfigPath := "/Users/sdillikar/kind-basic.yaml"
	namespaceName := "game"
	configmapName := "game-data"
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err)
	}
	dynamicClient, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		fmt.Printf("error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	configMapData := make(map[string]string)
	uiProperties := `
    color.good=purple
    color.bad=yellow
    color.normal=black
    allow.textmode=true
  `
	configMapData["ui.properties"] = uiProperties

	// configMap := corev1.ConfigMap{
	// 	TypeMeta: metav1.TypeMeta{
	// 		Kind:       "ConfigMap",
	// 		APIVersion: "v1",
	// 	},
	// 	ObjectMeta: metav1.ObjectMeta{
	// 		Name:      configmapName,
	// 		Namespace: namespaceName,
	// 	},
	// 	Data: configMapData,
	// }

	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "configmaps",
	}

	configMap := &unstructured.Unstructured{}
	configMap.SetUnstructuredContent(map[string]interface{}{
		"apiVersion": "v1",
		"kind":       "ConfigMap",
		"metadata": map[string]interface{}{
			"name":      configmapName,
			"namespace": namespaceName,
		},
		"Data": configMapData,
		// "spec": map[string]interface{}{
		// 	"Data": configMapData,
		// },
	})
	configmapClient := dynamicClient.Resource(gvr).Namespace(namespaceName)

	_, err = configmapClient.Get(context.TODO(), configmapName, metav1.GetOptions{})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Creating configmap...")
		cm, _ := configmapClient.Create(context.TODO(), configMap, metav1.CreateOptions{})
		fmt.Println("cm:", cm)
	} else {
		fmt.Println("Updating configmap...")
		cm, _ := configmapClient.Update(context.TODO(), configMap, metav1.UpdateOptions{})
		fmt.Println("cm:", cm)
	}
}
