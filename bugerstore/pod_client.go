package main

import (
	"context"
	"fmt"
	"os"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPods(clientset kubernetes.Interface, namespace string) ([]apiv1.Pod, error) {
	pods, err := clientset.CoreV1().
		Pods(namespace).
		List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		fmt.Printf("error listing pods : %v\n", err)
		os.Exit(1)
	}

	// kubectl get pods -n <namespace> displays 5 columns:

	// | NAME                          | READY | STATUS  | RESTARTS | AGE |
	// |-------------------------------|-------|---------|----------|-----|
	// | kube-flannel-ds-amd64-gdltp   | 1/1   | Running | 0        | 14d |

	fmt.Println("Listing Pods")
	for _, pod := range pods.Items {
		name := pod.Name
		status := pod.Status.Phase
		ready := 0
		var restarts int32 = 0
		if len(pod.Status.ContainerStatuses) > 0 {
			restarts = pod.Status.ContainerStatuses[0].RestartCount
		}
		diff := uint(time.Now().Sub(pod.Status.StartTime.Time).Hours())
		age := uint(diff / 24)

		fmt.Println(name, ready, status, restarts, age)

	}
	return pods.Items, nil
}

func GetPod(clientset kubernetes.Interface, namespace string, podname string) (*apiv1.Pod, error) {

	// kubectl get pod <pod_name> -n <namespace> -o yaml

	pod, err := clientset.CoreV1().
		Pods(namespace).
		Get(context.TODO(), podname, metav1.GetOptions{})

	if err != nil {
		fmt.Printf("error listing pods : %v\n", err)
		os.Exit(1)
	}

	status := pod.Status.Phase
	ready := 0
	var restarts int32 = 0
	if len(pod.Status.ContainerStatuses) > 0 {
		restarts = pod.Status.ContainerStatuses[0].RestartCount
	}

	age := uint(time.Since(pod.Status.StartTime.Time).Hours())

	fmt.Println("Get Single Pod object:")
	fmt.Println(podname, ready, status, restarts, age)

	return pod, err
}

func CreatePod(clientset kubernetes.Interface, namespace string, podname string) (*apiv1.Pod, error) {
	myPod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: podname},
		Spec: apiv1.PodSpec{
			RestartPolicy: apiv1.RestartPolicyNever,
			Containers: []apiv1.Container{
				apiv1.Container{
					Name:    "main",
					Image:   "python:3.8",
					Command: []string{"python"},
					Args:    []string{"-c", "print('hello world')"},
				},
			},
		},
	}

	pod, err := clientset.CoreV1().
		Pods(namespace).
		Create(context.TODO(), myPod, metav1.CreateOptions{})

	if err != nil {
		fmt.Printf("error creating pod : %v\n", err)
		os.Exit(1)
	}

	status := pod.Status.Phase
	ready := 0
	var restarts int32 = 0
	if len(pod.Status.ContainerStatuses) > 0 {
		restarts = pod.Status.ContainerStatuses[0].RestartCount
	}

	age := 0
	//age := uint(time.Since(pod.Status.StartTime.Time).Hours())

	fmt.Println("Created Single Pod object:")
	fmt.Println(podname, ready, status, restarts, age)

	return pod, err
}

func UpdatePod(clientset kubernetes.Interface, namespace string, mypod *apiv1.Pod) (*apiv1.Pod, error) {
	//cannot change anything in pod expect image
	mypod.Spec.Containers[0].Image = "python:3.9"

	pod, err := clientset.CoreV1().
		Pods(namespace).
		Update(context.TODO(), mypod, metav1.UpdateOptions{})

	if err != nil {
		fmt.Printf("error update pod : %v\n", err)
		os.Exit(1)
	}

	podname := pod.Name
	status := pod.Status.Phase
	ready := 0
	var restarts int32 = 0
	if len(pod.Status.ContainerStatuses) > 0 {
		restarts = pod.Status.ContainerStatuses[0].RestartCount
	}

	age := 0
	//age := uint(time.Since(pod.Status.StartTime.Time).Hours())

	fmt.Println("Updated Single Pod object:")
	fmt.Println(podname, ready, status, restarts, age)

	return pod, err
}

func DeletePod(clientset kubernetes.Interface, namespace string, podname string) error {

	err := clientset.CoreV1().
		Pods(namespace).
		Delete(context.TODO(), podname, metav1.DeleteOptions{})

	if err != nil {
		fmt.Printf("error deleting pod : %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Delete Single Pod object:", podname)
	return err
}
