package main

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testresourcev1beta1 "satya.com/my-k8s-test-ctrl/lib/testresource/v1beta1"
)

// CreateObject creates a TestResource object for the test purpose.
func (c *Controller) CreateObject() error {
	object := &testresourcev1beta1.TestResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-tr2",
			Namespace: corev1.NamespaceDefault,
		},
		Spec: testresourcev1beta1.TestResourceSpec{
			Command:        "echo Hello World!",
			CustomProperty: "asdasd=1234",
		},
	}

	_, err := c.testresourceclientset.SatyaV1beta1().TestResources(corev1.NamespaceDefault).Create(object)
	return err
}
