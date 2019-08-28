/*
Copyright 2019 The Knative Authors

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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing/pkg/apis/sources/v1alpha1"
	scheme "knative.dev/eventing/pkg/client/clientset/versioned/scheme"
)

// ApiServerSourcesGetter has a method to return a ApiServerSourceInterface.
// A group's client should implement this interface.
type ApiServerSourcesGetter interface {
	ApiServerSources(namespace string) ApiServerSourceInterface
}

// ApiServerSourceInterface has methods to work with ApiServerSource resources.
type ApiServerSourceInterface interface {
	Create(*v1alpha1.ApiServerSource) (*v1alpha1.ApiServerSource, error)
	Update(*v1alpha1.ApiServerSource) (*v1alpha1.ApiServerSource, error)
	UpdateStatus(*v1alpha1.ApiServerSource) (*v1alpha1.ApiServerSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiServerSource, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiServerSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiServerSource, err error)
	ApiServerSourceExpansion
}

// apiServerSources implements ApiServerSourceInterface
type apiServerSources struct {
	client rest.Interface
	ns     string
}

// newApiServerSources returns a ApiServerSources
func newApiServerSources(c *SourcesV1alpha1Client, namespace string) *apiServerSources {
	return &apiServerSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiServerSource, and returns the corresponding apiServerSource object, and an error if there is any.
func (c *apiServerSources) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiServerSource, err error) {
	result = &v1alpha1.ApiServerSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apiserversources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiServerSources that match those selectors.
func (c *apiServerSources) List(opts v1.ListOptions) (result *v1alpha1.ApiServerSourceList, err error) {
	result = &v1alpha1.ApiServerSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apiserversources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiServerSources.
func (c *apiServerSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apiserversources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a apiServerSource and creates it.  Returns the server's representation of the apiServerSource, and an error, if there is any.
func (c *apiServerSources) Create(apiServerSource *v1alpha1.ApiServerSource) (result *v1alpha1.ApiServerSource, err error) {
	result = &v1alpha1.ApiServerSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apiserversources").
		Body(apiServerSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiServerSource and updates it. Returns the server's representation of the apiServerSource, and an error, if there is any.
func (c *apiServerSources) Update(apiServerSource *v1alpha1.ApiServerSource) (result *v1alpha1.ApiServerSource, err error) {
	result = &v1alpha1.ApiServerSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apiserversources").
		Name(apiServerSource.Name).
		Body(apiServerSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiServerSources) UpdateStatus(apiServerSource *v1alpha1.ApiServerSource) (result *v1alpha1.ApiServerSource, err error) {
	result = &v1alpha1.ApiServerSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apiserversources").
		Name(apiServerSource.Name).
		SubResource("status").
		Body(apiServerSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiServerSource and deletes it. Returns an error if one occurs.
func (c *apiServerSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apiserversources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiServerSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apiserversources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiServerSource.
func (c *apiServerSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiServerSource, err error) {
	result = &v1alpha1.ApiServerSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apiserversources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}