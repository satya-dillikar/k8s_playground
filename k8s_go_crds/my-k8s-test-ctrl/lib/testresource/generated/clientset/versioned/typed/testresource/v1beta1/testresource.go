/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "satya.com/my-k8s-test-ctrl/lib/testresource/generated/clientset/versioned/scheme"
	v1beta1 "satya.com/my-k8s-test-ctrl/lib/testresource/v1beta1"
)

// TestResourcesGetter has a method to return a TestResourceInterface.
// A group's client should implement this interface.
type TestResourcesGetter interface {
	TestResources(namespace string) TestResourceInterface
}

// TestResourceInterface has methods to work with TestResource resources.
type TestResourceInterface interface {
	Create(*v1beta1.TestResource) (*v1beta1.TestResource, error)
	Update(*v1beta1.TestResource) (*v1beta1.TestResource, error)
	UpdateStatus(*v1beta1.TestResource) (*v1beta1.TestResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.TestResource, error)
	List(opts v1.ListOptions) (*v1beta1.TestResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.TestResource, err error)
	TestResourceExpansion
}

// testResources implements TestResourceInterface
type testResources struct {
	client rest.Interface
	ns     string
}

// newTestResources returns a TestResources
func newTestResources(c *SatyaV1beta1Client, namespace string) *testResources {
	return &testResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the testResource, and returns the corresponding testResource object, and an error if there is any.
func (c *testResources) Get(name string, options v1.GetOptions) (result *v1beta1.TestResource, err error) {
	result = &v1beta1.TestResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TestResources that match those selectors.
func (c *testResources) List(opts v1.ListOptions) (result *v1beta1.TestResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.TestResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("testresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested testResources.
func (c *testResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("testresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a testResource and creates it.  Returns the server's representation of the testResource, and an error, if there is any.
func (c *testResources) Create(testResource *v1beta1.TestResource) (result *v1beta1.TestResource, err error) {
	result = &v1beta1.TestResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("testresources").
		Body(testResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a testResource and updates it. Returns the server's representation of the testResource, and an error, if there is any.
func (c *testResources) Update(testResource *v1beta1.TestResource) (result *v1beta1.TestResource, err error) {
	result = &v1beta1.TestResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testresources").
		Name(testResource.Name).
		Body(testResource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *testResources) UpdateStatus(testResource *v1beta1.TestResource) (result *v1beta1.TestResource, err error) {
	result = &v1beta1.TestResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("testresources").
		Name(testResource.Name).
		SubResource("status").
		Body(testResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the testResource and deletes it. Returns an error if one occurs.
func (c *testResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *testResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("testresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched testResource.
func (c *testResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.TestResource, err error) {
	result = &v1beta1.TestResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("testresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
