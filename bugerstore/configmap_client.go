package main

import (
	"context"
	"fmt"
	"os"

	apiv1 "k8s.io/api/core/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListConfigMaps(clientset kubernetes.Interface, namespace string) ([]apiv1.ConfigMap, error) {
	ConfigMaps, err := clientset.CoreV1().
		ConfigMaps(namespace).
		List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error listing ConfigMaps : %v\n", err)
		os.Exit(1)
	}

	// kubectl get ConfigMaps -n <namespace> displays 3 columns:

	// | NAME                          | DATA  | AGE |
	// |-------------------------------|-------|-----|
	// | default-ns-rbac   						 |  0    | 14d |

	fmt.Println("Listing ConfigMaps")
	for _, ConfigMap := range ConfigMaps.Items {
		name := ConfigMap.Name
		data := ConfigMap.Data
		//diff := uint(time.Now().Sub(ConfigMap.StartTime.Time).Hours())
		age := 0

		fmt.Println(name, data, age)

	}
	return ConfigMaps.Items, nil
}

func GetConfigMap(clientset kubernetes.Interface, namespace string, ConfigMapName string) (*apiv1.ConfigMap, error) {

	// kubectl get ConfigMap <ConfigMap_name> -n <namespace> -o yaml

	ConfigMap, err := clientset.CoreV1().
		ConfigMaps(namespace).
		Get(context.TODO(), ConfigMapName, metav1.GetOptions{})

	if err != nil {
		fmt.Printf("error listing ConfigMaps : %v\n", err)
		os.Exit(1)
	}

	data := ConfigMap.Data
	//age := uint(time.Since(ConfigMap.Status.StartTime.Time).Hours())
	age := 0
	fmt.Println("Get Single ConfigMap object:")
	fmt.Println(ConfigMapName, data, age)

	return ConfigMap, err
}

func CreateConfigMap(clientset kubernetes.Interface, namespace string, ConfigMapName string) (*apiv1.ConfigMap, error) {

	configMapData := make(map[string]string)
	uiProperties := `
    color.good=purple
    color.bad=yellow
    color.normal=black
    allow.textmode=true
  `
	configMapData["ui.properties"] = uiProperties

	configMap := &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ConfigMapName,
			Namespace: namespace,
		},
		Data: configMapData,
	}

	cm, err := clientset.CoreV1().
		ConfigMaps(namespace).
		Create(context.TODO(), configMap, metav1.CreateOptions{})

	if err != nil {
		fmt.Printf("error creating ConfigMap : %v\n", err)
		os.Exit(1)
	}

	configmapname := cm.Name
	fmt.Println("Created Single ConfigMap object:", configmapname)

	return cm, err
}

func UpdateConfigMap(clientset kubernetes.Interface, namespace string, myconfigmap *apiv1.ConfigMap) (*apiv1.ConfigMap, error) {

	configMapData := make(map[string]string)
	uiProperties := `
	color.good=purple
	color.normal=white
	allow.textmode=false
`
	configMapData["ui.properties"] = uiProperties
	myconfigmap.Data = configMapData

	cm, err := clientset.CoreV1().
		ConfigMaps(namespace).
		Update(context.TODO(), myconfigmap, metav1.UpdateOptions{})

	if err != nil {
		fmt.Printf("error update ConfigMap : %v\n", err)
		os.Exit(1)
	}

	configmapname := cm.Name
	fmt.Println("Updated Single ConfigMap object:", configmapname)

	return cm, err
}

func DeleteConfigMap(clientset kubernetes.Interface, namespace string, ConfigMapName string) error {

	err := clientset.CoreV1().
		ConfigMaps(namespace).
		Delete(context.TODO(), ConfigMapName, metav1.DeleteOptions{})

	if err != nil {
		fmt.Printf("error Deleting ConfigMap : %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Delete Single ConfigMap object:", ConfigMapName)
	return err
}
