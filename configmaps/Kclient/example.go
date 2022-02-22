//http://blog.johandry.com/post/build-k8s-client/

package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	//kubeconfigPath := "/Users/sdillikar/.config/tanzu/tkg/standalone/tce/kube.conf"
	kubeconfigPath := "/Users/sdillikar/kind-basic.yaml"
	namespaceName := "game"
	configmapName := "game-data"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	configMapData := make(map[string]string)
	uiProperties := `
    color.good=purple
    color.bad=yellow
    color.normal=black
    allow.textmode=true
  `
	configMapData["ui.properties"] = uiProperties

	configMap := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      configmapName,
			Namespace: namespaceName,
		},
		Data: configMapData,
	}

	configmapClient := clientset.CoreV1().ConfigMaps(namespaceName)

	var cm *corev1.ConfigMap

	_, err = configmapClient.Get(context.TODO(), configmapName, metav1.GetOptions{})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println("Creating configmap...")
		cm, _ = configmapClient.Create(context.TODO(), &configMap, metav1.CreateOptions{})
	} else {
		fmt.Println("Updating configmap...")
		cm, _ = configmapClient.Update(context.TODO(), &configMap, metav1.UpdateOptions{})
	}
	fmt.Println("cm.Data:", cm.Data)
}
