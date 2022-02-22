package controller

import (
	"context"
	"fmt"
	"time"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
	burgerstorev1alpha1 "satya.com/burgerstore/pkg/apis/burgerstore.dev/v1alpha1"
)

func (c *Controller) doesCRDExist() (bool, error) {
	crd, err := c.apiextensionsclientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(context.TODO(), burgerstorev1alpha1.Name, metav1.GetOptions{})

	if err != nil {
		return false, err
	}

	// Check whether the CRD is accepted.
	for _, condition := range crd.Status.Conditions {
		if condition.Type == apiextensions.Established &&
			condition.Status == apiextensions.ConditionTrue {
			return true, nil
		}
	}

	return false, fmt.Errorf("CRD is not accepted")
}

func (c *Controller) waitCRDAccepted() error {
	err := wait.Poll(1*time.Second, 10*time.Second, func() (bool, error) {
		return c.doesCRDExist()
	})

	return err
}

// CreateCRD creates a custom resource definition,
// named TestResource.
func (c *Controller) CreateCRD() error {
	if result, _ := c.doesCRDExist(); result {
		return nil
	}

	crd := &apiextensions.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: burgerstorev1alpha1.Name,
		},
		Spec: apiextensions.CustomResourceDefinitionSpec{
			Group:   burgerstorev1alpha1.GroupName,
			Version: burgerstorev1alpha1.Version,
			Scope:   apiextensions.NamespaceScoped,
			Names: apiextensions.CustomResourceDefinitionNames{
				Plural:     burgerstorev1alpha1.Plural,
				Singular:   burgerstorev1alpha1.Singluar,
				Kind:       burgerstorev1alpha1.Kind,
				ShortNames: []string{burgerstorev1alpha1.ShortName},
			},
			Validation: &apiextensions.CustomResourceValidation{
				OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
					Type: "object",
					Properties: map[string]apiextensions.JSONSchemaProps{
						"spec": {
							Type: "object",
							Properties: map[string]apiextensions.JSONSchemaProps{
								"owner":      {Type: "string"},
								"address":    {Type: "string"},
								"currency":   {Type: "string"},
								"investment": {Type: "integer"},
							},
							Required: []string{"owner", "address"},
						},
					},
				},
			},
			AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
				{
					Name:     "owner",
					Type:     "string",
					JSONPath: ".spec.owner",
				},
				{
					Name:     "address",
					Type:     "string",
					JSONPath: ".spec.address",
				},
				{
					Name:     "currency",
					Type:     "string",
					JSONPath: ".spec.currency",
				},
				{
					Name:     "investment",
					Type:     "integer",
					JSONPath: ".spec.investment",
				},
			},
		},
	}

	_, err := c.apiextensionsclientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(context.TODO(), crd, metav1.CreateOptions{})

	if err != nil {
		klog.Fatalf(err.Error())
	}

	return c.waitCRDAccepted()
}
