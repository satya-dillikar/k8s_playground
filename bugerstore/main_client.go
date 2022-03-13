package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	//"k8s.io/client-go/informers"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"satya.com/burgerstore/configmaps"
	"satya.com/burgerstore/dyn_client"
	controller "satya.com/burgerstore/pkg/cmd"
	burgerStoreClient "satya.com/burgerstore/pkg/generated/clientset/versioned"
	informers "satya.com/burgerstore/pkg/generated/informers/externalversions"
	"satya.com/burgerstore/pods"
)

func NewK8sClient(masterUrl, kubeconfigPath string) (kubernetes.Interface, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
	if err != nil {
		fmt.Printf("error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	// create the clientset
	return kubernetes.NewForConfig(config)
}
func TestResourcePods(mykubeconfig string, namespace string) {

	myK8sClient, _ := NewK8sClient("", mykubeconfig)

	pods.ListPods(myK8sClient, namespace)
	podname := "testfoo"
	pod, _ := pods.CreatePod(myK8sClient, namespace, podname)
	_ = pod
	time.Sleep(5 * time.Second)
	pod, _ = pods.GetPod(myK8sClient, namespace, podname)
	time.Sleep(5 * time.Second)
	pod, _ = pods.UpdatePod(myK8sClient, namespace, pod)
	time.Sleep(5 * time.Second)
	pods.DeletePod(myK8sClient, namespace, pod.Name)
}

func TestResourceConfigMaps(mykubeconfig string, namespace string) {

	myK8sClient, _ := NewK8sClient("", mykubeconfig)

	configmaps.ListConfigMaps(myK8sClient, namespace)
	configmapname := "kube-root-ca.crt"
	configmaps.GetConfigMap(myK8sClient, namespace, configmapname)
	cmname := "testfoo-cm"
	cm, _ := configmaps.CreateConfigMap(myK8sClient, namespace, cmname)
	_ = cm
	time.Sleep(5 * time.Second)
	cm, _ = configmaps.GetConfigMap(myK8sClient, namespace, cmname)
	time.Sleep(5 * time.Second)
	cm, _ = configmaps.UpdateConfigMap(myK8sClient, namespace, cm)
	time.Sleep(5 * time.Second)
	configmaps.DeleteConfigMap(myK8sClient, namespace, cm.Name)
}

func NewClient(masterUrl, kubeconfigPath string) (*burgerStoreClient.Clientset, error) {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
	if err != nil {
		fmt.Printf("error creating dynamic client: %v\n", err)
		os.Exit(1)
	}
	// create the clientset
	return burgerStoreClient.NewForConfig(config)
}

func NewApiExtensionClient(masterUrl, kubeconfigPath string) apiextensionsclientset.Interface {
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
	if err != nil {
		fmt.Printf("error creating dynamic client: %v\n", err)
		os.Exit(1)
	}

	apiextensionsClient := apiextensionsclientset.NewForConfigOrDie(config)

	// create the clientset
	return apiextensionsClient
}

func TestResourceCRDs(mykubeconfig string, namespace string) {

	myK8sClient, _ := NewK8sClient("", mykubeconfig)
	myNewClient, _ := NewClient("", mykubeconfig)
	apiExtensionClient := NewApiExtensionClient("", mykubeconfig)

	exampleInformerFactory := informers.NewSharedInformerFactory(myNewClient, time.Second*30)

	stopCh := make(chan struct{})
	c := controller.NewController(myK8sClient, myNewClient, apiExtensionClient,
		exampleInformerFactory.Burgerstore().V1alpha1().BurgerStores())

	c.CreateCRD()

	CRDObjectObjectName := "testfoo-crd1"
	crd, err := controller.CreateCRDObject(myNewClient, namespace, CRDObjectObjectName)
	_ = crd
	if err != nil {
		klog.Infoln(err)
	}

	controller.ListCRDObject(myNewClient, namespace)
	crd, _ = controller.GetCRDObject(myNewClient, namespace, CRDObjectObjectName)
	time.Sleep(5 * time.Second)
	crd, _ = controller.UpdateCRDObject(myNewClient, namespace, crd)
	_ = crd
	time.Sleep(5 * time.Second)
	controller.DeleteCRDObject(myNewClient, namespace, CRDObjectObjectName)

	exampleInformerFactory.Start(stopCh)

	if err := c.Run(1, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func TestResourceCRDsDynamicClient(mykubeconfig string, namespace string) {

	myK8sClient, _ := NewK8sClient("", mykubeconfig)
	myNewClient, _ := NewClient("", mykubeconfig)
	apiExtensionClient := NewApiExtensionClient("", mykubeconfig)

	exampleInformerFactory := informers.NewSharedInformerFactory(myNewClient, time.Second*30)

	c := controller.NewController(myK8sClient, myNewClient, apiExtensionClient,
		exampleInformerFactory.Burgerstore().V1alpha1().BurgerStores())

	c.CreateCRD()

	CRDObjectObjectName := "testfoo-crd1"
	_, err := controller.CreateCRDObject(myNewClient, namespace, CRDObjectObjectName)

	if err != nil {
		klog.Infoln(err)
	}

	dyn_client.DynamicMain(mykubeconfig, namespace)
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("ERROR \n Usage: burgerstore <kubeconfig file path> <namespace>")
		return
	}
	kubeConfigPath := os.Args[1]
	namespace := os.Args[2]
	fmt.Printf("kubeConfigPath, %s\n", kubeConfigPath)
	fmt.Printf("namespace, %s\n", namespace)

	klog.SetOutput(os.Stderr)
	klog.SetOutput(os.Stdout)
	klog.InitFlags(nil) // initializing the flags
	defer klog.Flush()  // flushes all pending log I/O

	flag.Parse() // parses the command-line flags

	klog.Info("now you can see me")
	klog.V(4).Info("V4- now you can see me")
	fmt.Printf("Using kubeconfig: %s\n", kubeConfigPath)

	TestResourcePods(kubeConfigPath, namespace)
	// TestResourceConfigMaps(kubeConfigPath, namespace)

	// TestResourceCRDs(kubeConfigPath, namespace)

	// TestResourceCRDsDynamicClient(kubeConfigPath, namespace)

}
