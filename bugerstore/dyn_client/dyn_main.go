package dyn_client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"satya.com/burgerstore/pkg/apis/burgerstore.dev/v1alpha1"
)

func DynamicMain(mykubeconfig string, namespace string) {
	//kubeconfig := flag.String("kubeconfig", "/home/vivek/.kube/config", "location to your kubeconfig file")
	// mykubeconfig := "/Users/sdillikar/kind-basic.yaml"
	kubeconfig := &mykubeconfig
	fmt.Println("kubeconfig:", *kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		// handle error
		fmt.Printf("erorr %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s, getting inclusterconfig", err.Error())
		}
	}

	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, getting dyn client\n", err.Error())
	}

	unsObject, err := dynClient.Resource(schema.GroupVersionResource{
		Group:    "burgerstore.dev",
		Version:  "v1alpha1",
		Resource: "burgerstores",
	}).Namespace(namespace).Get(context.TODO(), "testfoo-crd1", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("error %s getting resource from dyn client\n", err.Error())
	}

	brgs := v1alpha1.BurgerStore{}
	// getting and setting fields on unsObject
	fmt.Printf("Go the object %s\n", unsObject.GetName())

	// how do I convert unsObject into a typed object
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unsObject.UnstructuredContent(), &brgs)
	if err != nil {
		fmt.Printf("error %s, converting unstructured to kluster type", err.Error())
	}
	fmt.Printf("the concrete type that we got is %+v\n", brgs)

	// encode back to JSON
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	enc.Encode(brgs)

	infFactory := dynamicinformer.NewDynamicSharedInformerFactory(dynClient, 10*time.Minute)

	c := newController(dynClient, infFactory)
	infFactory.Start(make(<-chan struct{}))
	c.run(make(<-chan struct{}))
}
